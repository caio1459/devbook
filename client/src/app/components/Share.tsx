import { IUser } from "@/interfaces/IUser"
import { UserImage } from "./UserImage"
import { FaPaperPlane, FaUserFriends } from "react-icons/fa"
import { TbPhoto } from "react-icons/tb"
import { ContainerButtons } from "./containers/ContainerButtons"
import { ActionsButtons } from "./buttons/ActionsButtons"
import { useMutation, useQueryClient } from "react-query"
import AxiosService from "@/services/AxiosService"
import { useContext, useState } from "react"
import { AuthContext } from "@/contexts/AuthContext"
import { GiBrain } from "react-icons/gi"
import { LuSubtitles } from "react-icons/lu"
import { CustomInput } from "./inputs/CustomInput"

interface IPropsShare {
  user: IUser | undefined
}

export const Share: React.FC<IPropsShare> = ({ user }) => {
  const queryClient = useQueryClient()
  const { headers } = useContext(AuthContext)
  const [title, setTitle] = useState("")
  const [content, setContent] = useState("")
  const [image_url, setImageUrl] = useState("")

  const mutation = useMutation({
    mutationFn: async (newValue: {}) => {
      await AxiosService.makeRequest().post("/publications", { title, content, image_url }, headers).then(
        res => res.data
      )
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["publications"] })
  })

  return (
    <div className="sm:w-1/3 bg-white rounded-lg p-4 shadow-md flex flex-col gap-3">
      {image_url && <img src={image_url} className="rounded-full" alt="Imagem da Publicação" />}
      <div className="flex gap-4 pt-6 flex-col">
        <div className="flex flex-row items-center space-x-2">
          <UserImage image_url={user?.image_url} />
          <span className="font-semibold text-xl text-gray-600">Criar publicação</span>
        </div>

        <CustomInput
          placeholder={"Titulo da Publicação"}
          value={title}
          onChange={(newValue => setTitle(newValue))}
        >
          <LuSubtitles />
        </CustomInput>

        <CustomInput
          placeholder={`No que está pensando ${user?.name}?`}
          value={content}
          onChange={(newValue => setContent(newValue))}
        >
          <GiBrain />
        </CustomInput>

        <CustomInput
          placeholder={"URL da imagem"}
          value={image_url}
          onChange={(newValue => setImageUrl(newValue))}
        >
          <TbPhoto />
        </CustomInput>
      </div>

      <ContainerButtons>
        <ActionsButtons text={"Marcar amigos"}>
          <FaUserFriends />
        </ActionsButtons>

        <ActionsButtons
          text={"Publicar"}
          onClick={() => {
            mutation.mutate({ title, content, image_url })
            setTitle("")
            setContent("")
            setImageUrl("")
          }}
        >
          <FaPaperPlane />
        </ActionsButtons>
      </ContainerButtons>
    </div>
  )
}