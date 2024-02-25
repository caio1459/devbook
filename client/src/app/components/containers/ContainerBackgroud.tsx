import React from "react"

interface IPropsContainerBackgroud {
  children: React.ReactNode
}
export const ContainerBackgroud: React.FC<IPropsContainerBackgroud> = ({ children }) => {
  return (
    <main className="bg-fundo bg-no-repeat bg-cover flex min-h-screen flex-col items-center justify-center">
      {children}
    </main>
  )
}