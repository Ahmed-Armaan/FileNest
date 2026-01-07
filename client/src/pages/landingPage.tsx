import Logo from "../components/logo"
import Header from "../components/header"
import GoogleButton from "react-google-button"

function LandingPage() {
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
            <GoogleButton label="Continue with Google" />
          </div>
        </div>
      </div>
    </div>
  )
}

export default LandingPage
