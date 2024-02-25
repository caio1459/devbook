import Link from "next/link"
import React from "react"

interface IPropsMenuItems {
  href: string
  icon: React.JSX.Element
  text: string
}

export const MenuItems: React.FC<IPropsMenuItems> = ({ href, icon, text }) => {
  return (
    <Link href={href} className="flex gap-3 items-center">
      {icon}
      {text}
    </Link>
  )
}