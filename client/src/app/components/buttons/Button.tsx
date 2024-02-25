import React from "react"

interface IPropsButton {
    text: string
    click: (e:React.MouseEvent<HTMLButtonElement, MouseEvent>) => void
}

export const Button: React.FC<IPropsButton> = ({ text, click }) => {
    return (
        <button
            className="bg-gradient-to-r from-sky-500 to-indigo-500 py-2 font-bold text-white rounded-full cursor-pointer"
            onClick={(e) => click(e)}
        >
            {text}
        </button>
    )
}