import React, { useContext, useRef, useState } from "react"
import { IPublication } from "../../interfaces/IPublication"
import { FaPaperPlane, FaRegComment, FaThumbsDown, FaThumbsUp } from "react-icons/fa"
import moment from "moment"
import 'moment/locale/pt-br';
import { UserImage } from "./UserImage"
import { ContainerButtons } from "./containers/ContainerButtons"
import { ActionsButtons } from "./buttons/ActionsButtons"
import { CustomInput } from "./inputs/CustomInput";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { IComment } from "@/interfaces/IComment";
import AxiosService from "@/services/AxiosService";
import { AuthContext } from "@/contexts/AuthContext";
import { Comments } from "./Comments";

interface IPropsPublication {
  publication: IPublication
  userImg?: string
  userID?: number
}

export const Publication: React.FC<IPropsPublication> = ({ publication, userImg, userID }) => {
  const [text, setText] = useState("")
  const [showComments, setShowComments] = useState(false)
  const inputRef = useRef<HTMLInputElement>(null);
  const { headers } = useContext(AuthContext)
  const queryClient = useQueryClient()

  const handleCommentClick = () => inputRef.current?.focus()

  const { data, error, isLoading } = useQuery<IComment[] | undefined>({
    queryKey: ["comments", publication?.pub_id],
    queryFn: () => AxiosService.makeRequest().get(`/comments/${publication?.pub_id}`, headers).then(
      res => res.data
    ),
    enabled: !!publication?.pub_id //Só ativa a Query se tiver o parametro id
  })

  if (error) {
    console.log(error)
  }

  const mutation = useMutation({
    mutationFn: async (newValue: {}) => {
      await AxiosService.makeRequest().post(`/comments/${publication?.pub_id}`, { userID, text }, headers)
        .then(res => res.data)
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["comments", publication?.pub_id] })
  })

  const shareComent = async () => {
    mutation.mutate({ userID, text })
    setText("")
  }

  return (
    <div className="sm:w-1/3 bg-white rounded-lg p-4 shadow-md">
      <header className="flex gap-2 pb-4 border-b items-center">
        <UserImage image_url={publication?.user_image} />
        <div className="flex flex-col">
          <span className="font-semibold">{publication?.actor_nick}</span>
          <span className="text-xs">{moment(publication?.register).fromNow()}</span>
        </div>
      </header>
      {publication?.content && (
        <div className="py-4 w-full flex flex-col">
          <p className="text-sky-900 font-bold text-center mb-2 text-xl">{publication.title}</p>
          <span>{publication.content}</span>
        </div>
      )}
      {publication?.image_url && (
        <img className="rounded-md w-full" src={publication.image_url} alt="Imagem de publicação" />)}
      <div className="flex justify-between py-4 border-b">
        <div className="flex gap-1 items-center">
          <span className="bg-blue-700 w-6 h-6 text-white rounded-full flex items-center justify-center text-sm">
            <FaThumbsUp />
          </span>
          {publication?.likes}
        </div>
        <button onClick={() => setShowComments(!showComments)}>
          <span>{data && data.length > 0 ? `${data.length} Comentarios` : ""}</span>
        </button>
      </div>
      <ContainerButtons>
        <ActionsButtons text={"Curtir"}>
          <FaThumbsUp />
        </ActionsButtons>
        <ActionsButtons text={"Descurtir"}>
          <FaThumbsDown />
        </ActionsButtons>
        <ActionsButtons
          text={"Comentar"}
          onClick={handleCommentClick}
        >
          <FaRegComment />
        </ActionsButtons>
      </ContainerButtons>
      {showComments && data?.map((comment, i) => <Comments key={i} comment={comment} />)}
      <div className="flex flex-row gap-2">
        <UserImage image_url={userImg} />
        <CustomInput
          ref={inputRef}
          placeholder={"Adicionar comentario"}
          onChange={newValue => setText(newValue)}
          value={text}
        >
          <button onClick={() => shareComent()}>
            <FaPaperPlane />
          </button>
        </CustomInput>
      </div>
    </div>
  )
}