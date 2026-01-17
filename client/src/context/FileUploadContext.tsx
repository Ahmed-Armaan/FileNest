import { createContext, useContext, useEffect, useState, type ReactNode } from "react"

export type FileUploadStatus = {
	uploadId: string
	file: File
	uploadedSize: number
}

interface FileUploadContextType {
	uploading: boolean
	fileData: FileUploadStatus[]
	addFile: (file: File) => FileUploadStatus
	clearFiles: () => void
	removeFile: (fileUpload: FileUploadStatus) => void
}

const FileUploadContext = createContext<FileUploadContextType | null>(null)

export function FileUploadContextProvider({ children }: { children: ReactNode }) {
	const [fileData, setFileData] = useState<FileUploadStatus[]>([])
	const [uploading, setUploading] = useState<boolean>(false)

	useEffect(() => {
		setUploading(() => fileData.length > 0)
	}, [fileData])

	const addFile = (file: File): FileUploadStatus => {
		const newFileData: FileUploadStatus = {
			uploadId: crypto.randomUUID(),
			file,
			uploadedSize: 0,
		}

		setFileData(prev => [...prev, newFileData])
		return newFileData
	}

	const clearFiles = () => {
		setFileData([])
	}

	const removeFile = (fileUpload: FileUploadStatus) => {
		setFileData((prev) => {
			return prev.filter((fileUploadItem) => {
				return fileUploadItem.uploadId !== fileUpload.uploadId
			})
		})
	}

	return (
		<FileUploadContext.Provider
			value={{
				uploading: uploading,
				fileData: fileData,
				addFile: addFile,
				clearFiles: clearFiles,
				removeFile: removeFile,
			}}>
			{children}
		</FileUploadContext.Provider>
	)
}

export function useFileUploadContext() {
	const ctx = useContext(FileUploadContext)
	if (!ctx) {
		throw new Error("useFileUploadContext must be used inside provider")
	}
	return ctx
}
