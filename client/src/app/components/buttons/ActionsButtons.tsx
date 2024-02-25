import React from "react"

interface IPropsActionsButtons {
  children: React.ReactNode
  text: string
  onClick?: () => void
}

export const ActionsButtons: React.FC<IPropsActionsButtons> = ({ children, text, onClick }) => {
  return (
    <button className="flex gap-1 items-center" onClick={onClick}>
      {children} {text}
    </button>
  )
}