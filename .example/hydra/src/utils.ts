import fs from "node:fs";

export async function readFileAsBytes(filePath: string): Promise<Uint8Array> {
    try {
        const buffer = fs.readFileSync(filePath);
        return new Uint8Array(buffer);
    } catch (error) {
        console.error('Error reading file:', error);
        throw error;
    }
}

