import React from "react";

interface IPropsContainerInputs {
    children: React.ReactNode
}

export const ContainerInputs: React.FC<IPropsContainerInputs> = ({ children }) => {
    return (
        <div className="flex flex-col justify-between items-start">
            {children}
        </div>
    )
}