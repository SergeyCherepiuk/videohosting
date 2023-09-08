/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    fontFamily: {
      "amazon-ember": ["Amazon Ember"],
    },
    colors: {
      "transparent": "transparent",
      "white": "#FFFFFF",
      "black": "#000000",
      "gray-30": "#555555",
      "gray-50": "#7B7B7B",
      "gray-60": "#AAAAAA",
      "gray-70": "#CCCCCC",
      "blue-50": "#1952E5",
      "blue-90": "#E2ECFB",
      "red-50": "#E12D2D",
      "red-90": "#FFE9E9",
    },
    extend: {
      container: {
        "reset-all": "initial",
      },
      flex: {
        "2": "2 2 0%",
        "3": "3 3 0%",
      },
    },
  },
  plugins: [],
}

