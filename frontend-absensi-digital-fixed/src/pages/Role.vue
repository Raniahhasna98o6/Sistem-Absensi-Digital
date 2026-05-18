<template>
  <div class="wrapper">
    <div class="phone">
      <div class="header">
        <h1>Sistem Absensi Digital</h1>
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
          <div class="icon-wrapper">
            <img :src="lecturer" class="card-icon" width="60" height="60"/>
          </div>
          <div class="text-group">
            <h2 style="color: cadetblue;">Lecturer</h2>
            <p>@telkomuniversity.ac.id</p>
          </div>
        </div>

        <!-- SELEKSI MAHASISWA -->
        <div class="card" :class="{ 'selected': role === 'student' }" @click="selectRole('student')">
          <div class="icon-wrapper">
            <img :src="student" class="card-icon" width="60" height="60 "/>
          </div>
          <div class="text-group">
            <h2 style="color: cadetblue;">Student</h2>
            <p>@student.telkomuniversity.ac.id</p>
          </div>
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
import lecturer from '../assets/ph--chalkboard-teacher-duotone.svg'
import student from '../assets/ph--student-duotone.svg'

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
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap');
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
  background-color: #ea3236;
  color: white;
  text-align: center;
  font-family: 'Poppins', sans-serif;
  padding: 45px 20px 40px;
  border-bottom-left-radius: 50% 25%;
  border-bottom-right-radius: 50% 25%;
}
.header h1 {
  margin: 0;
  font-size: 30px;
  font-weight: 700;
  letter-spacing: 1px;
  line-height: 1.2;
  color: rgba(255, 255, 255, 0.95);
  text-shadow: 0 3px 15px rgba(255, 255, 255, 0.2);
  font-family: 'Poppins', sans-serif;
}
.header p {
  margin-top: 6px;
  font-size: 14px;
  line-height: 15px;
}

.logo-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 15px;
  margin-bottom: -25px;
}

.logo {
  width: 75px;
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
  display: flex;
  padding: 18px;
  gap: 18px;
  align-items: center;
  border-radius: 18px;
  border: 1px solid #ddd;
  background: #f6f6f6;
  text-align: center;
  cursor: pointer;
  transition: 0.2s;
  box-shadow: 0 2px 8px rgba(0,0,0,0.4);
}
.card.selected {
  background: lightcyan;
  border: 1px solid #ea3236;
}
.card-icon {
  width: 50px;
  height: 50px;
  object-fit: contain;
}

.icon-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 45px;
}

.text-group {
  flex: 1;
  text-align: left;
}
.text-group h2 {
  margin: 0;
}

button {
  margin-top: 15px;
  padding: 14px;
  border: none;
  border-radius: 12px;
  background: #ea3236;
  color: white;
  font-weight: bold;
  cursor: pointer;
}

button:disabled {
  background: gray;
  cursor: not-allowed;
}
</style>