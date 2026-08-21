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
  resolve: {
    alias: {
      '$lib': path.resolve(__dirname, './src/lib'),
      '$wails': path.resolve(__dirname, './wailsjs'),
    }
  }
})
