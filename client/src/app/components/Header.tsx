import Link from "next/link"
import { useRouter } from "next/navigation";
import React, { useCallback, useState } from "react"
import { FaBell, FaSearch } from "react-icons/fa"
import { TbMessageCircle2 } from "react-icons/tb"
import { IUser } from "../interfaces/IUser";

interface IPropsHeader {
    user: IUser | undefined
    setUser: (value: React.SetStateAction<IUser | undefined>) => void
}

export const Header: React.FC<IPropsHeader> = ({ user, setUser }) => {
    const [showMenu, setShowMenu] = useState(false)
    const router = useRouter()

    const logout = useCallback((e: React.MouseEvent<HTMLAnchorElement, MouseEvent>) => {
        e.preventDefault()
        localStorage.removeItem("devbook:token")
        router.push("/login")
    }, [])

    return (
        <header className="w-full bg-white flex justify-between py-2 px-4 items-center shadow-md">
            <Link
                href={"/"}
                className="font-bold text-sky-900 text-lg"
            >
                Devbook
            </Link>
            <div className="flex bg-zinc-200 text-gray-600 px-3 py-1 rounded-full items-center">
                <input
                    type="text"
                    className="bg-zinc-200 outline-none"
                    placeholder="Pesquisar"
                />
                <FaSearch />
            </div>
            <div className="flex gap-2 items-center text-gray-600">
                <div className="flex gap-3">
                    <button className="bg-zinc-200 p-2 rounded-full hover:bg-zinc-300">
                        <TbMessageCircle2 />
                    </button>
                    <button className="bg-zinc-200 p-2 rounded-full hover:bg-zinc-300">
                        <FaBell />
                    </button>
                </div>
                <div className="relative" onMouseLeave={() => setShowMenu(false)}>
                    <button
                        className="flex gap-2 items-center"
                        onClick={() => setShowMenu(!showMenu)}
                    >
                        <img
                            src={user?.image_url ? user.image_url : "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTYNbWxNyDw0PPEQ88-OFj5ySYbAsNUH9gFIyVLMwy0tA&s"}
                            alt="Imagem de perfil"
                            className="w-8 h-8 rounded-full"
                        />
                        <span className="font-bold">{user?.name}</span>
                    </button>
                    {showMenu && (
                        <div className="absolute flex flex-col bg-white p-4 shadow-md rounded-md gap-2 border-t whitespace-nowrap right-[-5px]">
                            <Link href={""} className="border-b hover:text-sky-600">
                                Editar Perfil
                            </Link>
                            <Link
                                href={""}
                                className="border-b hover:text-sky-600"
                                onClick={(e) => logout(e)}
                            >
                                Logout
                            </Link>
                        </div>
                    )}
                </div>
            </div>
        </header>
    )
}