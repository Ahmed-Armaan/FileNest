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
import { useFileRefreshContext } from "../context/filesRefreshContext";
import { useNotification } from "../context/notificationContext";
import { NotificationTypes } from "./notificationPanel";

interface GridViewProps {
	fileTree: FileTreeNode;
	AlterFileNode: (node: FileTreeNode) => void;
	currPathAdd: (nextPath: TakenPath) => void;
}

function GridView(props: GridViewProps) {
	const [menuNode, setMenuNode] = useState<FileTreeNode | null>(null);
	const [menuPos, setMenuPos] = useState<{ x: number; y: number } | null>(null);
	const [showShare, setShowShare] = useState(false);
	const [sharePassword, setSharePassword] = useState("")
	const justOpenedRef = useRef(false);
	const { triggerFilesRefresh } = useFileRefreshContext()
	const showNotification = useNotification()

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

	const handleShare = async () => {
		const nodeId = menuNode?.nodeId
		const password = sharePassword || null

		try {
			const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/api/share`, {
				method: "POST",
				body: JSON.stringify({
					nodeId: nodeId,
					password: password,
				}),
				credentials: "include",
			})

			if (!res.ok) {
				throw new Error("Request failed")
			}

			const data = await res.json()
			const code = data.code
			showNotification({
				message: `The ${menuNode?.nodeType} ${menuNode?.nodeName} is shared at : ${import.meta.env.VITE_FRONTEND_URL}/share/${code}`,
				notificationType: NotificationTypes.sharedConfirmation,
			})
		}
		catch (err) {
			console.log(err)
		}
		finally {
			closeMenu()
		}
	};

	const deleteNode = async (nodeId: string) => {
		try {
			const reqUrl = new URL(`${import.meta.env.VITE_BACKEND_URL}/api/delete`)
			reqUrl.searchParams.set("nodeId", nodeId)

			const res = await fetch(reqUrl.toString(), {
				method: "DELETE",
				credentials: "include",
			})

			if (!res.ok) {
				throw new Error("Failed to delete")
			}

			closeMenu()
			triggerFilesRefresh()
		}
		catch (err) {
			console.log(err)
		}
	}

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

							<button className="flex w-full items-center gap-2 px-3 py-2 text-sm text-red-500 hover:bg-red-500/10"
								onClick={() => deleteNode(menuNode.nodeId)}>
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
