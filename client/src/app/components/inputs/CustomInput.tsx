import React from "react"

interface IPropsCustomInput {
  children: React.ReactNode
  placeholder: string
  value?: string
  onChange: (newValue: string) => void | undefined
}

export const CustomInput: React.FC<IPropsCustomInput> = ({ children, placeholder, onChange, value }) => {
  return (
    <div className="flex bg-zinc-200 text-gray-600 px-3 py-1 rounded-full items-center w-full">
      <input
        type="text"
        className="bg-zinc-200 outline-none w-full"
        placeholder={placeholder}
        onChange={(e) => onChange(e.target.value)}
        required
        value={value}
      />
      {children}
    </div>
  )
}