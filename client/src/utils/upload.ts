import type { FileUploadStatus } from "../context/FileUploadContext";
import type { UploadActionRes } from "../types/uploadTypes";

const FIVE_MB = 5 * 1024 * 1024 // bytes

export default async function UploadFile(
	fileUploadStatus: FileUploadStatus,
	removeFile: (fileUpload: FileUploadStatus) => void,
	setUploadData: (fileData: FileUploadStatus, UploadData: UploadActionRes) => void,
	setUploadedSize: (fileData: FileUploadStatus, size: { uploadedSize: number }) => void
) {
	const file = fileUploadStatus.file
	const totalChunks = Math.ceil(file.size / FIVE_MB)

	await createUpload(fileUploadStatus, setUploadData)

	for (let i = 0; i < totalChunks; i++) {
		const start = i * FIVE_MB
		const end = Math.min(start + FIVE_MB, file.size)

		const blob = file.slice(start, end)
		const buffer = await blob.arrayBuffer()

		const partNumber = i + 1
		await uploadChunk(fileUploadStatus, buffer, partNumber)

		setUploadedSize(fileUploadStatus, { uploadedSize: end })
	}

	// show notification
	// use a global state {showNotification, notificationMessage} to trigger.
	// So Can be reaised from anywhere in the app
	removeFile(fileUploadStatus)
}

// create a upload action
async function createUpload(file: FileUploadStatus,
	setUploadData: (fileData: FileUploadStatus, UploadData: UploadActionRes) => void
) {
	if (!file) throw new Error("No file Unavailable")
	if (!setUploadData) throw new Error("No data setter Unavailable")

	const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/get_upload_url`,
		{
			method: "POST",
			credentials: "include",
		}
	)
	if (!res.ok) return alert("Failed to get upload URL")
	const uploadingData: UploadActionRes = await res.json()
	setUploadData(file, uploadingData)
}

// upload a chunk
async function uploadChunk(file: FileUploadStatus, chunk: ArrayBuffer, chunkIndex: number) {
	const getUploadUrl = async (): Promise<string> => {
		const reqUrl = new URL(`${import.meta.env.VITE_BACKEND_URL}/get_upload_url/parts`)
		reqUrl.searchParams.set("uploadId", file.uploadId)
		reqUrl.searchParams.set("objectKey", file.objectKey)
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

		console.log(`ETAG: ${res.headers.get("ETag")}`)
	}

	// chunk upload entry point
	try {
		const url = await getUploadUrl()
		await upload(url)
	} catch (err) {
		console.log(`File Upload error: ${err}`)
	}
}
