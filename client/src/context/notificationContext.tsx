import { createContext, useContext, useState, useRef, type ReactNode } from "react";
import type { NotificationProps } from "../components/notificationPanel";
import Notification from "../components/notificationPanel";

type NotificationWithId = NotificationProps & { id: number };

const NotificationContext = createContext<
	((props: NotificationProps) => void) | undefined
>(undefined);

export function NotificationContextProvider({
	children,
}: {
	children: ReactNode;
}) {
	const [notifications, setNotifications] = useState<NotificationWithId[]>([]);
	const idCounter = useRef(0);

	const showNotification = (props: NotificationProps) => {
		const id = idCounter.current++;

		const newNotification: NotificationWithId = {
			...props,
			id,
		};

		setNotifications((prev) => [...prev, newNotification]);

		setTimeout(() => {
			setNotifications((prev) => prev.filter((n) => n.id !== id));
		}, 3000);
	};

	return (
		<NotificationContext.Provider value={showNotification}>
			{children}

			<div className="fixed top-4 right-4 flex flex-col gap-2 z-50">
				{notifications.map((n) => (
					<Notification
						key={n.id}
						message={n.message}
						notificationType={n.notificationType}
					/>
				))}
			</div>
		</NotificationContext.Provider>
	);
}

export function useNotification() {
	const ctx = useContext(NotificationContext);
	if (!ctx) {
		throw new Error("useNotification must be used inside provider");
	}
	return ctx;
}
