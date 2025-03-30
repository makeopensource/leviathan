export * from "./generated/docker_rpc/v1/docker_pb"
export * from "./generated/jobs/v1/jobs_pb"
export * from "./generated/labs/v1/labs_pb"
export * from "./generated/stats/v1/stats_pb"
export * from "./generated/types/v1/types_pb"
export * from "@connectrpc/connect-node"
export * from "@connectrpc/connect"

type FileData = {
    fieldName: string,
    filename: string,
    filedata: Blob
}

export type SubmissionFile = Omit<FileData, 'fieldName'> & {
    fieldName: 'submissionFiles'  // Preset field name for submission files
}

export type DockerFile = Omit<FileData, 'fieldName'> & {
    fieldName: "dockerfile"  // Preset field name for submission files
}

export type LabFile = Omit<FileData, 'fieldName'> & {
    fieldName: 'labFiles'  // Preset field name
}


export async function UploadLabFiles(basePath: String, dockerfile: DockerFile, files: Array<LabFile>) {
    const url = `${basePath}/v1/files/upload/submission`
    return UploadMultipartForm(url, files)
}

export async function UploadSubmissionFiles(basePath: String, files: Array<SubmissionFile>) {
    const url = `${basePath}/v1/files/upload/lab`
    return UploadMultipartForm(url, files)
}

async function UploadMultipartForm(url: string, files: Array<FileData>,) {
    const formData = new FormData();

    for (const file of files) {
        formData.append(
            file.fieldName,
            file.filedata,
            file.filename,
        );
    }

    const response = await fetch(url, {
        method: 'POST',
        body: formData,
    });

    if (!response.ok) {
        throw new Error(`Upload failed with status: ${response.status}\n${response.statusText}\n${response.body}`);
    }

    const data = await response.json();
    if (data && data.folderId) {
        return data.folderId as string;
    } else {
        throw new Error('Response did not contain an folderID');
    }
}

