import type { FileUploadStatus } from "../context/FileUploadContext";

const FIVE_MB = 5 * 1024 * 1024 // bytes

export default function UploadFile(parentDirId: string,
	FileUploadData: FileUploadStatus,
	addFiles: (file: File) => FileUploadStatus,
	clearFiles: () => void,
	removeFile: (fileUpload: FileUploadStatus) => void) {

	const file = FileUploadData.file
	if (file) {
		addFiles(file)
		var totalChunks = Math.ceil(file.size / FIVE_MB)
	}
	var chunkIndex = 0
	//var uploadedChunks = 0

	//chunked file upload
	const onFileReadEvent = (err: ProgressEvent<FileReader>) => {
		if (err.target?.error === null) {
			if (err.target.result instanceof ArrayBuffer) {
				uploadToS3(err.target.result, parentDirId).then(() => {
					chunkIndex++
				}).finally(() => {
					chunkReader(chunkIndex, file)
				})
			}
			else throw new Error("Cant read file as bytes")
		}

		else throw new Error("No file uploaded")
	}

	//read the file into Array buffer and set the event to handle update
	const chunkReader = (offset: number, file: File) => {
		if (offset === totalChunks) {
			removeFile(FileUploadData)
		}
		const fileReader = new FileReader();
		const currChunk = file.slice(offset * FIVE_MB, (offset + 1) * FIVE_MB)
		fileReader.onload = onFileReadEvent
		fileReader.readAsArrayBuffer(currChunk)
	}

	// Upload starting point
	chunkReader(0, file)
}

async function uploadToS3(file: ArrayBuffer, parentDirId: string) {
	const res = await fetch(
		`${import.meta.env.VITE_BACKEND_URL}/api/upload`,
		{ credentials: "include" }
	)

	if (!res.ok) return alert("Failed to get upload URL")

	const { uploadUrl } = await res.json()

	await fetch(uploadUrl, {
		method: "PUT",
		body: file,
	})
}
