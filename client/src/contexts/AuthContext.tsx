import { AxiosRequestConfig } from "axios"
import React, { createContext } from "react"

interface IPropsAuthContext {
  headers: AxiosRequestConfig<any>
}

interface IPropsAuthContextProvider {
  children: React.ReactNode
}

export const AuthContext = createContext<IPropsAuthContext>({} as IPropsAuthContext)

export const AuthContextProvider: React.FC<IPropsAuthContextProvider> = ({ children }) => {
  let storedToken = "";
  let token = "";

  if (typeof window !== 'undefined') {
    // Verificar se o localStorage está disponível no ambiente do navegador
    storedToken = localStorage.getItem("devbook:token") || "";
    token = storedToken.replace(/^"(.*)"$/, '$1');
  }

  const headers: AxiosRequestConfig = {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };

  return (
    <AuthContext.Provider value={{
      headers
    }}>
      {children}
    </AuthContext.Provider>
  )
}