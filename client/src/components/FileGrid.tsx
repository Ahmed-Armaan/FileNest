import { useEffect, useState } from "react"
import { FileTreeNode } from "../utils/fileTree"
import { useFileRefreshContext } from "../context/filesRefreshContext"
import { FileUploadContextProvider } from "../context/FileUploadContext"
import GridView from "./gridView"
import SideBar from "./sideBar"
import BreadCrumbs from "./BreadCrumbs"
import type { DirectoryMetaData } from "../types/types"
import type { Loaderdata } from "../utils/homeLoader"
import { useLoaderData } from "react-router"
import PasswordPromt from "./passwordPrompt"

export type TakenPath = {
	dirName: string
	TreeNode: FileTreeNode
}

export const ViewSelector = {
	gridView: 0,
	treeView: 1,
}

type rootNodeRes = {
	rootNodeId: string
	rootNodeUpdatedAt: string
}

function FileGrid() {
	const [FileTree, AlterFileTree] = useState<FileTreeNode>()
	const [currPath, setCurrPath] = useState<TakenPath[]>([])
	const [currView, setCurrView] = useState<number>(0)
	const [rootNode, setRootNode] = useState<rootNodeRes>()
	const [passwordInput, ShowPasswordInput] = useState(false)
	const [password, setPassword] = useState("")
	const [sharedRootChildren, setSharedRootChildren] = useState<FileTreeNode[] | null>(null)

	const { fileRefreshTrigger } = useFileRefreshContext()
	const loaderData = useLoaderData() as Loaderdata

	const currPathBack = () => {
		if (currPath.length <= 1 || !FileTree?.parent) return

		AlterFileTree(FileTree.parent)
		setCurrPath(prev => prev.slice(0, -1))
	}

	const currPathSet = (target: TakenPath) => {
		const newPath: TakenPath[] = []

		for (const p of currPath) {
			newPath.push(p)
			if (p.dirName === target.dirName) break
		}

		setCurrPath(newPath)
		AlterFileTree(target.TreeNode)
	}

	const currPathAdd = (nextPath: TakenPath) => {
		setCurrPath(prev => [...prev, nextPath])
	}

	const setChildren = (children: FileTreeNode[]) => {
		AlterFileTree(prev => {
			if (!prev) return prev

			return new FileTreeNode(
				prev.nodeId,
				prev.nodeName,
				prev.nodeType,
				prev.updatedAt,
				children,
				prev.parent
			)
		})
	}

	// initial mode handling
	useEffect(() => {
		const init = async () => {
			console.log(loaderData.mode)
			if (loaderData.mode === "normal") {
				getRootNode()
				return
			}

			if (loaderData.mode === "share" && loaderData.code) {
				const needsPassword = await getPasswordStatus(loaderData.code)
				if (needsPassword === null) {
					ShowPasswordInput(true)
				} else {
					getSharedRoot(loaderData.code)
				}
			}
		}

		init()
	}, [loaderData])

	// password submit for shared
	useEffect(() => {
		if (!password || !loaderData.code) return

		ShowPasswordInput(false)
		getSharedRoot(loaderData.code, password)
	}, [password])

	// construct root node
	useEffect(() => {
		if (!rootNode) return

		const root = new FileTreeNode(
			rootNode.rootNodeId,
			"/",
			"directory",
			rootNode.rootNodeUpdatedAt,
			sharedRootChildren ?? [],
			null
		)

		AlterFileTree(root)
		setCurrPath([{ dirName: "/", TreeNode: root }])
	}, [rootNode])

	// fetch children (normal mode only)
	useEffect(() => {
		if (loaderData.mode === "share") return
		if (!FileTree) return

		fetchChildElements(FileTree)
	}, [currPath, fileRefreshTrigger])

	const getRootNode = async () => {
		if (rootNode) return

		try {
			const res = await fetch(
				`${import.meta.env.VITE_BACKEND_URL}/api/root_node`,
				{ credentials: "include" }
			)

			if (!res.ok) throw new Error("fetch failed")

			const data: rootNodeRes = await res.json()
			setRootNode(data)
		} catch (err) {
			console.log(err)
		}
	}

	const getPasswordStatus = async (code: string): Promise<boolean> => {
		try {
			const reqUrl = new URL(`${import.meta.env.VITE_BACKEND_URL}/shared_password_status`)
			reqUrl.searchParams.set("code", code)

			const res = await fetch(reqUrl.toString())
			if (!res.ok) throw new Error("password status fetch failed")

			const data: { status: boolean } = await res.json()
			return data.status
		} catch {
			return false
		}
	}

	const getSharedRoot = async (code: string, password?: string) => {
		try {
			const res = await fetch(
				`${import.meta.env.VITE_BACKEND_URL}/get_share`,
				{
					method: "POST",
					headers: { "Content-Type": "application/json" },
					body: JSON.stringify({
						code,
						...(password && { password }),
					}),
				}
			)

			if (!res.ok) throw new Error("shared root fetch failed")

			const data: DirectoryMetaData[] = await res.json()

			const children = data.map(
				ele =>
					new FileTreeNode(
						ele.id,
						ele.name,
						ele.type,
						ele.updatedAt,
						[],
						null
					)
			)

			setSharedRootChildren(children)
			setRootNode({
				rootNodeId: code,
				rootNodeUpdatedAt: "shared",
			})
		} catch {
			ShowPasswordInput(true)
		}
	}

	const fetchChildElements = async (node: FileTreeNode) => {
		try {
			const reqUrl = new URL(`${import.meta.env.VITE_BACKEND_URL}/api/get_elements`)
			reqUrl.searchParams.append("parentId", node.nodeId)

			const res = await fetch(reqUrl.toString(), {
				credentials: "include",
			})

			if (!res.ok) throw new Error("fetch children failed")

			const data: DirectoryMetaData[] = await res.json()
			const children = data.map(
				ele =>
					new FileTreeNode(
						ele.id,
						ele.name,
						ele.type,
						ele.updatedAt,
						[],
						node
					)
			)

			setChildren(children)
		} catch (err) {
			console.log(err)
		}
	}

	return (
		<FileUploadContextProvider>
			<div className="max-h-screen flex flex-row border grow">
				{passwordInput && loaderData.code && (
					<PasswordPromt code={loaderData.code} setPassword={setPassword} />
				)}

				<div className="flex-4">
					<BreadCrumbs
						currPath={currPath}
						currPathBack={currPathBack}
						currPathSet={currPathSet}
						setCurrView={setCurrView}
					/>

					{FileTree && currView === ViewSelector.gridView && (
						<GridView
							fileTree={FileTree}
							AlterFileNode={AlterFileTree}
							currPathAdd={currPathAdd}
						/>
					)}
				</div>

				<div className="flex-1">
					<SideBar currDirId={FileTree?.nodeId || ""} />
				</div>
			</div>
		</FileUploadContextProvider>
	)
}

export default FileGrid
