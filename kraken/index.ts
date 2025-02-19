import express, {Request, Response} from 'express';
import multer from 'multer';

import {createClient, createConnectTransport, JobLogRequest, JobService, NewJobRequest} from "leviathan-node-sdk"
import path from "node:path";

import {WebSocketServer} from 'ws';

const leviUrl = process.env.LEVIATHAN_URL || 'http://localhost:9221';
console.log(`Leviathan url set to ${leviUrl}`)

const transport = createConnectTransport({
    baseUrl: leviUrl,
    httpVersion: "2"
});
const jobService = createClient(JobService, transport)

const app = express();
const upload = multer();
const port = 3000;

app.use(express.static(path.join(__dirname, 'ui')));
// Define the endpoint
app.post('/submit',
    upload.fields([
        {name: 'grader', maxCount: 1},
        {name: 'makefile', maxCount: 1},
        {name: 'student', maxCount: 1},
        {name: 'dockerfile', maxCount: 1},
    ]),
    async (req: Request, res: Response) => {
        try {
            const imageTag = req.body.imageTag as string;
            const jobTimeout = parseInt(req.body.timeoutInSeconds, 10);
            if (isNaN(jobTimeout)) {
                res.status(400).send('Invalid timeout');
                return;
            }

            const files = req.files as { [fieldname: string]: Express.Multer.File[] };

            const grader = files['grader'][0]
            const makefile = files['makefile'][0]
            const student = files['student'][0]
            const dockerfile = files['dockerfile'][0]

            const job = <NewJobRequest>{
                jobTimeoutInSeconds: BigInt(jobTimeout),
                imageName: imageTag,
                makeFile: {
                    content: new Uint8Array(makefile.buffer),
                    filename: makefile.originalname,
                },
                graderFile: {
                    content: new Uint8Array(grader.buffer),
                    filename: grader.originalname,
                },
                studentSubmission: {
                    content: new Uint8Array(student.buffer),
                    filename: student.originalname,
                },
                dockerFile: {
                    content: new Uint8Array(dockerfile.buffer),
                    filename: dockerfile.originalname,
                },
            }

            const jobRes = await jobService.newJob(job)
            res.status(200).redirect(`/results.html?jobid=${jobRes.jobId}`);
        } catch (error: any) {
            res.status(400).json({error: error.message});
        }
    });

const server = app.listen(port, () => {
    console.log(`Server is running on port ${port}`);
});

const wss = new WebSocketServer({server, path: "/ws"});

wss.on('connection', async (ws, req) => {
    const url = new URL(req.url!, `ws://${req.headers.host}`); // Important: Construct a full URL
    const searchParams = new URLSearchParams(url.search);
    const jobId = searchParams.get('jobid') as string;
    console.log("Job ID:", jobId);

    if (!jobId) {
        ws.close(400, "Invalid job ID");
        return;
    }

    const dataStream = jobService.streamStatus(<JobLogRequest>{jobId: jobId})

    for await (const chunk of dataStream) {
        if (!chunk.jobInfo) {
            console.warn("Empty job state")
            continue
        }

        const {jobTimeout, $unknown, $typeName, ...rest} = chunk.jobInfo!

        console.log("Job", rest);
        console.log(chunk.logs)

        ws.send(JSON.stringify({
            logs: chunk.logs,
            jobStatus: rest,
        }));
    }

    console.log("Job ID:", jobId, "done streaming");
});

wss.on('close', () => {

})