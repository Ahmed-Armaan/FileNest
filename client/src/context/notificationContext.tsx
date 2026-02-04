import { createContext, useContext, useState, type ReactNode } from "react";
import type { NotificationProps } from "../components/notificationPanel";
import Notification from "../components/notificationPanel";

const NotificationContext = createContext<((props: NotificationProps) => void) | undefined>(undefined)

export function NotificationContextProvider({ children }: { children: ReactNode }) {
	const [notificationProps, setnotificationProps] = useState<NotificationProps>({ message: "", notificationType: 0 })

	const showNotification = (props: NotificationProps) => {
		setnotificationProps(props)
	}

	return (
		<NotificationContext.Provider
			value={showNotification}>
			{children}
			<Notification message={notificationProps?.message} notificationType={notificationProps?.notificationType} />
		</NotificationContext.Provider>
	)
}

export function useNotification() {
	const ctx = useContext(NotificationContext)
	if (!ctx) {
		throw new Error("useNotification must be used inside provider")
	}
	return ctx
}
