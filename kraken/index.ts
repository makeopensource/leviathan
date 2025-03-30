import express, {Request, Response} from 'express';
import multer from 'multer';

import {
    createClient,
    createConnectTransport,
    DockerFile,
    JobLogRequest,
    JobService,
    LabData,
    LabFile,
    LabService,
    NewJobRequest,
    NewLabRequest,
    SubmissionFile,
    UploadLabFiles,
    UploadSubmissionFiles
} from "leviathan-node-sdk"
import path from "node:path";

import {WebSocketServer} from 'ws';

const leviUrl = process.env.LEVIATHAN_URL || 'http://localhost:9221';
console.log(`Leviathan url set to ${leviUrl}`)

const transport = createConnectTransport({
    baseUrl: leviUrl,
    httpVersion: "2"
});
const jobService = createClient(JobService, transport)
const labService = createClient(LabService, transport)

const app = express();
const upload = multer();
const port = 3000;

app.use(express.static(path.join('ui/')));
// Define the endpoint
app.post('/submit',
    upload.fields([
        {name: 'fileList', maxCount: 10},
        {name: 'dockerfile', maxCount: 1}
    ]),
    async (req: Request, res: Response) => {
        try {
            const jobTimeout = parseInt(req.body.timeoutInSeconds, 10);
            if (isNaN(jobTimeout)) {
                res.status(400).send('Invalid timeout');
                return;
            }

            let autolabMode = (req.body.autolabCompatibilityMode as String) == "on";

            let entryCmd = autolabMode ? "" : req.body.entryCmd as string;
            entryCmd = entryCmd.trim()
            if ((!autolabMode && entryCmd === "") || entryCmd.startsWith("&&") || entryCmd.endsWith("&&")) {
                res.status(400).send('Invalid entry command must not start or end with && or empty');
                return
            }

            let memory = req.body.memory as number
            let cpuCore = req.body.cpuCores as number
            let pids = req.body.pidLimit as number

            if (!memory || !cpuCore || !pids) {
                res.status(400).send('Invalid machine limits');
                return
            }

            const files = req.files as { [fieldname: string]: Express.Multer.File[] };
            const jobFiles = files['fileList'].map(value => <LabFile>{
                fieldName: "labFiles",
                filename: value.originalname,
                filedata: bufferToBlob(value),
            })
            if (jobFiles.length < 2) {
                res.status(400).send('at least 2 files are required');
                return
            }

            const {filename, filedata} = jobFiles.pop()!
            const submission = <SubmissionFile>{
                fieldName: "submissionFiles",
                filename: filename,
                filedata: filedata
            }


            const dockerfile = files['dockerfile'][0]
            const jobFolderID = await UploadLabFiles(leviUrl, <DockerFile>{
                fieldName: "dockerfile",
                filename: dockerfile.filename,
                filedata: bufferToBlob(dockerfile),
            }, jobFiles)

            const lab = <NewLabRequest>{
                labData: <LabData>{
                    autolabCompatibilityMode: autolabMode,
                    entryCmd: entryCmd,
                    jobTimeoutInSeconds: BigInt(jobTimeout),
                    limits: {
                        PidLimit: pids,
                        CPUCores: cpuCore,
                        memoryInMb: memory,
                    },
                    labname: "testlab"
                },
                tmpFolderId: jobFolderID
            }
            const labId = await labService.newLab(lab)

            const submissionTmpID = await UploadSubmissionFiles(leviUrl, [submission])
            const job = <NewJobRequest>{
                labID: labId.labId,
                tmpSubmissionFolderId: submissionTmpID,
            }

            const jobRes = await jobService.newJob(job)
            res.status(200).redirect(`/results.html?jobid=${jobRes.jobId}`);
        } catch (error: any) {
            res.status(400).json({error: error.message});
        }
    });

const server = app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});

const wss = new WebSocketServer({server, path: "/ws"});

wss.on('connection', async (ws, req) => {
    const url = new URL(req.url!, `${req.headers.origin!.startsWith("https") ? "wss" : "ws"}://${req.headers.host}`); // Important: Construct a full URL
    const searchParams = new URLSearchParams(url.search);
    const jobId = searchParams.get('jobid') as string;
    console.log("Job ID:", jobId);

    if (!jobId) {
        ws.close(400, "Invalid job ID");
        return;
    }

    const controller = new AbortController();
    const dataStream = jobService.streamStatus(<JobLogRequest>{jobId: jobId}, {signal: controller.signal})

    ws.on("close", () => {
        console.log("disconnected")
        controller.abort()
    })
    try {
        for await (const chunk of dataStream) {
            if (!chunk.jobInfo) {
                console.warn("Empty job state")
                continue
            }

            const {$unknown, $typeName, ...rest} = chunk.jobInfo!

            console.log("Job", rest);
            console.log(chunk.logs)

            ws.send(JSON.stringify({
                logs: chunk.logs,
                jobStatus: rest,
            }));
        }
    } catch (e) {
        console.error(e)
    }

    console.log("Job ID:", jobId, "done streaming");
});

function bufferToBlob(multerFile: Express.Multer.File): Blob {
    return new Blob([multerFile.buffer], {type: multerFile.mimetype});
}