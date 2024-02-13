"use client";
import { useEffect, useState } from "react";
import { Header } from "./components/Header"
import { useRouter } from "next/navigation";
import { Sidebar } from "./components/Sidebar";
import { IUser } from "./interfaces/IUser";

export default function Home() {
  const router = useRouter()
  const [user, setUser] = useState<IUser | undefined>(undefined)

  useEffect(() => {
    let value = localStorage.getItem("devbook:user")
    if (value) {
      setUser(JSON.parse(value))
    }
  }, [])

  useEffect(() => {
    let value = localStorage.getItem("devbook:token")
    if (!value) {
      router.push("/login")
    }
  }, [])

  return (
    <main className="flex min-h-screen flex-col items-center bg-zinc-200">
      <Header user={user} setUser={setUser} />
      <div className="w-full flex justify-start">
        <Sidebar user={user} setUser={setUser} />
      </div>
    </main>
  );
}
