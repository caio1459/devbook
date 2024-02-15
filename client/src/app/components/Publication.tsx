import React from "react"
import { IPublication } from "../interfaces/IPublication"
import { FaPaperPlane, FaRegComment, FaThumbsDown, FaThumbsUp } from "react-icons/fa"
import moment from "moment"
import { UserImage } from "./UserImage"

interface IPropsPublication {
    publication: IPublication | undefined
}

export const Publication: React.FC<IPropsPublication> = ({ publication }) => {
    const formattedDate = moment(publication?.register).format("DD/MM/YYYY HH:mm:ss");

    return (
        <div className="sm:w-1/3 bg-white rounded-lg p-4 shadow-md">
            <header className="flex gap-2 pb-4 border-b items-center">
                <UserImage image_url={publication?.user_image} />
                <div className="flex flex-col">
                    <span className="font-semibold">{publication?.actor_nick}</span>
                    <span className="text-xs">{formattedDate}</span>
                </div>
            </header>
            {publication?.content && (
                <div className="py-4 w-full flex flex-col">
                    <p className="text-sky-900 font-bold text-center">{publication.title}</p>
                    <span>{publication.content}</span>
                </div>
            )}
            {publication?.image_url && (
                <img className="rounded-md" src={publication.image_url} alt="Imagem de publicação" />)}
            <div className="flex justify-between py-4 border-b">
                <div className="flex gap-1 items-center">
                    <span className="bg-blue-700 w-6 h-6 text-white rounded-full flex items-center justify-center text-sm">
                        <FaThumbsUp />
                    </span>
                    {publication?.likes}
                </div>
                <span>5 Comentarios</span>
            </div>
            <div className="flex justify-around py-4 text-gray-600 b-4">
                <button className="flex gap-1 items-center">
                    <FaThumbsUp /> Curtir
                </button>
                <button className="flex gap-1 items-center">
                    <FaThumbsDown /> Descurtir
                </button>
                <button className="flex gap-1 items-center">
                    <FaRegComment /> Comentar
                </button>
            </div>
            <div>
                <div className="flex bg-zinc-200 text-gray-600 px-3 py-1 rounded-full items-center w-full">
                    <input type="text" className="bg-zinc-200 outline-none w-full" placeholder="Comentar" />
                    <FaPaperPlane />
                </div>
            </div>
        </div>
    )
}