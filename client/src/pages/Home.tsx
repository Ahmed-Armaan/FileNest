import { useEffect } from "react"
import { useNavigate } from "react-router"
import { useUserInfoContext } from "../context/userInfoContext"
import Header from "../components/header"
import FileGrid from "../components/FileGrid"

// TODO: fetch the file tree, create context for breadcrums and the UI

function Home() {
  const navigate = useNavigate()
  const { setUser, clearUser } = useUserInfoContext()

  useEffect(() => {
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
    <div className="theme_applier h-screen flex flex-col">
      <Header />
      <FileGrid />
    </div>
  )
}

export default Home
