import { createBrowserRouter } from "react-router"
import { RouterProvider } from "react-router"
import LandingPage from "./pages/landingPage"
import Home from "./pages/Home"
import { HomeLoader } from "./utils/homeLoader"

const router = createBrowserRouter([
	{
		path: "/",
		element: <LandingPage />
	},
	{
		path: "/home",
		element: <Home />,
		loader: HomeLoader,
	},
	{
		path: "/share/:share_id",
		element: <Home />,
		loader: HomeLoader,
	},
	/*{
		path: "*",
		element: <LandingPage />
	}*/
])

function App() {
	return (
		<RouterProvider router={router} />
	)
}

export default App
