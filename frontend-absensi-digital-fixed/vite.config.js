import { defineConfig } from 'vite'
import vue from '@vue/dist/vue.esm-bundler.js' // atau import vue biasa tergantung setup lu

export default defineConfig({
  plugins: [vue()],
  server: {
    // Tambahin ini biar Ngrok diizinin masuk
    allowedHosts: ['brisket-simply-chafe.ngrok-free.dev'] 
    // Atau kalau mau bebas (buat testing):
    // allowedHosts: 'all'
  }
})