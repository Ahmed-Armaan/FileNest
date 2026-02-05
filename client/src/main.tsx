import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { UserInfoContextProvider } from './context/userInfoContext.tsx'
import { NotificationContextProvider } from './context/notificationContext.tsx'
import './index.css'
import App from './App.tsx'

createRoot(document.getElementById('root')!).render(
	<NotificationContextProvider>
		<UserInfoContextProvider>
			{/*	<StrictMode>
			</StrictMode>*/}
			<App />
		</UserInfoContextProvider>
	</NotificationContextProvider>
)
