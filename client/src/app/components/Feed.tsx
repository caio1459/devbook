import React, { useContext, useEffect, useState } from "react";
import { Publication } from "./Publication";
import { IPublication } from "../../interfaces/IPublication";
import AxiosService from "../../services/AxiosService";
import { Share } from "./Share";
import { IUser } from "@/interfaces/IUser";
import { useQuery } from "react-query";
import { AuthContext } from "@/contexts/AuthContext";

interface IPropsFeed {
  user: IUser | undefined
}

export const Feed: React.FC<IPropsFeed> = ({ user }) => {
  // const [user, setUser] = useState<IUser | undefined>(undefined)
  const { headers } = useContext(AuthContext)

  const { data, isLoading, error } = useQuery<IPublication[] | undefined>({
    queryKey: ["publications"],
    queryFn: () => AxiosService.makeRequest().get("/publications", headers).then(res => res.data)
  })

  if (error) {
    console.log(error)
  }

  return (
    <div className="flex flex-col items-center gap-5 mt-4 w-full mb-4">
      <Share user={user} />
      {isLoading ? <span>Carregando...</span> :
        (
          data?.map((publication, i) => {
            return (
              <div className="w-full flex flex-col items-center" key={i}>
                <Publication
                  publication={publication}
                  userImg={user?.image_url}
                  userID={user?.user_id}
                />
              </div>
            )
          })
        )
      }
    </div>
  )
}