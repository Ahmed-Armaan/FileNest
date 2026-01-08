import { useNavigate, useSearchParams } from "react-router"
import Logo from "../components/logo"
import Header from "../components/header"
import GoogleButton from "react-google-button"
import { useEffect, useState } from "react"

function LandingPage() {
  const [loginError, toggleErrorState] = useState(false)
  const navigate = useNavigate()
  const [searchParams] = useSearchParams()

  useEffect(() => {
    const errorMsg = searchParams.get("error")
    if (errorMsg && errorMsg.length && errorMsg.length > 0) {
      toggleErrorState(true)
    }
  }, [searchParams])

  useEffect(() => {

  })

  const loginWithGoogle = () => {
    var url = new URL("https://accounts.google.com/o/oauth2/v2/auth")
    const client_id = import.meta.env.VITE_OAUTH_CLIENT_ID
    const redirect_url = import.meta.env.VITE_OAUTH_REDIRECT_URI
    const scopes = import.meta.env.VITE_OAUTH_SCOPE

    url.searchParams.append("client_id", client_id)
    url.searchParams.append("redirect_uri", redirect_url)
    url.searchParams.append("response_type", "code")
    url.searchParams.append("scope", scopes)
    url.searchParams.append("access_type", "offline")
    url.searchParams.append("prompt", "consent")

    window.location.href = url.toString()
  }

  // TODO: check logged in and redirect

  return (
    <div style={{ background: "var(--bg-primary)", color: "var(--text-primary)" }}>
      <Header />

      <div className="min-h-screen flex items-center justify-center">
        <div
          className="flex flex-col items-center gap-6 px-14 py-12 
          rounded-2xl border hover:shadow-md transition-shadow"
          style={{ borderColor: "var(--border-primary)" }}
        >
          <Logo className="text-6xl" />

          <p className="text-sm text-center max-w-md leading-relaxed">
            Store, organize, access, and share your files through a virtual file system.
          </p>

          <div className="pt-2">
            <GoogleButton label="Continue with Google" onClick={loginWithGoogle} />
          </div>

          {loginError &&
            <div className="text-red-400">
              Could not login, please try again
            </div>
          }
        </div>
      </div>
    </div>
  )
}

export default LandingPage
