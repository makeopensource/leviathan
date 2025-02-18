import express, {Request, Response} from 'express';
import multer from 'multer';

import {createClient, createConnectTransport, JobService, NewJobRequest} from "leviathan-node-sdk"
import path from "node:path";


const transport = createConnectTransport({
    baseUrl: "http://localhost:9221",
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
            res.status(200).json({jobId: jobRes.jobId});
        } catch (error: any) {
            res.status(400).json({error: error.message});
        }
    });

app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
});
