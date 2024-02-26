import React, { useContext, useState } from "react"
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
import { IUser } from "@/interfaces/IUser";

interface IPropsPublication {
  publication: IPublication | undefined
  user: IUser | undefined
}

export const Publication: React.FC<IPropsPublication> = ({ publication, user }) => {
  const [text, setText] = useState("")
  const { headers } = useContext(AuthContext)
  const queryClient = useQueryClient()

  const { data, error, isLoading } = useQuery<IComment[] | undefined>({
    queryKey: ["comments", publication?.pub_id],
    queryFn: () => AxiosService.makeRequest().get(`/comments/${publication?.pub_id}`, headers).then(
      res => res.data
    ),
    enabled: !!publication?.pub_id //Só ativa a Query se tiver o parametro id
  })

  const mutation = useMutation({
    mutationFn: async (newValue: {}) => {
      await AxiosService.makeRequest().post(`/comments/${publication?.pub_id}`, { text }, headers).then(
        res => res.data
      )
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["comments", publication?.pub_id] })
  })

  if (error) {
    console.log(error)
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
        <span>{data && data.length > 0 ? `${data.length} Comentarios` : ""}</span>
      </div>
      <ContainerButtons>
        <ActionsButtons text={"Curtir"}>
          <FaThumbsUp />
        </ActionsButtons>
        <ActionsButtons text={"Descurtir"}>
          <FaThumbsDown />
        </ActionsButtons>
        <ActionsButtons text={"Comentar"}>
          <FaRegComment />
        </ActionsButtons>
      </ContainerButtons>
      <div>
        <CustomInput
          placeholder={"Adicionar comentario"}
          onChange={newValue => setText(newValue)}
          value={text}
        >
          <FaPaperPlane />
        </CustomInput>
      </div>
    </div>
  )
}