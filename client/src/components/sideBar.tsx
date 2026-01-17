import { useEffect, useState } from "react"
import DragAndDrop from "./dragAndDrop";
import UploadFile from "../utils/upload";
import { useFileUploadContext, type FileUploadStatus } from "../context/FileUploadContext";

function SideBar(props: { currDirId: string }) {
	const { uploading, fileData, addFile, clearFiles, removeFile } = useFileUploadContext()

	const [showUploader, toggleUploader] = useState<boolean>(false)
	const [files, setFiles] = useState<File[]>([]);
	const [startUpload, setStartUpload] = useState<boolean>(false)

	useEffect(() => {
		if (!startUpload) return

		files.forEach((file) => {
			const fileData: FileUploadStatus = addFile(file)
			UploadFile(props.currDirId, fileData, addFile, clearFiles, removeFile)
		})
	}, [startUpload])

	// const uploadFile = async (file: File) => {
	// 	const res = await fetch(
	// 		`${import.meta.env.VITE_BACKEND_URL}/api/upload`,
	// 		{ credentials: "include" }
	// 	)
	//
	// 	if (!res.ok) return alert("Failed to get upload URL")
	//
	// 	const { uploadUrl } = await res.json()
	//
	// 	await fetch(uploadUrl, {
	// 		method: "PUT",
	// 		body: file,
	// 	})
	// }

	return (
		<div
			className="flex flex-col h-full p-5 gap-3"
			style={{ background: "var(--bg-secondary)" }}
		>
			{/*<div>
				<input
					type="file"
					onChange={(e) => {
						if (e.target.files?.[0]) {
							uploadFile(e.target.files[0])
						}
					}}
				/>
			</div>*/}

			<div>
				<button
					onClick={() => toggleUploader(!showUploader)}
					className="w-full text-left px-3 py-2 border rounded-md hover:opacity-80 transition">
					📤 Upload File
				</button>
			</div>

			{showUploader && (
				<DragAndDrop onFileUpload={setFiles} toggleUploader={toggleUploader} setStartUpload={setStartUpload} />
			)}

			<div>
				<button className="w-full text-left px-3 py-2 border rounded-md hover:opacity-80 transition">
					➕ Create Folder
				</button>
			</div>
		</div>
	)
}

export default SideBar
