import React, { useEffect, useState } from "react"

type DragAndDropProps = {
	onFileUpload: (files: File[]) => void
	ParentId: string
	toggleUploader: (bool: boolean) => void
	setStartUpload: (bool: boolean) => void
}

function DragAndDrop(props: DragAndDropProps) {
	const [files, setFiles] = useState<File[]>([])

	useEffect(() => {
		props.onFileUpload(files)
	}, [files, props.onFileUpload])

	const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
		const uploadedFiles = event.target.files
		if (!uploadedFiles || uploadedFiles.length === 0) return

		const filesArray = Array.from(uploadedFiles)
		setFiles(prevFiles => [...prevFiles, ...filesArray])
	}

	const handleDrop = (event: React.DragEvent<HTMLDivElement>) => {
		console.log("Got new File Drsopped")
		event.preventDefault()
		const droppedFiles = event.dataTransfer.files
		if (droppedFiles.length > 0) {
			const newFiles = Array.from(droppedFiles)
			setFiles((prevFiles) => [...prevFiles, ...newFiles])
		}
	}

	const cancelUpload = () => {
		setFiles([])
		props.onFileUpload([])
		props.toggleUploader(false)
	}

	const startUpload = () => {
		props.setStartUpload(true)
		props.toggleUploader(false)
	}

	return (
		<>
			<div
				className="fixed inset-0 flex items-center justify-center"
				onDrop={handleDrop}
				onDragOver={(e) => e.preventDefault()}
			>
				<div>

					<label
						htmlFor="file-input"
						className="flex flex-col items-center justify-center gap-3
						w-96 min-h-[10rem] border border-dashed rounded-md
                     cursor-pointer select-none px-6 py-4 text-center"
					>
						<input
							id="file-input" type="file" multiple className="hidden" onChange={handleFileChange} />

						<div className="flex flex-col justify-center gap-3 max-w-full">
							<div className="text-sm opacity-80">
								Drag and drop files here<br />or click to browse
							</div>

							{files.length > 0 && (
								<div className="mt-4 w-full text-left">
									<div className="text-xs opacity-60 mb-2">
										Selected files
									</div>

									<ul className="space-y-1 text-sm">
										{files.map((file) => (
											<li key={file.name} className="flex items-start gap-2">
												<span aria-hidden className="shrink-0">📄</span>
												<span className="min-w-0 break-words">
													{file.name}
												</span>
											</li>
										))}
									</ul>
								</div>
							)}

							<button className="w-full text-left px-3 py-2 border rounded-md flex items-center justify-center
							hover:opacity-80 transition"
								onClick={cancelUpload} >
								Cancel
							</button>
							<button className="w-full text-left px-3 py-2 border rounded-md flex items-center justify-center
							hover:opacity-80 transition"
								onClick={startUpload}>
								Start Upload
							</button>

						</div>
					</label>

				</div>
			</div>
		</>
	)
}

export default DragAndDrop
