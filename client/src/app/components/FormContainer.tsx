import React from "react"

interface IPropsFormContainer {
    children: React.ReactNode
    title: string
}

export const FormContainer: React.FC<IPropsFormContainer> = ({ children, title }) => {
    return (
        <form className="flex flex-col bg-white px-6 py-4 my-2 rounded-2xl gap-6 text-sky-700 border-solid border-2 border-cyan-600 2xl:w-1/4 xl:w-1/4 md:w-72 sm:w-64">
            <h1 className="font-bold text-2xl flex justify-center">
                {title}
            </h1>
            {children}
        </form>
    )
}