/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        // Palette officielle NOORCHAIN 1.1
        primary: "#1A6AFF",      // Bleu signature
        light: "#00D1B2",        // Turquoise NOOR
        navy: "#0C2D48",         // Texte fort / institutionnel
        "gray-soft": "#E5E7EB",  // Backgrounds légers
        paper: "#F9FAFB",        // Blanc cassé
        "indigo-deep": "#3B3F75", // Accent premium
        "mint-light": "#A0F1DD", // Accent doux

        // Sémantique
        success: "#22C55E",
        warning: "#FACC15",
        danger: "#EF4444",
      },
    },
  },
  plugins: [],
};
