import { useEffect, useState } from "react"
import DragAndDrop from "./dragAndDrop";
import UploadFile from "../utils/upload";
import { useFileUploadContext, type FileUploadStatus } from "../context/FileUploadContext";
import { useFileRefreshContext } from "../context/filesRefreshContext";
import { FiCheck, FiX } from "react-icons/fi";
import { SharedList } from "./sharedList";

function SideBar(props: { currDirId: string }) {
	const { addFile, removeFile, setUploadData, setUploadedSize } = useFileUploadContext()
	const { triggerFilesRefresh } = useFileRefreshContext()

	const [showUploader, toggleUploader] = useState<boolean>(false)
	const [files, setFiles] = useState<File[]>([]);
	const [startUpload, setStartUpload] = useState<boolean>(false)
	const [creatingDir, setCreatingDir] = useState(false)
	const [dirName, setDirName] = useState("")

	useEffect(() => {
		if (!startUpload) return

		files.forEach((file) => {
			const fileData: FileUploadStatus = addFile(file)
			UploadFile(fileData, removeFile, setUploadData, setUploadedSize, props.currDirId, triggerFilesRefresh)
		})
	}, [startUpload])

	const createDirectory = async (dirName: string) => {
		const reqUrl = new URL(`${import.meta.env.VITE_BACKEND_URL}/api/create_directory`)
		reqUrl.searchParams.set("dirName", dirName)
		reqUrl.searchParams.set("parentId", props.currDirId)

		const res = await fetch(reqUrl, {
			method: "PUT",
			credentials: "include",
		})

		if (!res.ok) {
			throw new Error("failed to create directory")
		}

		triggerFilesRefresh()
	}

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

			<div className="flex flex-col gap-2">
				<button
					className="w-full text-left px-3 py-2 border rounded-md hover:opacity-80 transition"
					onClick={() => {
						setCreatingDir(!creatingDir)
						setDirName("")
					}}
				>
					➕ Create Folder
				</button>

				{creatingDir && (
					<div className="flex gap-2">
						<input
							autoFocus
							className="flex-1 px-3 py-2 border rounded-md bg-transparent"
							placeholder="Folder name"
							value={dirName}
							onChange={(e) => setDirName(e.target.value)}
							onKeyDown={(e) => {
								if (e.key === "Enter") {
									if (!dirName.trim()) return
									createDirectory(dirName)
									setCreatingDir(false)
								}
								if (e.key === "Escape") {
									setCreatingDir(false)
								}
							}}
						/>

						<button
							className="p-2 border rounded-md hover:opacity-80 transition"
							onClick={() => {
								if (!dirName.trim()) return
								createDirectory(dirName)
								setCreatingDir(false)
							}}
							title="Create"
						>
							<FiCheck />
						</button>

						<button
							className="p-2 border rounded-md hover:opacity-80 transition"
							onClick={() => setCreatingDir(false)}
							title="Cancel"
						>
							<FiX />
						</button>
					</div>
				)}

			</div>
			<SharedList />
			<div>

			</div>
		</div>
	)
}

export default SideBar
