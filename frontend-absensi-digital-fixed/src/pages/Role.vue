<template>
  <div class="wrapper">
    <div class="phone">
      <div class="header">
        <h2>Choose Role</h2>
        <p>Please select your role to continue</p>
        <p style="font-size: smaller;">Use the format below on your email</p>
      </div>

      <!-- LOGO SEKSI -->
      <div class="logo-wrapper">
        <img :src="logo" class="logo" />
      </div>
      
      <div class="content">
        <!-- SELEKSI DOSEN -->
        <div class="card" :class="{ 'selected': role === 'lecturer' }" @click="selectRole('lecturer')">
          <h2 style="color: cadetblue;">Lecturer</h2>
          <p>@telkomuniversity.ac.id</p>
        </div>

        <!-- SELEKSI MAHASISWA -->
        <div class="card" :class="{ 'selected': role === 'student' }" @click="selectRole('student')">
          <h2 style="color: cadetblue;">Student</h2>
          <p>@student.telkomuniversity.ac.id</p>
        </div>

        <button :disabled="!role" @click="goToLogin">
          Continue
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import logo from '../assets/logo.png'

const router = useRouter()
const role = ref(null)

// --- MEMBERSIHKAN DATA LOGIN LAMA ---
onMounted(() => {
  localStorage.clear() // Memastikan sesi benar-benar bersih saat pilih role baru[cite: 1]
})

// --- LOGIKA PEMILIHAN ROLE ---
const selectRole = (selectedRole) => {
  role.value = selectedRole
}

// --- NAVIGASI KE HALAMAN LOGIN ---
const goToLogin = () => {
  if (!role.value) {
    alert('Please select a role first!')
    return
  }

  // Mengirim info role via query parameter ke Login_2.vue[cite: 1, 5]
  router.push({ 
    path: '/login',
    query: { role: role.value }
  })
}
</script>

<style scoped>
/* --- KONFIGURASI TAMPILAN UTAMA --- */
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
  background-color: #ff2d2d;
  color: white;
  text-align: center;
  padding: 50px 20px 100px;
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

/* --- KARTU ROLE & TOMBOL --- */
.content {
  padding: 25px;
  margin-top: 10px;
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.card {
  padding: 18px;
  border-radius: 15px;
  border: 1px solid #ddd;
  background: #f6f6f6;
  text-align: center;
  cursor: pointer;
  transition: 0.2s;
  box-shadow: 0 2px 8px rgba(0,0,0,0.4);
}

.card.selected {
  background: lightcyan;
  border: 1px solid #ff2d2d;
}

button {
  margin-top: 15px;
  padding: 14px;
  border: none;
  border-radius: 12px;
  background: #ff2d2d;
  color: white;
  font-weight: bold;
  cursor: pointer;
}

button:disabled {
  background: gray;
  cursor: not-allowed;
}
</style>