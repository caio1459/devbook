import React from "react"

interface IPropsInput {
    idLabel: string
    type: 'email' | 'password' | 'text'
    textLabel: string
    placeholder?: string
    required?: boolean
    value: string
    onChange: (newValue: string) => void
}

export const Input: React.FC<IPropsInput> = ({ idLabel, type, placeholder, textLabel, required, value, onChange }) => {
    return (
        <>
            <label htmlFor={idLabel}>{textLabel}</label>
            <input
                className="border-sky-400 border-b w-full focus-visible:border-indigo-800 focus-visible:border-b focus-visible:outline-none"
                type={type}
                id={idLabel}
                placeholder={placeholder}
                required={required}
                value={value}
                onChange={e => onChange(e.target.value)}
            />
        </>
    )
}