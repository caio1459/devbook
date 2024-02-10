import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
        fundo:
          "url('https://th.bing.com/th/id/OIG3._61uKV7T7_Kdv_BDbKGP?w=1024&h=1024&rs=1&pid=ImgDetMain')",
      },
    },
  },
  plugins: [],
};
export default config;
