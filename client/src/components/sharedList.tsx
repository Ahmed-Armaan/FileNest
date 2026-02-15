import { useState } from "react"
import {
	FiFolder,
	FiFileText,
	FiChevronDown,
	FiChevronRight,
	FiUsers,
	FiCopy,
	FiX,
} from "react-icons/fi"
import { useNotification } from "../context/notificationContext";
import { NotificationTypes } from "./notificationPanel";

type SharedItem = {
	id: string
	name: string
	type: "file" | "folder"
	code: string
}

export function SharedList() {
	const [open, setOpen] = useState(false)
	const [sharedItems, setSharedItems] = useState<SharedItem[]>([])
	const showNotification = useNotification()

	const handleOpenClick = () => {
		if (open) {
			setOpen(false)
		} else {
			getSharedItems()
		}
	}

	const getSharedItems = async () => {
		try {
			const res = await fetch(
				`${import.meta.env.VITE_BACKEND_URL}/api/all_shared`,
				{ credentials: "include" }
			)

			if (!res.ok) {
				throw new Error("Failed to get shared nodes")
			}

			const data: SharedItem[] = await res.json()
			setSharedItems(data)
		} catch (err) {
			console.error(err)
		} finally {
			setOpen(true)
		}
	}

	const redirectToShared = (code: string) => {
		window.location.href = `/share/${code}`
	}

	const handleCopyLink = async (code: string) => {
		const url = `${window.location.origin}/share/${code}`
		try {
			await navigator.clipboard.writeText(url).
				then(() => {
					showNotification({
						message: 'Share link copied to clipboard',
						notificationType: NotificationTypes.sharedConfirmation,
					})
				})
		} catch (err) {
			console.error("Failed to copy link", err)
			showNotification({
				message: 'Failed to copy link',
				notificationType: NotificationTypes.error,
			})
		}
	}

	//// need to do this
	const handleStopSharing = async (id: string) => {
		try {
			const res = await fetch(
				`${import.meta.env.VITE_BACKEND_URL}/api/remove_shared?nodeId=${id}`,
				{
					method: "GET",
					credentials: "include",
				}
			)

			if (!res.ok) {
				throw new Error("Failed to stop sharing")
			}

			setSharedItems((items) =>
				items.filter((item) => item.id !== id)
			)

			showNotification({
				message: "Sharing stopped",
				notificationType: NotificationTypes.sharedConfirmation,
			})
		} catch (err) {
			console.error("Failed to stop sharing", err)

			showNotification({
				message: "Failed to stop sharing",
				notificationType: NotificationTypes.error,
			})
		}
	}

	return (
		<div className="flex flex-col gap-2">
			<button
				onClick={handleOpenClick}
				className="w-full flex items-center justify-between px-3 py-2 border rounded-md hover:opacity-80 transition"
			>
				<div className="flex items-center gap-2">
					<FiUsers />
					<span>Shared</span>
				</div>

				{open ? <FiChevronDown /> : <FiChevronRight />}
			</button>

			{open && (
				<div className="flex flex-col gap-1 pl-2">
					{sharedItems && sharedItems.map(item => (
						<div
							key={item.id}
							className="flex items-center justify-between gap-3 px-3 py-2 rounded-md hover:bg-[var(--bg-tertiary)] cursor-pointer transition"
							onClick={() => redirectToShared(item.code)}
						>
							<div className="flex items-center gap-3 overflow-hidden">
								{item.type === "folder" ? (
									<FiFolder className="text-[var(--accent-primary)]" />
								) : (
									<FiFileText className="text-[var(--text-secondary)]" />
								)}

								<span className="truncate text-sm">
									{item.name}
								</span>
							</div>

							<div className="flex items-center gap-1">
								<button
									onClick={(e) => {
										e.stopPropagation()
										handleCopyLink(item.code)
									}}
									title="Copy link"
									className="
										p-2 rounded-md
										bg-[var(--bg-secondary)]
										text-[var(--text-secondary)]
										hover:bg-[var(--accent-primary)]
										hover:text-white
										hover:shadow-sm
										transition duration-150
										cursor-pointer
									"
								>
									<FiCopy />
								</button>

								<button
									onClick={(e) => {
										e.stopPropagation()
										handleStopSharing(item.id)
									}}
									title="Stop sharing"
									className="
										p-2 rounded-md
										bg-[var(--bg-secondary)]
										text-[var(--text-secondary)]
										hover:bg-red-500
										hover:text-white
										hover:shadow-sm
										transition duration-150
										cursor-pointer
									"
								>
									<FiX />
								</button>
							</div>
						</div>
					))}

					{sharedItems === null && (
						<div className="px-3 py-2 text-sm text-[var(--text-secondary)]">
							No shared items
						</div>
					)}
				</div>
			)}
		</div>
	)
}
