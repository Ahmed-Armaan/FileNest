import { createBrowserRouter } from "react-router"
import { RouterProvider } from "react-router"
import LandingPage from "./pages/landingPage"
import Home from "./pages/Home"

const router = createBrowserRouter([
  {
    path: "/",
    element: <LandingPage />
  },
  {
    path: "/home",
    element: <Home />,
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
