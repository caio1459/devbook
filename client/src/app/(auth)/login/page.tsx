"use client";
import Link from "next/link"
import { Input } from "../../components/inputs/Input"
import { ContainerInputs } from "../../components/containers/ContainerInputs"
import { Button } from "../../components/buttons/Button"
import { useCallback, useState } from "react"
import { ContainerBackgroud } from "../../components/containers/ContainerBackgroud";
import { FormContainer } from "../../components/containers/FormContainer";
import AxiosService from "../../../services/AxiosService";
import { useRouter } from "next/navigation";

const Login = () => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const router = useRouter()

  const clearInputs = () => {
    setEmail('')
    setPassword('')
  }

  const handleLogin = useCallback((e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault()
    AxiosService.makeRequest().post("/login", { email, password }).then(res => {
      localStorage.setItem("devbook:token", JSON.stringify(res.data.token))
      localStorage.setItem("devbook:user", JSON.stringify(res.data.user))
      setError("")
      router.push("/")
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