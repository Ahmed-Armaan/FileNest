import { useEffect, useState } from "react"
import type { FileMEtaData } from "../types/types"
import SideBar from "./sideBar"
import BreadCrumbs from "./BreadCrumbs"

function FileGrid() {
  const [currElements, setCurrElements] = useState<FileMEtaData[]>([])
  const rootPath: FileMEtaData = {
    id: "",
    name: "",
    type: "directory",
    updatedAt: "",
  }
  const [currPath, setCurrPath] = useState<FileMEtaData[]>([rootPath])

  const currPathBack = () => {
    setCurrPath(currPath => {
      currPath.pop()
      return currPath
    })
  }

  //  const currPathAdd = (nextDir: FileMEtaData) => {
  //    setCurrPath(currPath => {
  //      currPath = [...currPath, nextDir]
  //      return currPath
  //    })
  //  }

  const currPathSet = (newDirId: string) => {
    var newPath: FileMEtaData[] = []

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

      const data: FileMEtaData[] = await res.json()
      setCurrElements(data)
    }
    catch (err) {
      console.log(err)
    }

  }

  return (
    <>
      <div className="flex flex-row m-5 border-1 grow">
        {/* files */}
        <div className="w-4/5">
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
        <div className="w-1/5">
          <SideBar currDirId={currPath[currPath.length - 1].id} />
        </div>
      </div>
    </>
  )
}

export default FileGrid
