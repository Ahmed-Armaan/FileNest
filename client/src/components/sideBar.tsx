import { useEffect, useState } from "react"
import DragAndDrop from "./dragAndDrop";
import UploadFile from "../utils/upload";
import { useFileUploadContext, type FileUploadStatus } from "../context/FileUploadContext";
import { useFileRefreshContext } from "../context/filesRefreshContext";

function SideBar(props: { currDirId: string }) {
	const { addFile, removeFile, setUploadData, setUploadedSize } = useFileUploadContext()
	const { triggerFilesRefresh } = useFileRefreshContext()

	const [showUploader, toggleUploader] = useState<boolean>(false)
	const [files, setFiles] = useState<File[]>([]);
	const [startUpload, setStartUpload] = useState<boolean>(false)

	useEffect(() => {
		if (!startUpload) return
		console.log(`Files`)
		// console.log(props.currDirId) -- Empty string

		files.forEach((file) => {
			const fileData: FileUploadStatus = addFile(file)

			console.log(`Files: ${props.currDirId}`)

			UploadFile(fileData, removeFile, setUploadData, setUploadedSize, props.currDirId, triggerFilesRefresh)
		})
	}, [startUpload])

	return (
		<div
			className="flex flex-col h-full p-5 gap-3"
			style={{ background: "var(--bg-secondary)" }}
		>
			<div>
				<button
					onClick={() => toggleUploader(!showUploader)}
					className="w-full text-left px-3 py-2 border rounded-md hover:opacity-80 transition">
					📤 Upload File
				</button>
			</div>

			{showUploader && (
				<DragAndDrop ParentId={props.currDirId} onFileUpload={setFiles} toggleUploader={toggleUploader} setStartUpload={setStartUpload} />
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
