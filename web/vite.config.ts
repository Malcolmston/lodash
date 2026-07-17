import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { fileURLToPath } from 'node:url';

// The lodash repo is served as a GitHub *project* page at
// https://malcolmston.github.io/lodash/, so assets must be based under /lodash/.
export default defineConfig({
  base: '/lodash/',
  plugins: [react()],
  resolve: {
    alias: {
      // Import the vendored shared library from source.
      'go-ui': fileURLToPath(new URL('./vendor/go/ui/src/index.ts', import.meta.url)),
    },
  },
  build: { outDir: 'dist', emptyOutDir: true },
});
