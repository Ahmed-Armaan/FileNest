import type { FileTreeNode } from "./fileTree";

type downloadStartRes = {
	objectKey: string;
	size: string;
	name: string;
	url: string;
}

export async function DownloadFile(node: FileTreeNode) {
	if (!node) return
	const downloadMetaData: downloadStartRes = await getUrl(node.nodeId)

	const a = document.createElement("a")
	a.href = downloadMetaData.url
	a.download = ""
	document.body.append(a)
	a.click()
	a.remove()
}

async function getUrl(fileId: string): Promise<downloadStartRes> {
	const reqUrl = new URL(`${import.meta.env.VITE_BACKEND_URL}/api/start_download`)
	reqUrl.searchParams.set("fileId", fileId)

	const res = await fetch(reqUrl.toString(), {
		method: "POST",
		credentials: "include",
	})

	if (!res.ok) {
		throw new Error("failed to start download")
	}

	const data: downloadStartRes = await res.json()
	return data
}

//async function getObjectKey(fileId: string): Promise<downloadStartRes> {
//	const reqUrl = new URL(`${import.meta.env.VITE_BACKEND_URL}/api/start_download`)
//	reqUrl.searchParams.set("fileId", fileId)
//
//	const res = await fetch(reqUrl.toString(), {
//		method: "POST",
//		credentials: "include",
//	})
//
//	if (!res.ok) {
//		throw new Error("failed to start download")
//	}
//
//	const data: downloadStartRes = await res.json()
//	return data
//}
//
//async function chunkUrl(objectKey: string, partNumber: number) {
//	const reqUrl = new URL(`${import.meta.env.VITE_BACKEND_URL}/api/get_download_url/parts`)
//	reqUrl.searchParams.set("objectKey", objectKey)
//	reqUrl.searchParams.set("partNumber", partNumber.toString())
//
//	const res = await fetch(reqUrl.toString(), {
//		method: "POST",
//		credentials: "include",
//	})
//
//	if (!res.ok) {
//		throw new Error("failed to download chunk")
//	}
//
//	const data = await res.json()
//	return data.url
//}
