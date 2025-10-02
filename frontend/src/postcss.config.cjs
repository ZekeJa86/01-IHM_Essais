// La nouvelle configuration correcte
import tailwindcss from '@tailwindcss/postcss'; // <-- La seule ligne qui change
import autoprefixer from 'autoprefixer';

export default {
  plugins: [
    tailwindcss,
    autoprefixer,
  ],
};