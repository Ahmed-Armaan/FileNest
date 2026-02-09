import { useState } from "react"
import { FiFolder, FiFileText, FiChevronDown, FiChevronRight, FiUsers } from "react-icons/fi"

type SharedItem = {
	id: string
	name: string
	type: "file" | "folder"
	code: string
}

export function SharedList() {
	const [open, setOpen] = useState(false)
	const [sharedItems, setSharedItems] = useState<SharedItem[]>([])

	const handleOpenClick = () => {
		if (open) setOpen(false)
		else getSharedItems()
	}

	const getSharedItems = async () => {
		try {
			const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/api/all_shared`, {
				credentials: "include",
			})

			if (!res.ok) {
				throw new Error("Failed to get shared nodes")
			}

			const data: SharedItem[] = await res.json()
			setSharedItems(data)
		}
		catch (err) {
			console.log(err)
		}
		finally {
			setOpen(true)
		}
	}

	const redirectToShared = (code: string) => {
		window.location.href = `/share/${code}`
	}

	return (
		<div className="flex flex-col gap-2">
			<button
				onClick={() => handleOpenClick()}
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
					{sharedItems.map(item => (
						<div
							key={item.id}
							className="flex items-center gap-3 px-3 py-2 rounded-md hover:bg-[var(--bg-tertiary)] cursor-pointer transition"
							onClick={() => { redirectToShared(item.code) }}
						>
							{item.type === "folder" ? (
								<FiFolder className="text-[var(--accent-primary)]" />
							) : (
								<FiFileText className="text-[var(--text-secondary)]" />
							)}

							<div className="flex flex-col text-sm overflow-hidden">
								<span className="truncate">{item.name}</span>
							</div>
						</div>
					))}

					{sharedItems.length === 0 && (
						<div className="px-3 py-2 text-sm text-[var(--text-secondary)]">
							No shared items
						</div>
					)}
				</div>
			)}
		</div>
	)
}
