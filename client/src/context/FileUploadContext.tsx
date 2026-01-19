import { createContext, useContext, useEffect, useState, type ReactNode } from "react"
import type { UploadActionRes } from "../types/uploadTypes"

export type FileUploadStatus = {
	Id: string
	file: File
	uploadedSize: number
	uploadId: string
	objectKey: string
}

interface FileUploadContextType {
	uploading: boolean
	fileData: FileUploadStatus[]
	addFile: (file: File) => FileUploadStatus
	clearFiles: () => void
	removeFile: (fileUpload: FileUploadStatus) => void
	setUploadData: (fileData: FileUploadStatus, UploadData: UploadActionRes) => void
	setUploadedSize: (fileData: FileUploadStatus, size: { uploadedSize: number }) => void
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
			Id: crypto.randomUUID(),
			file,
			uploadedSize: 0,
			uploadId: "",
			objectKey: "",
		}

		setFileData(prev => [...prev, newFileData])
		return newFileData
	}

	const setUploadData = (fileData: FileUploadStatus, UploadData: UploadActionRes) => {
		setFileData((prev) =>
			prev.map((item) =>
				(item.Id === fileData.Id) ? {
					...item,
					...UploadData,
				} : item
			)
		)
	}

	const setUploadedSize = (fileData: FileUploadStatus, size: { uploadedSize: number }) => {
		setFileData((prev) =>
			prev.map((item) =>
				(item.Id === fileData.Id) ? {
					...item,
					...size,
				} : item
			)
		)
	}

	const clearFiles = () => {
		setFileData([])
	}

	const removeFile = (fileUpload: FileUploadStatus) => {
		setFileData((prev) => {
			return prev.filter((fileUploadItem) => {
				return fileUploadItem.Id !== fileUpload.Id
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
				setUploadData: setUploadData,
				setUploadedSize: setUploadedSize,
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
