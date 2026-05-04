<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER PENDAFTARAN -->
      <div class="header">
        <h2>Sign Up</h2>
        <p>Sistem Absensi Digital</p>
      </div>

      <!-- LOGO APLIKASI -->
      <div class="logo-wrapper">
        <img :src="logo" class="logo" />
      </div>

      <!-- FORM PENDAFTARAN DINAMIS -->
      <div class="form">

        <!-- INPUT KHUSUS MAHASISWA -->
        <div v-if="role === 'student'">
          <input v-model="nama" placeholder="Nama Lengkap" />
          <input v-model="nim" placeholder="NIM" />
          <input v-model="prodi" placeholder="Program Studi" />
          <input v-model="fakultas" placeholder="Fakultas" />
          <input v-model="angkatan" placeholder="Tahun Angkatan" />
          <input v-model="email" placeholder="Email" />
          <input v-model="nohp" placeholder="No HP" />
        </div>

        <!-- INPUT KHUSUS DOSEN -->
        <div v-if="role === 'lecturer'">
          <input v-model="nama" placeholder="Nama Lengkap" />
          <input v-model="nidn" placeholder="NIDN" />
          <input v-model="email" placeholder="Email" />
          <input v-model="nohp" placeholder="No HP" />
          <input v-model="fakultas" placeholder="Fakultas" />
        </div>

        <input type="password" v-model="password" placeholder="Password" />

        <button @click="register">REGISTER</button>

        <p class="signup">
          Already have an account?
          <span @click="goToLogin">Login</span>
        </p>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import axios from 'axios' 
import logo from '../assets/logo.png'

// --- INISIALISASI ROUTING & ROLE ---
const router = useRouter()
const route = useRoute()
const role = computed(() => route.query.role || 'student')

// --- STATE FORM DATA ---
const nama = ref(''), email = ref(''), nohp = ref(''), password = ref('')
const nim = ref(''), prodi = ref(''), fakultas = ref(''), angkatan = ref(''), nidn = ref('')

// --- LOGIKA PENGIRIMAN DATA KE AZURE ---
const register = async () => {
  // Validasi input wajib
  if (!nama.value || !email.value || !password.value) {
    alert('Lengkapi data dulu!')
    return
  }

  // Susun payload berdasarkan role yang dipilih
  let payload = {
    nama: nama.value,
    email: email.value,
    nohp: nohp.value,
    password: password.value,
    role: role.value,
    fakultas: fakultas.value
  }

  if (role.value === 'student') {
    payload = { ...payload, nim: nim.value, prodi: prodi.value, angkatan: angkatan.value }
  } else {
    payload = { ...payload, nidn: nidn.value }
  }

  try {
    // Tembak endpoint register di backend Azure
    await axios.post(`${import.meta.env.VITE_API_URL}/register`, payload)
    
    alert('Registrasi Berhasil! Silakan Login.')
    router.push({ path: '/login', query: { role: role.value } })
  } catch (error) {
    console.error('Register Error:', error)
    alert(error.response?.data?.message || 'Gagal mendaftar ke server Azure.')
  }
}

// --- NAVIGASI KE HALAMAN LOGIN ---
const goToLogin = () => {
  router.push({ path: '/login', query: { role: role.value } })
}
</script>

<style scoped>
/* --- KONFIGURASI LAYOUT UTAMA --- */
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

/* --- TAMPILAN HEADER & LOGO --- */
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

/* --- STYLE FORM INPUT --- */
.form {
  padding: 25px;
}

input {
  width: 90%;
  padding: 14px;
  margin-bottom: 12px;
  border-radius: 12px;
  border: 1px solid #ccc;
  background: #f5f5f5;
  color: #000000;
}

button {
  width: 100%;
  padding: 14px;
  background: #2f80ed;
  color: white;
  border: none;
  border-radius: 12px;
  font-weight: bold;
  margin-top: 10px;
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
}
</style>