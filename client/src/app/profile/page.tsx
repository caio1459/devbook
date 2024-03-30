"use client"

import { useContext, useEffect, useState } from "react"
import { Header } from "../components/Header"
import { IUser } from "@/interfaces/IUser"
import router from "next/router"
import { useQuery } from "react-query"
import AxiosService from "@/services/AxiosService"
import { AuthContext } from "@/contexts/AuthContext"


const Profile = ({ searchParams }: { searchParams: { id: string } }) => {
  const [user, setUser] = useState<IUser | undefined>(undefined)
  const { headers } = useContext(AuthContext)

  useEffect(() => {
    let userStorage = localStorage.getItem("devbook:user")
    if (userStorage) setUser(JSON.parse(userStorage))

    let token = localStorage.getItem("devbook:token")
    if (!token) router.push("/login")
  }, [])

  const { data, error } = useQuery({
    queryKey: ["profile", searchParams.id],
    queryFn: () => AxiosService.makeRequest().get(`/users/${searchParams.id}`, headers).then(res => res.data),
  })

  console.log(headers)

  if (data) {
    console.log(data)
  }

  if (error) {
    console.error(error)
  }

  return (
    <>
      <Header user={user} />
      <div>
        {searchParams.id}
      </div>
    </>
  )
}

export default Profile