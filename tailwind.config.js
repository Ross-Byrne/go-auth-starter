/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.templ"],
  theme: {
    extend: {
      lineHeight: {
        11: "2.75rem",
      },
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      "light",
      // "dark", // Remove for now
    ],
  },
  // darkMode: ["selector", '[data-theme="dark"]'],
};
