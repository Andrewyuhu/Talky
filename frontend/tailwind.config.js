/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,vue,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        "primary-green": "var( --color-primary-green)",
        "primary-background": "var( --color-background)",
        "primary-text": "var(--color-text)",
        "secondary-green": "var(--color-light-green)",
      },
    },
  },
  plugins: [],
  safelist: [
    "bg-red-200",
    "bg-blue-200",
    "bg-green-200",
    "bg-yellow-200",
    "bg-purple-200",
  ],
};
