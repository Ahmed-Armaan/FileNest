import { useEffect, useRef, useState } from "react"
import Logo from "./logo"
import { useUserInfoContext } from "../context/userInfoContext"
import { LuSunDim } from "react-icons/lu"
import { IoMoonOutline } from "react-icons/io5"
import { useNavigate } from "react-router"

function Header() {
  const { loggedIn, profilePicture, userName, email, clearUser } = useUserInfoContext()
  const navigate = useNavigate()

  const [theme, setTheme] = useState("")
  const [open, setOpen] = useState(false)
  const dropdownRef = useRef<HTMLDivElement>(null)

  const toggleTheme = () => {
    setTheme(t => (t === "" ? "light" : ""))
  }

  useEffect(() => {
    document.documentElement.className = theme
  }, [theme])

  useEffect(() => {
    const handler = (e: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(e.target as Node)) {
        setOpen(false)
      }
    }
    document.addEventListener("mousedown", handler)
    return () => document.removeEventListener("mousedown", handler)
  }, [])

  const logout = async () => {
    clearUser()
    navigate("/")
  }

  return (
    <div
      className="sticky top-0 z-10 flex justify-between p-5 min-w-full"
      style={{ background: "var(--bg-secondary)" }}
    >
      <Logo className="text-lg" />

      <div className="flex items-center gap-4 relative">
        {/* Theme toggle */}
        <div
          className="w-10 h-10 rounded-full border flex items-center justify-center cursor-pointer"
          style={{ borderColor: "var(--border-primary)" }}
          onClick={toggleTheme}
        >
          {theme === "light" ? <LuSunDim /> : <IoMoonOutline />}
        </div>

        {/* Profile */}
        {loggedIn && (
          <div ref={dropdownRef} className="relative">
            <div
              className="w-10 h-10 rounded-full border overflow-hidden cursor-pointer"
              style={{ borderColor: "var(--border-primary)" }}
              onClick={() => setOpen(o => !o)}
            >
              <img
                src={profilePicture}
                className="w-full h-full object-cover"
                referrerPolicy="no-referrer"
              />
            </div>

            {/* Dropdown */}
            {open && (
              <div
                className="absolute right-0 mt-3 w-56 rounded-xl border shadow-md p-4"
                style={{
                  background: "var(--bg-primary)",
                  borderColor: "var(--border-primary)",
                }}
              >
                <div className="flex items-center gap-3 mb-3">
                  <img
                    src={profilePicture}
                    className="w-10 h-10 rounded-full"
                    referrerPolicy="no-referrer"
                  />
                  <div className="text-sm">
                    <div className="font-medium">{userName}</div>
                    <div className="text-xs opacity-70">{email}</div>
                  </div>
                </div>

                <button
                  onClick={logout}
                  className="w-full text-sm py-2 rounded-md hover:bg-red-500 hover:text-white transition"
                >
                  Logout
                </button>
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  )
}

export default Header
