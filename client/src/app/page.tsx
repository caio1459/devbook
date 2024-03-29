"use client";
import { useEffect, useState } from "react";
import { Header } from "./components/Header"
import { useRouter } from "next/navigation";
import { Sidebar } from "./components/Sidebar";
import { IUser } from "../interfaces/IUser";
import { Feed } from "./components/Feed";
import { AuthContextProvider } from "@/contexts/AuthContext";

export default function Home() {
  const router = useRouter()
  const [user, setUser] = useState<IUser | undefined>(undefined)

  useEffect(() => {
    let userStorage = localStorage.getItem("devbook:user")
    if (userStorage) setUser(JSON.parse(userStorage))

    let token = localStorage.getItem("devbook:token")
    if (!token) router.push("/login")
  }, [])

  return (
    <AuthContextProvider>
      <main className="flex min-h-screen flex-col items-center bg-zinc-200">
        <Header user={user} />
        <div className="w-full flex justify-start">
          <Sidebar user={user} />
          <Feed user={user} />
        </div>
      </main>
    </AuthContextProvider>
  );
}
