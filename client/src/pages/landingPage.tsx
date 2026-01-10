import { useNavigate, useSearchParams } from "react-router"
import { useEffect, useState } from "react"
import { useUserInfoContext } from "../context/userInfoContext"
import Logo from "../components/logo"
import Header from "../components/header"
import GoogleButton from "react-google-button"

function LandingPage() {
  const [loginError, setLoginError] = useState(false)
  const navigate = useNavigate()
  const [searchParams] = useSearchParams()

  const { clearUser } = useUserInfoContext()

  useEffect(() => {
    const errorMsg = searchParams.get("error")
    if (errorMsg) {
      setLoginError(true)
    }
  }, [searchParams])

  useEffect(() => {
    const checkLogin = async () => {
      try {
        const res = await fetch(
          `${import.meta.env.VITE_BACKEND_URL}/api/me`,
          { credentials: "include" }
        )

        if (res.status === 200) {
          navigate("/home")
        }
      } catch (err) {
        clearUser()
      }
    }

    checkLogin()
  },)

  const loginWithGoogle = () => {
    const url = new URL("https://accounts.google.com/o/oauth2/v2/auth")
    const scope = import.meta.env.VITE_OAUTH_SCOPE
    url.searchParams.append("client_id", import.meta.env.VITE_OAUTH_CLIENT_ID)
    url.searchParams.append("redirect_uri", import.meta.env.VITE_OAUTH_REDIRECT_URI)
    url.searchParams.append("response_type", "code")
    url.searchParams.append("scope", scope)
    url.searchParams.append("access_type", "offline")
    url.searchParams.append("prompt", "consent")

    window.location.href = url.toString()
  }

  return (
    <div className="h-screen flex flex-col theme_applier">
      <Header />

      <div className="flex-1 overflow-hidden flex items-center justify-center">
        <div
          className="flex flex-col items-center gap-6 px-14 py-12 rounded-2xl border hover:shadow-md transition-shadow"
          style={{ borderColor: "var(--border-primary)" }}
        >
          <Logo className="text-6xl" />

          <p className="text-sm text-center max-w-md leading-relaxed">
            Store, organize, access, and share your files through a virtual file system.
          </p>

          <GoogleButton label="Continue with Google" onClick={loginWithGoogle} />

          {loginError && (
            <div className="text-red-400 text-sm">
              Could not login, please try again
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default LandingPage
