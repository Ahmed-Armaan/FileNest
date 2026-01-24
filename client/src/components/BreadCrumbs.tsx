import { IoMdArrowRoundBack } from "react-icons/io";
import { IoGrid } from "react-icons/io5";
import { ViewSelector, type TakenPath } from "./FileGrid";

interface BreadCrumbsPropsType {
	currPath: TakenPath[],
	currPathBack: () => void,
	currPathSet: (Node: TakenPath) => void,
	setCurrView: (view: number) => void,
}

function BreadCrumbs(props: BreadCrumbsPropsType) {
	return (
		< div className="flex items-center justify-between gap-2 border-b px-3 py-2">
			<div className="flex gap-2">

				{/* Back button */}
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

				{/* Breadcrumbs */}
				<div className="flex items-center gap-1 overflow-x-auto">
					{props.currPath.map((dir, idx) => (
						<span key={idx} className="flex items-center gap-1">

							<span
								onClick={() => { props.currPathSet(dir) }}
								className="
						cursor-pointer
						text-sm
						hover:underline
						whitespace-nowrap">
								{(dir.dirName === "/") ? "root" : dir.dirName}
							</span>

							{idx !== props.currPath.length - 1 && (
								<span className="text-gray-400">/</span>
							)}
						</span>
					))}
				</div>
			</div>

			{/* View selector buttons */}
			<div className="flex gap-2">
				<button className="border rounded-xl p-2"
					onClick={() => props.setCurrView(ViewSelector.gridView)}>
					<IoGrid />
				</button>
				{/*<div className="border rounded-xl p-2"
					onClick={() => props.setCurrView(ViewSelector.treeView)}>
					<FaFolderTree />
				</div>*/}
			</div>
		</div >
	)
}

export default BreadCrumbs
