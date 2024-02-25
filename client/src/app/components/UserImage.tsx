import React from "react";

interface IPropsUserImage {
  image_url: string | undefined
}

export const UserImage: React.FC<IPropsUserImage> = ({ image_url }) => {
  return (
    <>
      <img
        className="w-10 h-10 rounded-full"
        src={image_url ? image_url : "https://cdn-icons-png.flaticon.com/512/6558/6558722.png"}
      />
    </>
  )
}