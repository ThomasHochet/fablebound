import {defineConfig} from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import tailwindcss from '@tailwindcss/vite'
import Icons from 'unplugin-icons/vite'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    tailwindcss(),
    svelte(),
    Icons({
      compiler: 'svelte',
    })
  ],
  optimizeDeps: {
    exclude: ['@tiptap/pm']
  },
  server: {
    host: '127.0.0.1',
    port: 9245,
    strictPort: true
  },
  resolve: {
    alias: {
      '$lib': path.resolve(__dirname, './src/lib'),
      '$wails': path.resolve(__dirname, './bindings'),
    }
  }
})
