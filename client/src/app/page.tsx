"use client";
import { useEffect } from "react";
import { Header } from "./components/Header"
import { useRouter } from "next/navigation";
import { Sidebar } from "./components/Sidebar";

export default function Home() {
  const router = useRouter()

  useEffect(() => {
    let value = localStorage.getItem("devbook:token")
    if (!value) {
      router.push("/login")
    }
  }, [])

  return (
    <main className="flex min-h-screen flex-col items-center bg-zinc-200">
      <Header />
      <Sidebar />
    </main>
  );
}
