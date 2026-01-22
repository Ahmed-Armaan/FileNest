import { createContext, useContext, useState, type ReactNode } from "react"

interface FilesRefreshContextType {
	fileRefreshTrigger: number
	triggerFilesRefresh: () => void
}

const FilesRefreshContext = createContext<FilesRefreshContextType | null>(null)

export function FilesRefreshContextprovided({ children }: { children: ReactNode }) {
	const [filesRefresh, triggerFilesRefresh] = useState<number>(0)

	const RefreshFiles = () => {
		triggerFilesRefresh(prev => (prev + 1) % 2)
	}

	return (
		<FilesRefreshContext.Provider
			value={{
				fileRefreshTrigger: filesRefresh,
				triggerFilesRefresh: RefreshFiles,
			}}>
			{children}
		</ FilesRefreshContext.Provider>
	)
}

export function useFileRefreshContext() {
	const ctx = useContext(FilesRefreshContext)
	if (!ctx) {
		throw new Error("useFileRefreshContext must be used inside provider")
	}
	return ctx
}
