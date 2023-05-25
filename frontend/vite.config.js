import { defineConfig } from 'vite';
import { createHtmlPlugin } from 'vite-plugin-html';
import react from '@vitejs/plugin-react';

export default defineConfig({
  root: 'src',
  build: {
    outDir: '../dist',
  },
  plugins: [
    createHtmlPlugin({
      inject: {
        data: {
          title: 'Shotdawn',
        },
      },
    }),
    react({
      include: '**/*.{jsx,tsx}',
    }),
  ],
  server: {
    host: '0.0.0.0',
    port: 8080,
  },
});
