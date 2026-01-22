import { IoMdArrowRoundBack } from "react-icons/io";
import type { TakenPath } from "./FileGrid";

interface BreadCrumbsPropsType {
	currPath: TakenPath[],
	currPathBack: () => void,
	currPathSet: (Node: TakenPath) => void,
}

function BreadCrumbs(props: BreadCrumbsPropsType) {
	return (
		<div className="flex items-center gap-2 border-b px-3 py-2">

			{/* back button */}
			<button
				onClick={props.currPathBack}
				className="
				flex items-center justify-center w-9 h-9 rounded-full
				border border-gray-300
				hover:bg-gray-100
				active:scale-95
				transition">
				<IoMdArrowRoundBack />
			</button>

			{/* path */}
			<div className="flex items-center gap-1 overflow-x-auto">
				{props.currPath.map((dir, idx) => (
					<span key={idx} className="flex items-center gap-1">

						<span
							onClick={() => { }}
							className="
						cursor-pointer
						text-sm
						hover:underline
						whitespace-nowrap">
							{dir.dirName || "root"}
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
