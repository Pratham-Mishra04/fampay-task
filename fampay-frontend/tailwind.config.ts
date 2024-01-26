import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        primary: ['var(--inter-font)'],
      },
      colors: {
        primary_black: '#2e2c2c',
      },
      animation: {
        fade_third: 'fade 0.3s ease-in-out',
        fade_half: 'fade 0.5s ease-in-out',
        fade_1: 'fade 1s ease-in-out',
        fade_2: 'fade 2s ease-in-out',
      },
      keyframes: {
        fade: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' },
        },
      },
    },
  },
  plugins: [],
};
export default config;
