import { useEffect, useState } from "react"

export const NotificationTypes = {
	error: 0,
	sharedConfirmation: 1,
}

export type NotificationProps = {
	message: string
	notificationType: number
}

export default function Notification({ message, notificationType }: NotificationProps) {
	const [notifTimeOut, setNotifTimeOut] = useState(0)
	const [showNotiff, toggleNotiff] = useState(false)

	useEffect(() => {
		switch (notificationType) {
			case NotificationTypes.error:
				setNotifTimeOut(1_000)
				break
			case NotificationTypes.sharedConfirmation:
				setNotifTimeOut(30_000)
				break
			default:
				setNotifTimeOut(5_000)
		}
	}, [notificationType])

	useEffect(() => {
		if (notifTimeOut === 0) return

		toggleNotiff(true)

		const timer = setTimeout(() => {
			toggleNotiff(false)
		}, notifTimeOut)

		return () => clearTimeout(timer)
	}, [notifTimeOut])

	if (!showNotiff) return null

	return (
		<div
			className="
				fixed bottom-6 right-6 z-[9999]
				min-w-[260px] max-w-[360px]
				rounded-xl border
				shadow-xl
			"
			style={{
				background: "var(--bg-tertiary)",
				borderColor: "var(--border-primary)",
				color: "var(--text-primary)",
			}}
		>
			<div className="flex items-start justify-between gap-3 p-4">
				<p
					className="text-sm leading-relaxed"
					style={{ color: "var(--text-primary)" }}
				>
					{message}
				</p>

				<button
					onClick={() => toggleNotiff(false)}
					className="
						text-sm font-medium
						opacity-60 hover:opacity-100
						transition-opacity
					"
					style={{ color: "var(--text-secondary)" }}
					aria-label="Dismiss notification"
				>
					✕
				</button>
			</div>
		</div>
	)
}
