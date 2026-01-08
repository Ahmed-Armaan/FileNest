import { useEffect } from "react"
import { useNavigate } from "react-router"
import { useUserInfoContext } from "../context/userInfoContext"
import Header from "../components/header"

// TODO: fetch the file tree, create context for breadcrums and the UI

function Home() {
  const navigate = useNavigate()
  const { setUser, clearUser } = useUserInfoContext()

  useEffect(() => {
    const loadUser = async () => {
      try {
        const res = await fetch(
          `${import.meta.env.VITE_BACKEND_URL}/me`,
          { credentials: "include" }
        )

        if (!res.ok) {
          clearUser()
          navigate("/")
          return
        }

        const user = await res.json()
        setUser(user.user, user.profile, user.email)
      } catch {
        clearUser()
        navigate("/")
      }
    }

    loadUser()
  }, [])

  return (
    <div className="theme_applier">
      <Header />
      <h1>Welcome</h1>
    </div>
  )
}

export default Home
