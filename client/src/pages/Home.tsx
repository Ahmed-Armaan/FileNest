import { useEffect } from "react"
import { useLoaderData, useNavigate } from "react-router"
import { useUserInfoContext } from "../context/userInfoContext"
import Header from "../components/header"
import FileGrid from "../components/FileGrid"
import { FilesRefreshContextprovided } from "../context/filesRefreshContext"
import type { Loaderdata } from "../utils/homeLoader"

function Home() {
	const navigate = useNavigate()
	const { setUser, clearUser } = useUserInfoContext()
	const loaderData = useLoaderData() as Loaderdata

	useEffect(() => {
		console.log(loaderData.mode)
		if (loaderData.mode === "share") return
		loadUser()
	}, [])

	const loadUser = async () => {
		try {
			const res = await fetch(
				`${import.meta.env.VITE_BACKEND_URL}/api/me`,
				{ credentials: "include" }
			)

			if (!res.ok) {
				throw Error("Login failed")
			}

			const user = await res.json()
			setUser(user.user, user.profile, user.email)
		} catch {
			clearUser()
			navigate("/")
		}
	}

	return (
		<FilesRefreshContextprovided>
			<div className="theme_applier h-screen flex flex-col">
				<Header />
				<FileGrid />
			</div>
		</FilesRefreshContextprovided>
	)
}

export default Home
