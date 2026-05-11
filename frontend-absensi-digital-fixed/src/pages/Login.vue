<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER LOGIN -->
      <div class="header">
        <h2>Sistem Absensi Digital</h2>
        <p>Silakan masuk ke akun Anda</p>
      </div>

      <!-- LOGO APLIKASI -->
      <div class="logo-wrapper">
        <img :src="logo" class="logo" />
      </div>

      <!-- FORM INPUT & AKSI -->
      <div class="form">
        <div class="input-group">
          <input type="text" v-model="email" required />
          <label for="email">Email</label>
        </div>
        <div class="input-group">
          <input type="password" v-model="password" required />
          <label for="password">Password</label>
        </div>

        <button @click="login">LOGIN</button>

        <p class="signup">
          Don’t have an account?
          <span @click="goToRegister">Sign up</span>
        </p>

        <button style="width: 60%;" @click="goToRole">
          Back to Role Selection
        </button>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import axios from 'axios' 
import logo from '../assets/logo.png'

// --- INISIALISASI ROUTING & STATE ---
const email = ref('')
const password = ref('')
const router = useRouter()
const route = useRoute()

// Ambil role dari URL (misal: /login?role=student)
// Kalau kosong, kita set default ke 'student' biar formnya gak langsung error
const role = route.query.role || 'student' 

// --- LOGIKA NAVIGASI KE ROLE SELECTION ---
const goToRole = () => {
  router.push('/role')
}

// --- LOGIKA NAVIGASI KE REGISTER ---
const goToRegister = () => {
  router.push({
    path: '/register',
    query: { role }
  })
}

// --- PROSES AUTENTIKASI KE SERVER AZURE ---
const login = async () => {
  // 1. Validasi Input Lokal
  if (!email.value || !password.value) return alert('Email dan password wajib diisi!')

  // 2. Proteksi Domain Email (Sesuai kampus Telkom University)
  const domain = role === 'lecturer' ? '@telkomuniversity.ac.id' : '@student.telkomuniversity.ac.id'
  if (!email.value.endsWith(domain)) {
    alert(`Gunakan email resmi institusi untuk ${role}!`)
    return
  }

  try {
    // 3. Eksekusi Request dengan Fallback URL Azure (Anti-Gagal Netlify)
    const baseURL = import.meta.env.VITE_API_URL || 'https://sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net'
    
    // Pastikan nggak ada double slash pas URL digabung
    const cleanBaseURL = baseURL.replace(/\/$/, "")
    const endpoint = role === 'lecturer' ? '/login/dosen' : '/login/mahasiswa'
    
    const finalURL = `${cleanBaseURL}${endpoint}`

    // 4. Kirim Request ke Backend Golang
    const response = await axios.post(finalURL, {
      email: email.value,
      password: password.value
    }, {
      withCredentials: true // WAJIB supaya Azure bisa ngasih Cookie 'nim_user'
    })

    // 5. Manajemen Session Lokal
    localStorage.setItem('user_nama', response.data.nama) 
    localStorage.setItem('role', role)

    alert('Login Berhasil!')

    // 6. Routing Berdasarkan Role User
    router.push(role === 'lecturer' ? '/BerandaDosen' : '/Beranda')
    
  } catch (error) {
    // 7. Penanganan Error
    console.error('Login Error:', error)
    // Tampilkan pesan error spesifik dari Golang lu (misal: "Email atau Password salah")
    alert(error.response?.data?.error || 'Gagal terhubung ke server Azure!')
  }
}
</script>

<style scoped>
/* --- LAYOUT UTAMA --- */
.wrapper {
  background: #0f172a;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.phone {
  width: 390px;
  min-height: 780px;
  background: white;
  border-radius: 30px;
  overflow: hidden;
}

/* --- TAMPILAN VISUAL HEADER --- */
.header {
  background: #ff2d2d;
  color: white;
  text-align: center;
  padding: 40px 20px 90px;
  border-bottom-left-radius: 50% 25%;
  border-bottom-right-radius: 50% 25%;
}

.logo-wrapper {
  display: flex;
  justify-content: center;
  margin-top: -70px;
}

.logo {
  width: 95px;
}

/* --- ELEMEN FORM & INPUT --- */
.form {
  padding: 25px;
}

.input-group {
  position: relative;
  width: 320px;
  margin-bottom: 15px;
}
.input-group input {
  width: 100%;
  padding: 14px 10px;
  font-size: 14px;
  border: 1px solid black;
  background-color: white;
  border-radius: 12px;
  outline: none;
  margin-bottom: 0px;
  color:#0f172a;
}
.input-group label {
  position: absolute;
  top: 50%;
  left: 10px;
  transform: translateY(-50%);
  background: white;
  padding: 0 5px;
  color: #999;
  font-size: 12px;
  pointer-events: none;
  transition: all 0.3s ease-out;
}
.input-group input:focus + label,
.input-group input:valid+label {
  top: -1px;
  font-size: 10px;
  color: #2f80ed;
}

button {
  width: 100%;
  padding: 14px;
  background: #2f80ed;
  color: white;
  border: none;
  border-radius: 12px;
  font-weight: bold;
  cursor: pointer;
}

.signup {
  text-align: center;
  margin-top: 15px;
  margin-bottom: 15px;
  color: gray;
}

.signup span {
  color: #2f80ed;
  cursor: pointer;
  font-weight: bold;
}
</style>