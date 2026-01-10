import type { FileMEtaData } from "../types/types";
import { IoMdArrowRoundBack } from "react-icons/io";

interface BreadCrumbsPropsType {
  currPath: FileMEtaData[],
  currPathBack: () => void,
  currPathSet: (newDirId: string) => void,
}

function BreadCrumbs(props: BreadCrumbsPropsType) {
  return (
    <>
      <div className="flex gap-2 border-b-1">
        <div className="
          flex items-center justify-center w-9 h-9 rounded-full m-2
          border border-gray-300
          hover:bg-gray-100 hover:text-gray-900
          active:scale-95
          transition
          "
          onClick={() => props.currPathBack}>
          <IoMdArrowRoundBack />
        </div>

        <div>
          {
            props.currPath.map((dir) => {
              return <span id={dir.id} className="underline cursor-pointer"
                onClick={() => props.currPathSet(dir.id)}>
                {`${dir.name}/`}
              </span>
            })
          }
        </div>
      </div>
    </>
  )
}

export default BreadCrumbs
