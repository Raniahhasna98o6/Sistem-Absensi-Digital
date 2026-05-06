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
        <input type="text" placeholder="Email" v-model="email" />
        <input type="password" placeholder="Password" v-model="password" />

        <button @click="login">LOGIN</button>

        <p class="signup">
          Don’t have an account?
          <span @click="goToRegister">Sign up</span>
        </p>
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
const role = route.query.role || '' 

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
  if (!role) return alert('Pilih role terlebih dahulu!')
  if (!email.value || !password.value) return alert('Email dan password wajib diisi!')

  // 2. Proteksi Domain Email 
  const domain = role === 'lecturer' ? '@telkomuniversity.ac.id' : '@student.telkomuniversity.ac.id'
  if (!email.value.endsWith(domain)) {
    alert(`Gunakan email resmi institusi untuk ${role}!`)
    return
  }

  try {
    // 3. Eksekusi Request: Dinamis milih rute Golang lu
    const endpoint = role === 'lecturer' ? '/login/dosen' : '/login/mahasiswa'
    
    // PERBAIKAN KRUSIAL: Tambahkan withCredentials supaya Cookie dari Azure bisa masuk!
    const response = await axios.post(`${import.meta.env.VITE_API_URL}${endpoint}`, {
      email: email.value,
      password: password.value
    }, {
      withCredentials: true 
    })

    // 4. Manajemen Session
    localStorage.setItem('user_nama', response.data.nama) 
    localStorage.setItem('role', role)

    alert('Login Berhasil!')

    // 5. Routing Berdasarkan Role User
    router.push(role === 'lecturer' ? '/beranda-dosen' : '/beranda')
    
  } catch (error) {
    // 6. Penanganan Error Koneksi/Server
    console.error('Login Error:', error)
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

input {
  width: 100%;
  padding: 14px;
  margin-bottom: 15px;
  border-radius: 12px;
  border: 1px solid #ccc;
  background: #f5f5f5;
  color: black;
  box-sizing: border-box;
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
  color: gray;
}

.signup span {
  color: #2f80ed;
  cursor: pointer;
  font-weight: bold;
}
</style>