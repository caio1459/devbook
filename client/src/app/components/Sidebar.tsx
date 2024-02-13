import Link from "next/link"
import React from "react"
import { FaAlignLeft, FaCalendar, FaClock, FaFlag, FaPeopleArrows, FaStore, FaUserFriends } from "react-icons/fa"
import { FaBookBookmark } from "react-icons/fa6"
import { IUser } from "../interfaces/IUser"

interface IPropsSidebar {
    user: IUser | undefined
    setUser: (value: React.SetStateAction<IUser | undefined>) => void
}

export const Sidebar: React.FC<IPropsSidebar> = ({ user, setUser }) => {
    return (
        <aside className="p-4 pt-10">
            <nav className="flex flex-col gap-6 text-gray-600 font-semibold">
                <Link href={""} className="flex gap-2 pb-6 items-center">
                    <img
                        className="w-10 h-10 rounded-full"
                        src={user?.image_url ? user.image_url : "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTYNbWxNyDw0PPEQ88-OFj5ySYbAsNUH9gFIyVLMwy0tA&s"}
                    />
                    <span className="font-bold">{user?.name}</span>
                </Link >
                <Link href={""} className="flex gap-3 items-center">
                    <FaUserFriends className="w-6 h-6" />
                    Amigos
                </Link>
                <Link href={""} className="flex gap-3 items-center">
                    <FaAlignLeft className="w-6 h-6" />
                    Feed
                </Link>
                <Link href={""} className="flex gap-3 items-center">
                    <FaPeopleArrows className="w-6 h-6" />
                    Grupos
                </Link>
                <Link href={""} className="flex gap-3 items-center">
                    <FaStore className="w-6 h-6" />
                    Loja
                </Link>
                <Link href={""} className="flex gap-3 items-center">
                    <FaBookBookmark className="w-6 h-6" />
                    Salvos
                </Link>
                <Link href={""} className="flex gap-3 items-center">
                    <FaFlag className="w-6 h-6" />
                    Paginas
                </Link>
                <Link href={""} className="flex gap-3 items-center">
                    <FaCalendar className="w-6 h-6" />
                    Eventos
                </Link>
                <Link href={""} className="flex gap-3 items-center">
                    <FaClock className="w-6 h-6" />
                    Lembraças
                </Link>
            </nav>
        </aside>
    )
}