import { useEffect, useState } from "react"
import type { DirectoryMetaData } from "../types/types"
import SideBar from "./sideBar"
import BreadCrumbs from "./BreadCrumbs"
import { FileTreeNode } from "../utils/fileTree"
import { useFileRefreshContext } from "../context/filesRefreshContext"
import { FileUploadContextProvider } from "../context/FileUploadContext"
import { getCookie } from "../utils/cookieExtractor"

export type TakenPath = {
	dirName: string
	TreeNode: FileTreeNode
}

function FileGrid() {
	//const [currElements, setCurrElements] = useState<DirectoryMetaData[]>([])

	//const rootPath: DirectoryMetaData = {
	//	id: "",
	//	name: "",
	//	type: "directory",
	//	updatedAt: "",
	//}
	//const [currPath, setCurrPath] = useState<DirectoryMetaData[]>([rootPath])


	//const currPathBack = () => {
	//	setCurrPath(currPath => {
	//		currPath.pop()
	//		return currPath
	//	})
	//}

	//  const currPathAdd = (nextDir: FileMetaData) => {
	//    setCurrPath(currPath => {
	//      currPath = [...currPath, nextDir]
	//      return currPath
	//    })
	//  }

	//const currPathSet = (newDirId: string) => {
	//	var newPath: DirectoryMetaData[] = []

	//	setCurrPath(currPath => {
	//		for (const dir of currPath) {
	//			newPath = [...newPath, dir]
	//			if (dir.id === newDirId) break
	//		}
	//		return newPath
	//	})
	//}


	const [FileTree, AlterFileTree] = useState<FileTreeNode>()
	const [currPath, setCurrPath] = useState<TakenPath[]>([])
	const { fileRefreshTrigger } = useFileRefreshContext()

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

	// fetch root dir data
	useEffect(() => {
		const rootId = getCookie("rootNodeId")
		const updated_at = getCookie("rootNodeUpdatedAt")

		if (rootId && updated_at) {
			const fileTree = new FileTreeNode(rootId, "/", "directory", updated_at, [], null)
			AlterFileTree(fileTree)
			if (FileTree)
				setCurrPath([{ dirName: "/", TreeNode: FileTree, }])
		}
	}, [])

	//fetch children on change
	useEffect(() => {
		if (FileTree)
			fetchChildElements(FileTree)
	}, [FileTree, fileRefreshTrigger])

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
			node.children = children
		}
		catch (err) {
			console.log(err)
		}
	}

	return (
		<FileUploadContextProvider>
			<div className="max-h-screen flex flex-row border grow">

				{/* files */}
				<div className="flex-4">
					<div>
						<BreadCrumbs currPath={currPath} currPathBack={currPathBack} currPathSet={currPathSet} />
					</div>

					<div>
						<ul>
							{
								FileTree?.children.map((ele) => {
									return <li key={ele.nodeId}>{`${ele.nodeType} - ${ele.nodeName}`}</li>
								})
							}
						</ul>
					</div>
				</div>

				{/* side bar */}
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
