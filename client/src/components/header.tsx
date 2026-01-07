import { useEffect, useState } from "react"
import Logo from "./logo"

function Header() {
  const [theme, setTheme] = useState("")
  const toggleTheme = () => {
    setTheme(prevTheme => {
      const currTheme = (prevTheme === "") ? "light" : ""
      return currTheme
    })
  }

  useEffect(() => {
    document.documentElement.className = theme
  }, [theme])

  return (
    <>
      <div className="sticky top-0 z-1 flex justify-between p-5 min-w-full" style={{ background: "var(--bg-secondary)" }}>
        <Logo className="text-lg" />

        <div className="flex">
          <div onClick={toggleTheme}>SEt Theme</div>
          <div>User</div>
        </div>
      </div>
    </>
  )
}

export default Header
