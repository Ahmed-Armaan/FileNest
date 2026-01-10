import type { FileMEtaData } from "../types/types";
import { IoMdArrowRoundBack } from "react-icons/io";

interface BreadCrumbsPropsType {
  currPath: FileMEtaData[],
  currPathBack: () => void,
  currPathSet: (newDirId: string) => void,
}

function BreadCrumbs(props: BreadCrumbsPropsType) {
  return (
    <div className="flex items-center gap-2 border-b border-gray-300 px-3 py-2">

      {/* back button */}
      <button
        onClick={props.currPathBack}
        className="
          flex items-center justify-center w-9 h-9 rounded-full
          border border-gray-300
          hover:bg-gray-100
          active:scale-95
          transition
        "
      >
        <IoMdArrowRoundBack />
      </button>

      {/* path */}
      <div className="flex items-center gap-1 overflow-x-auto">
        {props.currPath.map((dir, idx) => (
          <span key={dir.id || idx} className="flex items-center gap-1">

            <span
              onClick={() => props.currPathSet(dir.id)}
              className="
                cursor-pointer
                text-sm
                hover:underline
                whitespace-nowrap
              "
            >
              {dir.name || "root"}
            </span>

            {idx !== props.currPath.length - 1 && (
              <span className="text-gray-400">/</span>
            )}
          </span>
        ))}
      </div>
    </div>
  )
}

export default BreadCrumbs
