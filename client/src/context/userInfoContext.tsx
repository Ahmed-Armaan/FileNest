import { createContext, useContext, useState, type ReactNode } from "react";

interface userInfoContextType {
  loggedIn: Boolean;
  userName: string;
  profilePicture: string;
  email: string,
  setUser: (userName: string, profilePicture: string, email: string) => void;
  clearUser: () => void;
}

const userInfoContext = createContext<userInfoContextType | undefined>(undefined)

export function UserInfoContextProvider({ children }: { children: ReactNode }) {
  const [loggedIn, setLoggedIn] = useState(false)
  const [userName, setUserName] = useState("")
  const [profilePicture, setProfilePicture] = useState("")
  const [email, setEmail] = useState("")

  const setUser = (userName: string, profilePicture: string, email: string) => {
    setLoggedIn(true)
    setUserName(userName)
    setProfilePicture(profilePicture)
    setEmail(email)
  }

  const clearUser = () => {
    setLoggedIn(false)
    setUserName("")
    setProfilePicture("")
    setEmail("")
  }

  return (
    <userInfoContext.Provider value={{
      loggedIn: loggedIn,
      userName: userName,
      profilePicture: profilePicture,
      email: email,
      setUser: setUser,
      clearUser: clearUser,
    }}>
      {children}
    </userInfoContext.Provider>
  )
}

export function useUserInfoContext() {
  const currUserInfoContext = useContext(userInfoContext)
  if (!currUserInfoContext) {
    throw new Error("Context can only be used inside a context Provider")
  }
  return currUserInfoContext
}
