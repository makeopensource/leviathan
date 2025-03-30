export * from "./generated/docker_rpc/v1/docker_pb"
export * from "./generated/jobs/v1/jobs_pb"
export * from "./generated/labs/v1/labs_pb"
export * from "./generated/stats/v1/stats_pb"
export * from "./generated/types/v1/types_pb"
export * from "@connectrpc/connect-node"
export * from "@connectrpc/connect"

// Base type with common properties
type BaseFileData = {
    filename: string,
    filedata: Blob
}

// Create specific types with preset field names
export type SubmissionFile = BaseFileData & { fieldName: 'submissionFiles' }
export type DockerFile = BaseFileData & { fieldName: 'dockerfile' }
export type LabFile = BaseFileData & { fieldName: 'labFiles' }

// Union type of all possible file data types
type FileData = SubmissionFile | DockerFile | LabFile

export async function UploadLabFiles(basePath: String, dockerfile: DockerFile, files: Array<LabFile>) {
    const url = `${basePath}/files.v1/upload/lab`
    return UploadMultipartForm(url, [...files, dockerfile])
}

export async function UploadSubmissionFiles(basePath: String, files: Array<SubmissionFile>) {
    const url = `${basePath}/files.v1/upload/submission`
    return UploadMultipartForm(url, files)
}

async function UploadMultipartForm(url: string, files: Array<FileData>,) {
    const formData = new FormData();

    for (const info of files) {
        formData.append(
            info.fieldName,
            info.filedata,
            info.filename,
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

