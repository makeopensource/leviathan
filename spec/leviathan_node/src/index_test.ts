import {DockerFile, UploadLabFiles, UploadSubmissionFiles} from "./index";

async function test_lab_upload() {
    const basePath = "http://localhost:9221"

    const res = await UploadLabFiles(basePath, <DockerFile>{
        fieldName: 'dockerfile',
        filename: "test.txt",
        filedata: new Blob([Buffer.from("test")], {type: "text/plain"})
    }, [
        {
            fieldName: 'labFiles',
            filename: "test.txt",
            filedata: new Blob([Buffer.from("test")], {type: "text/plain"}),

        }, {
            fieldName: 'labFiles',
            filename: "test.txt",
            filedata: new Blob([Buffer.from("test")], {type: "text/plain"}),

        }, {
            fieldName: 'labFiles',
            filename: "test.txt",
            filedata: new Blob([Buffer.from("test")], {type: "text/plain"}),

        }, {
            fieldName: 'labFiles',
            filename: "test.txt",
            filedata: new Blob([Buffer.from("test")], {type: "text/plain"}),

        }, {
            fieldName: 'labFiles',
            filename: "test.txt",
            filedata: new Blob([Buffer.from("test")], {type: "text/plain"}),
        },
    ])


    console.log(res)
}

const basePath = "http://localhost:9221"

async function test_submission_upload() {

    const res = await UploadSubmissionFiles(basePath, [{
        fieldName: 'submissionFiles',
        filename: "test.txt",
        filedata: new Blob([Buffer.from("test")], {type: "text/plain"})
    }])

    console.log(res)
}

test_submission_upload()
test_lab_upload()
// test()