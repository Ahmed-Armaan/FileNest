import { createBrowserRouter } from "react-router"
import { RouterProvider } from "react-router"
import LandingPage from "./pages/landingPage"

const router = createBrowserRouter([
  {
    path: "/",
    element: <LandingPage />
  },
  {
    path: "*",
    element: <LandingPage />
  }
])

function App() {
  return (
    <RouterProvider router={router} />
  )
}

export default App
