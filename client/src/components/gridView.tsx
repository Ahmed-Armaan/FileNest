import { useState, useEffect, useRef } from "react";
import { NodeTypeDirectory, NodeTypeFile, type FileTreeNode } from "../utils/fileTree";
import {
	FiFolder,
	FiFile,
	FiDownload,
	FiShare2,
	FiTrash2,
} from "react-icons/fi";
import type { TakenPath } from "./FileGrid";
import { DownloadFile } from "../utils/download";

interface GridViewProps {
	fileTree: FileTreeNode;
	AlterFileNode: (node: FileTreeNode) => void;
	currPathAdd: (nextPath: TakenPath) => void;
}

function GridView(props: GridViewProps) {
	const [menuNode, setMenuNode] = useState<FileTreeNode | null>(null);
	const [menuPos, setMenuPos] = useState<{ x: number; y: number } | null>(null);

	const [showShare, setShowShare] = useState(false);
	const [sharePassword, setSharePassword] = useState("");

	const justOpenedRef = useRef(false);

	const handleDirectoryClick = (node: FileTreeNode) => {
		props.AlterFileNode(node);
		props.currPathAdd({ dirName: node.nodeName, TreeNode: node });
	};

	const handleContextMenu = (e: React.MouseEvent, node: FileTreeNode) => {
		e.preventDefault();

		justOpenedRef.current = true;

		setMenuNode(node);
		setMenuPos({ x: e.clientX, y: e.clientY });
		setShowShare(false);
		setSharePassword("");
	};

	const closeMenu = () => {
		setMenuNode(null);
		setMenuPos(null);
		setShowShare(false);
		setSharePassword("");
	};

	useEffect(() => {
		if (!menuNode) return;

		const handleMouseDown = () => {
			if (justOpenedRef.current) {
				justOpenedRef.current = false;
				return;
			}
			closeMenu();
		};

		document.addEventListener("mousedown", handleMouseDown);
		return () => document.removeEventListener("mousedown", handleMouseDown);
	}, [menuNode]);

	const handleShare = () => {
		console.log("share", {
			nodeId: menuNode?.nodeId,
			password: sharePassword || null,
		});
		closeMenu();
	};

	return (
		<div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-4 p-4">
			{props.fileTree.children.map((node) => (
				<div
					key={node.nodeId}
					onClick={() =>
						node.nodeType === NodeTypeDirectory && handleDirectoryClick(node)
					}
					onContextMenu={(e) => handleContextMenu(e, node)}
					className="flex flex-col items-center justify-center gap-2 rounded-lg border border-[var(--border-primary)] bg-[var(--bg-secondary)] p-4 hover:bg-[var(--bg-tertiary)] transition cursor-pointer"
				>
					{node.nodeType === NodeTypeDirectory ? (
						<div className="text-3xl text-[var(--accent-primary)]">
							<FiFolder />
						</div>
					) : (
						<div className="text-3xl text-[var(--text-secondary)]">
							<FiFile />
						</div>
					)}

					<div className="text-sm text-center w-full break-words">
						{node.nodeName}
					</div>
				</div>
			))}

			{menuNode && menuPos && (
				<div
					style={{ top: menuPos.y, left: menuPos.x }}
					className="fixed z-50 w-48 rounded-md border border-[var(--border-primary)] bg-[var(--bg-secondary)] shadow-lg"
					onMouseDown={(e) => e.stopPropagation()}
				>
					{!showShare && (
						<>
							{menuNode.nodeType === NodeTypeFile && (
								<button className="flex w-full items-center gap-2 px-3 py-2 text-sm hover:bg-[var(--bg-tertiary)]"
									onClick={() => DownloadFile(menuNode)}>
									<FiDownload /> Download
								</button>
							)}

							<button
								onClick={() => setShowShare(true)}
								className="flex w-full items-center gap-2 px-3 py-2 text-sm hover:bg-[var(--bg-tertiary)]"
							>
								<FiShare2 /> Share
							</button>

							<button className="flex w-full items-center gap-2 px-3 py-2 text-sm text-red-500 hover:bg-red-500/10">
								<FiTrash2 /> Delete
							</button>
						</>
					)}

					{showShare && (
						<div className="p-3 flex flex-col gap-2">
							<input
								type="password"
								placeholder="Password (optional)"
								value={sharePassword}
								onChange={(e) => setSharePassword(e.target.value)}
								className="w-full rounded border border-[var(--border-primary)] bg-[var(--bg-primary)] px-2 py-1 text-sm"
							/>

							<button
								onClick={handleShare}
								className="w-full rounded bg-[var(--accent-primary)] px-2 py-1 text-sm text-white hover:opacity-90"
							>
								Create link
							</button>

							<button
								onClick={() => setShowShare(false)}
								className="text-xs text-center text-[var(--text-secondary)] hover:underline"
							>
								Back
							</button>
						</div>
					)}
				</div>
			)}
		</div>
	);
}

export default GridView;
