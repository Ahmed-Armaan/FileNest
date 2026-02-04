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
	// const [TreeRoot, setTreeRoot] = useState<FileTreeNode>()
	const [currPath, setCurrPath] = useState<TakenPath[]>([])
	const [currView, setCurrView] = useState<number>(0)
	const [rootNode, setRootNode] = useState<rootNodeRes | undefined>(undefined)
	const { fileRefreshTrigger } = useFileRefreshContext()
	const loaderData = useLoaderData() as Loaderdata

	const currPathBack = () => {
		if (currPath.length <= 1 || !FileTree || !FileTree.parent) return

		AlterFileTree(FileTree?.parent)
		setCurrPath(prev => prev.slice(0, -1))
	}

	const currPathSet = (target: TakenPath) => {
		const newPath: TakenPath[] = []

		for (const p of currPath) {
			newPath.push(p)
			if (p.dirName === target.dirName) {
				break
			}
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

	// fetch root dir data
	useEffect(() => {
		if (loaderData.mode === "normal") {
			getRootNode()
		} else if (loaderData.mode === "share" && loaderData.code) {
			getSharedRoot(loaderData.code,)
		} else {
			// raise an error notification
		}
	}, [loaderData])

	useEffect(() => {
		const rootId = rootNode?.rootNodeId
		const updated_at = rootNode?.rootNodeUpdatedAt

		if (rootId && updated_at) {
			const fileTree = new FileTreeNode(
				rootId,
				"/",
				"directory",
				updated_at,
				[],
				null
			)

			AlterFileTree(fileTree)
			setCurrPath([{ dirName: "/", TreeNode: fileTree }])
		}
	}, [rootNode])

	//fetch children on change
	useEffect(() => {
		console.log(currPath)
		if (FileTree)
			fetchChildElements(FileTree)
	}, [currPath, fileRefreshTrigger])


	const getRootNode = async () => {
		if (rootNode) return
		try {
			const reqUrl = new URL(`${import.meta.env.VITE_BACKEND_URL}/api/root_node`)
			const res = await fetch(reqUrl.toString(), {
				credentials: "include",
			})

			if (!res.ok) {
				throw new Error("fetch request failed")
			}

			const data: rootNodeRes = await res.json()
			setRootNode(data)
		}
		catch (err) {
			console.log(err)
		}
	}

	//type ShareFetchReq struct {
	//	Code     string `json:"code"`
	//	Password string `json:"password"`
	//}
	const getSharedRoot = async (code: string, password?: string) => {
		try {
			const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/api/root_node`, {
				body: JSON.stringify({
					code,
					...(password && { password }),
				}),
			})

			if (!res.ok) {
				throw new Error("Failed to fetch")
			}

			const data = await res.json()
		}
		catch (err) { }
	}

	//fetch children 
	const fetchChildElements = async (node: FileTreeNode) => {
		try {
			if (!node?.nodeId) {
				throw new Error("file tree is empty")
			}

			const reqUrl = new URL(`${import.meta.env.VITE_BACKEND_URL}/api/get_elements`)
			reqUrl.searchParams.append("parentId", node?.nodeId)

			const res = await fetch(reqUrl.toString(), {
				credentials: "include",
			})

			if (!res.ok) {
				throw new Error("element fetch failed")
			}

			const data: DirectoryMetaData[] = await res.json()
			const children = data.map((ele) => new FileTreeNode(ele.id, ele.name, ele.type, ele.updatedAt, [], node))
			setChildren(children)
		}
		catch (err) {
			console.log(err)
		}
	}

	return (
		<FileUploadContextProvider>
			<div className="max-h-screen flex flex-row border grow">

				<div className="flex-4">
					<div>
						<BreadCrumbs currPath={currPath} currPathBack={currPathBack} currPathSet={currPathSet} setCurrView={setCurrView} />
					</div>

					<div>
						{(FileTree && currView === ViewSelector.gridView) &&
							<GridView fileTree={FileTree} AlterFileNode={AlterFileTree} currPathAdd={currPathAdd} />
						}
						{/*{(TreeRoot && currView === ViewSelector.treeView) &&
							<TreeView treeRoot={TreeRoot} />
						}*/}
					</div>
				</div>

				<div className="flex-1">
					{
						<SideBar currDirId={FileTree?.nodeId || ""} />
					}
				</div>
			</div>
		</FileUploadContextProvider>
	)
}

export default FileGrid
