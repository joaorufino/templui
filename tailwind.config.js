
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "cmd/tailwindconfig/**/*.{templ,go}",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('flowbite/plugin'),
  ],
}
