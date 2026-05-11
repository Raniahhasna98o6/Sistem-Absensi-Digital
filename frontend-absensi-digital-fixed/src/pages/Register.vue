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
          <div class="input-group">
            <input type="text" v-model="nama" required />
            <label for="nama">Nama Lengkap</label>
          </div>
          <div class="input-group">
            <input type="text" v-model="nim" required />
            <label for="nim">NIM</label>
          </div>
          <div class="input-group">
            <input type="text" v-model="prodi" required />
            <label for="prodi">Program Studi</label>
          </div>
          <div class="input-group">
            <input type="text" v-model="fakultas" required />
            <label for="fakultas">Fakultas</label>
          </div>
          <div class="input-group">
            <input type="number" v-model="angkatan" required />
            <label for="angkatan">Tahun Angkatan</label>
          </div>
          <div class="input-group">
            <input type="email" v-model="email" required />
            <label for="email">Email</label>
          </div>
          <div class="input-group">
            <input type="number" v-model="nohp" required />
            <label for="nohp">No HP</label>
          </div>
        </div>

        <!-- INPUT KHUSUS DOSEN -->
        <div v-if="role === 'lecturer'">
          <div class="input-group">
             <input type="text" v-model="nama" required />
             <label for="nama">Nama Lengkap</label>
          </div>
          <div class="input-group">
            <input type="text" v-model="nidn" required />
            <label for="nidn">NIDN</label>
          </div>
          <div class="input-group">
            <input type="email" v-model="email" required />
            <label for="email">Email</label>
          </div>
          <div class="input-group">
            <input type="number" v-model="nohp" required />
            <label for="nohp">No HP</label>
          </div>
          <div class="input-group">
            <input type="text" v-model="fakultas" required />
            <label for="fakultas">Fakultas</label>
          </div>
        </div>
        <div class="input-group">
          <input type="password" v-model="password" required />
          <label for="password">Password</label>
        </div>

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
  color:#1e293b;
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
.input-group input:valid + label {
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