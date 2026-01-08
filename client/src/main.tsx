import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { UserInfoContextProvider } from './context/userInfoContext.tsx'
import './index.css'
import App from './App.tsx'

createRoot(document.getElementById('root')!).render(
  <UserInfoContextProvider>
    <StrictMode>
      <App />
    </StrictMode>
  </UserInfoContextProvider>,
)
