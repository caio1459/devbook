"use client";
import { useCallback, useState } from "react";
import { ContainerBackgroud } from "../../components/ContainerBackgroud"
import { ContainerInputs } from "../../components/ContainerInputs"
import { FormContainer } from "../../components/FormContainer"
import { Input } from "../../components/Input"
import { Button } from "../../components/Button";
import axios from "axios";
import { redirect } from "next/navigation";
import AxiosService from "../../services/AxiosService";

const Register = () => {
    const [name, setName] = useState('')
    const [nick, setNick] = useState('')
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [image_url, setImage_url] = useState('')
    const [error, setError] = useState('')

    const clearInputs = () => {
        setName('')
        setEmail('')
        setImage_url('')
        setNick('')
        setPassword('')
    }

    const createUser = useCallback((e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
        e.preventDefault()
        AxiosService.makeRequest().post("/users", { name, nick, email, password, image_url }).then((res) => {
            console.log(res.data)
        }).catch((err) => {
            setError(err.response.data.erro)
            clearInputs()
        })
    }, [name, nick, email, password, image_url])

    return (
        <ContainerBackgroud>
            <FormContainer title="Registrar-se">
                <ContainerInputs>
                    <Input
                        idLabel={"name"}
                        type={"text"}
                        textLabel={"Nome"}
                        placeholder={"Digite o seu nome"}
                        onChange={(newValue) => setName(newValue)}
                        required={true}
                        value={name} />
                </ContainerInputs>
                <ContainerInputs>
                    <Input
                        idLabel={"email"}
                        type={"email"}
                        textLabel={"Email"}
                        placeholder={"Digite o seu email"}
                        onChange={(newValue) => setEmail(newValue)}
                        required={true}
                        value={email} />
                </ContainerInputs>
                <ContainerInputs>
                    <Input
                        idLabel={"nick"}
                        type={"text"}
                        textLabel={"Apelido"}
                        placeholder={"Digite o seu apelido"}
                        onChange={(newValue) => setNick(newValue)}
                        value={nick} />
                </ContainerInputs>
                <ContainerInputs>
                    <Input
                        idLabel={"senha"}
                        type={"password"}
                        textLabel={"Senha"}
                        placeholder={"Digite a sua senha"}
                        onChange={(newValue) => setPassword(newValue)}
                        required={true}
                        value={password} />
                </ContainerInputs>
                <ContainerInputs>
                    <Input
                        idLabel={"imagem"}
                        type={"text"}
                        textLabel={"Imagem de Perfil"}
                        placeholder="Cole a URL da imagem"
                        onChange={(newValue) => setImage_url(newValue)}
                        value={image_url} />
                </ContainerInputs>
                {error.length > 0 && <span className="text-red-600">* {error}</span>}
                <Button
                    text={"Salvar"}
                    click={(e) => createUser(e)}>
                </Button>
            </FormContainer>
        </ContainerBackgroud>
    )
}

export default Register