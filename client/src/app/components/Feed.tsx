import { useEffect, useState } from "react";
import { Publication } from "./Publication";
import { IPublication } from "../interfaces/IPublication";
import AxiosService from "../services/AxiosService";

export const Feed = () => {
    const [publications, setPublications] = useState<IPublication[] | undefined>(undefined)

    useEffect(() => {
        const storedToken = localStorage.getItem("devbook:token")
        //Formata o token para retirar as aspas
        const token = storedToken?.replace(/^"(.*)"$/, '$1');
        //Define o cabeçario da requisição
        const headers = {
            "Authorization": `Bearer ${token}`,
        };
        AxiosService.makeRequest().get("/publications", { headers }).then(res => {
            setPublications(res.data)
        }).catch(err => {
            console.log(err)
        })
    }, [])

    return (
        <div className="flex flex-col items-center gap-5 mt-4 w-full mb-4">
            {
                publications?.map((publication, i) =>
                    (<Publication publication={publication} key={i} />)
                )
            }
        </div>
    )
}