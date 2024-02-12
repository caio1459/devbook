import Link from "next/link"
import { FaAlignLeft, FaCalendar, FaFlag, FaPeopleArrows, FaStore, FaUserFriends } from "react-icons/fa"
import { FaBookBookmark } from "react-icons/fa6"

export const Sidebar = () => {

    return (
        <aside>
            <nav>
                <Link href={""}>
                    <img src={""} alt={""} />
                    <span>Usuário</span >
                </Link >
                <Link href={""}>
                    <FaUserFriends />
                    Amigos
                </Link>
                <Link href={""}>
                    <FaAlignLeft />
                    Feed
                </Link>
                <Link href={""}>
                    <FaPeopleArrows />
                    Grupos
                </Link>
                <Link href={""}>
                    <FaStore />
                    Loja
                </Link>
                <Link href={""}>
                    <FaBookBookmark />
                    Salvos
                </Link>
                <Link href={""}>
                    <FaFlag />
                    Paginas
                </Link>
                <Link href={""}>
                    <FaCalendar />
                    Eventos
                </Link>
            </nav>
        </aside>
    )
}