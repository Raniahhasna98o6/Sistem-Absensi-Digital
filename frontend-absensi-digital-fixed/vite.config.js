import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue' // <--- Ganti jadi ini!

export default defineConfig({
  plugins: [vue()],
  server: {
    // Biar Ngrok diizinin masuk tanpa protes "Blocked Request"
    allowedHosts: ['brisket-simply-chafe.ngrok-free.dev']
  }
})