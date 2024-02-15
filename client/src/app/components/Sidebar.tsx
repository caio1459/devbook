import Link from "next/link"
import React from "react"
import { FaAlignLeft, FaCalendar, FaClock, FaFlag, FaPeopleArrows, FaStore, FaUserFriends } from "react-icons/fa"
import { FaBookBookmark } from "react-icons/fa6"
import { IUser } from "../interfaces/IUser"
import { MenuItems } from "./MenuItems"
import { UserImage } from "./UserImage"

interface IPropsSidebar {
    user: IUser | undefined
}

const menuItens = [
    { href: "", text: "Amigos", icon: (<FaUserFriends className="w-6 h-6" />) },
    { href: "", text: "Feed", icon: (<FaAlignLeft className="w-6 h-6" />) },
    { href: "", text: "Grupos", icon: (<FaPeopleArrows className="w-6 h-6" />) },
    { href: "", text: "Loja", icon: (<FaStore className="w-6 h-6" />) },
    { href: "", text: "Salvos", icon: (<FaBookBookmark className="w-6 h-6" />) },
    { href: "", text: "Paginas", icon: (<FaFlag className="w-6 h-6" />) },
    { href: "", text: "Eventos", icon: (<FaCalendar className="w-6 h-6" />) },
    { href: "", text: "Lembraças", icon: (<FaClock className="w-6 h-6" />) },
]

export const Sidebar: React.FC<IPropsSidebar> = ({ user }) => {
    return (
        <aside className="p-4 pt-10">
            <nav className="flex flex-col gap-6 text-gray-600 font-semibold">
                <Link href={""} className="flex gap-2 pb-6 items-center">
                    <UserImage image_url={user?.image_url} />
                    <span className="font-bold">{user?.name}</span>
                </Link >
                {
                    menuItens.map((item, i) =>
                        <MenuItems href={item.href} icon={item.icon} text={item.text} key={i} />
                    )
                }
            </nav>
        </aside>
    )
}