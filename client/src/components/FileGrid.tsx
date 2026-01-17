import { useEffect, useState } from "react"
import type { FileMetaData } from "../types/types"
import SideBar from "./sideBar"
import BreadCrumbs from "./BreadCrumbs"
import { FileUploadContextProvider } from "../context/FileUploadContext"

function FileGrid() {
	const [currElements, setCurrElements] = useState<FileMetaData[]>([])
	const rootPath: FileMetaData = {
		id: "",
		name: "",
		type: "directory",
		updatedAt: "",
	}
	const [currPath, setCurrPath] = useState<FileMetaData[]>([rootPath])

	const currPathBack = () => {
		setCurrPath(currPath => {
			currPath.pop()
			return currPath
		})
	}

	//  const currPathAdd = (nextDir: FileMetaData) => {
	//    setCurrPath(currPath => {
	//      currPath = [...currPath, nextDir]
	//      return currPath
	//    })
	//  }

	const currPathSet = (newDirId: string) => {
		var newPath: FileMetaData[] = []

		setCurrPath(currPath => {
			for (const dir of currPath) {
				newPath = [...newPath, dir]
				if (dir.id === newDirId) break
			}
			return newPath
		})
	}

	useEffect(() => {
		fetchRootElements()
	}, [currPath])

	const fetchRootElements = async () => {
		try {
			const reqUrl = new URL(`${import.meta.env.VITE_OAUTH_REDIRECT_URI}/api/get_elements`)
			reqUrl.searchParams.append("parentId", currPath[currPath.length - 1].id)

			const res = await fetch(reqUrl.toString(), {
				credentials: "include",
			})

			if (!res.ok) {
				// show error!!
			}

			const data: FileMetaData[] = await res.json()
			setCurrElements(data)
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
								currElements.map((ele) => {
									return <li key={ele.id}>{`${ele.type} - ${ele.name}`}</li>
								})
							}
						</ul>
					</div>
				</div>

				{/* side bar */}
				<div className="flex-1">
					<SideBar currDirId={currPath[currPath.length - 1].id} />
				</div>
			</div>
		</FileUploadContextProvider>
	)
}

export default FileGrid
