import { IComment } from "@/interfaces/IComment"
import React from "react"
import { UserImage } from "./UserImage"
import moment from "moment"
import 'moment/locale/pt-br';

interface IPropsCommets {
  comment: IComment
}

export const Comments: React.FC<IPropsCommets> = ({ comment }) => {
  return (
    <div className="mt-5 flex gap-2">
      <UserImage image_url={comment.user_image} />
      <div className="text-zinc-600 w-full">
        <div className="flex flex-col bg-zinc-100 px-4 py-1 rounded-md">
          <span className="font-semibold">{comment.actor_nick}</span>
          <span>{comment.text}</span>
        </div>
        <span className="text-xs">{moment(comment.register).fromNow()}</span>
      </div>
    </div>
  )
}