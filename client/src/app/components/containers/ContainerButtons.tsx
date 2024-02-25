import React from "react"

interface IPropsContainerButtons {
  children: React.ReactNode
}

export const ContainerButtons: React.FC<IPropsContainerButtons> = ({ children }) => {
  return (
    <div className="flex justify-around py-4 text-gray-600 b-4">
      {children}
    </div>
  )
}