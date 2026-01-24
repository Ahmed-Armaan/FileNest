import type { FileUploadStatus } from "../context/FileUploadContext";
import type { UploadActionRes } from "../types/uploadTypes";

type CompletedPart = {
	etag: string,
	partNumber: number,
}

const CompletedParts: CompletedPart[] = []
const FIVE_MB = 5 * 1024 * 1024 // bytes

export default async function UploadFile(
	fileUploadStatus: FileUploadStatus,
	removeFile: (fileUpload: FileUploadStatus) => void,
	setUploadData: (fileData: FileUploadStatus, UploadData: UploadActionRes) => void,
	setUploadedSize: (fileData: FileUploadStatus, size: { uploadedSize: number }) => void,
	currDirId: string,
	triggerFileRefresh: () => void
) {
	const file = fileUploadStatus.file
	const totalChunks = Math.ceil(file.size / FIVE_MB)

	const uploadData = await createUpload(fileUploadStatus, setUploadData)

	for (let i = 0; i < totalChunks; i++) {
		const start = i * FIVE_MB
		const end = Math.min(start + FIVE_MB, file.size)

		const blob = file.slice(start, end)
		const buffer = await blob.arrayBuffer()

		const partNumber = i + 1
		await uploadChunk(buffer, partNumber, uploadData)

		setUploadedSize(fileUploadStatus, { uploadedSize: end })
	}

	// show notification
	// use a global state {showNotification, notificationMessage} to trigger.
	// So Can be reaised from anywhere in the app
	const completed = await completeUpload(fileUploadStatus, currDirId, uploadData.objectKey, uploadData.uploadId, triggerFileRefresh)
	if (completed)
		removeFile(fileUploadStatus)
}

// create a upload action
async function createUpload(file: FileUploadStatus,
	setUploadData: (fileData: FileUploadStatus, UploadData: UploadActionRes) => void
): Promise<UploadActionRes> {
	if (!file) throw new Error("No file Unavailable")
	if (!setUploadData) throw new Error("No data setter Unavailable")

	const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/api/get_upload_url`,
		{
			method: "POST",
			credentials: "include",
		}
	)
	if (!res.ok) alert("Failed to get upload URL")
	const uploadingData: UploadActionRes = await res.json()
	setUploadData(file, uploadingData)
	return uploadingData
}

// upload a chunk
async function uploadChunk(chunk: ArrayBuffer, chunkIndex: number, uploadData: UploadActionRes) {

	// this function gets the presigned URl for the current chunk
	const getUploadUrl = async (): Promise<string> => {
		const reqUrl = new URL(`${import.meta.env.VITE_BACKEND_URL}/api/get_upload_url/parts`)
		reqUrl.searchParams.set("uploadId", uploadData.uploadId)
		reqUrl.searchParams.set("objectKey", uploadData.objectKey)
		reqUrl.searchParams.set("partNumber", chunkIndex.toString())

		const res = await fetch(reqUrl.toString(), {
			method: "POST",
			credentials: "include",
		})
		if (!res.ok) {
			throw new Error("Failed to fetch URL")
		}

		const data = await res.json()
		const uploadUrl = data.url
		return uploadUrl
	}

	const upload = async (url: string) => {
		const res = await fetch(url, {
			method: "PUT",
			body: chunk,
		})
		if (!res.ok) {
			throw new Error("Failed to upload chunk")
		}

		const etag = res.headers.get("ETag")
		if (!etag) {
			throw new Error("no Etag received")
		}
		console.log(`ETAG: ${etag}`)
		CompletedParts.push({
			etag: etag,
			partNumber: chunkIndex,
		})
	}

	// chunk upload entry point
	try {
		const url = await getUploadUrl()
		await upload(url)
	} catch (err) {
		console.log(`File Upload error: ${err}`)
	}
}

async function completeUpload(file: FileUploadStatus, currDirId: string, objectKey: string, uploadId: string,
	triggerFileRefresh: () => void): Promise<boolean> {
	const reqUrl = new URL(`${import.meta.env.VITE_BACKEND_URL}/api/complete_upload`)
	reqUrl.searchParams.set("name", file.file.name)
	reqUrl.searchParams.set("parentId", currDirId)
	reqUrl.searchParams.set("objectKey", objectKey)
	reqUrl.searchParams.set("uploadId", uploadId)
	reqUrl.searchParams.set("size", file.file.size.toString())

	const res = await fetch(reqUrl.toString(), {
		method: "POST",
		credentials: "include",
		headers: { "Content-Type": "application/json" },
		body: JSON.stringify(CompletedParts),
	})

	if (!res.ok) {
		throw new Error("Failed to complete upload")
	}
	triggerFileRefresh()
	return true
}
