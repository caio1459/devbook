"use client";
import Link from "next/link"
import { Input } from "../components/Input"
import { ContainerInputs } from "../components/ContainerInputs"
import { Button } from "../components/Button"
import { useCallback, useState } from "react"
import axios from "axios";
import { ContainerBackgroud } from "../components/ContainerBackgroud";
import { FormContainer } from "../components/FormContainer";
import { redirect } from "next/navigation";

const Login = () => {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [error, setError] = useState('')

    const clearInputs = () => {
        setEmail('')
        setPassword('')
    }

    const handleLogin = useCallback((e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
        e.preventDefault()
        axios.post("http://localhost:5000/api/login", { email, password }).then(resp => {
            redirect("/home")
        }).catch(err => {
            setError(err.response.data.erro)
            clearInputs()
        })
    }, [email, password])

    return (
        <ContainerBackgroud>
            <FormContainer title={"LOGIN"}>
                <ContainerInputs>
                    <Input
                        idLabel={"email"}
                        type={"email"}
                        textLabel={"Email"}
                        placeholder={"Digite o seu Email"}
                        required={true}
                        onChange={(newValue) => setEmail(newValue)}
                        value={email} />
                </ContainerInputs>
                <ContainerInputs>
                    <Input
                        idLabel={"senha"}
                        type={"password"}
                        textLabel={"Senha"}
                        placeholder={"Digite a sua Senha"}
                        required={true}
                        onChange={(newValue) => setPassword(newValue)}
                        value={password} />
                </ContainerInputs>
                {error.length > 0 && <span className="text-red-600">* {error}</span>}
                <Button
                    text={"Logar"}
                    click={(e) => handleLogin(e)}
                >
                </Button>
                <Link
                    href="/register"
                    className="text-center underline hover:text-indigo-800 cursor-pointer"
                >
                    Cadastrar-se
                </Link>
            </FormContainer>
        </ContainerBackgroud>
    )
}
export default Login