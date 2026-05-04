<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER -->
      <div class="header">
        <h2>Sistem Absensi Digital Mahasiswa</h2>
      </div>

      <!-- LOGO -->
      <div class="logo-wrapper">
        <img :src="logo" class="logo" />
      </div>

      <!-- FORM -->
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
import logo from '../assets/logo.png'

const email = ref('')
const password = ref('')
const router = useRouter()
const route = useRoute()

// 🔥 AMBIL ROLE DARI HALAMAN ROLE
const role = route.query.role || ''

const goToRegister = () => {
  router.push({
    path: '/register',
    query: { role }
  })
}

const login = () => {
  // VALIDASI ROLE
  if (!role) {
    alert('Please select role first!')
    return
  }

  // VALIDASI INPUT
  if (!email.value || !password.value) {
    alert('Please fill email & password!')
    return
  }

  // VALIDASI EMAIL SESUAI ROLE
  if (role === 'lecturer' && !email.value.endsWith('@lecturer.university.ac.id')) {
    alert('Gunakan email dosen!')
    return
  }

  if (role === 'student' && !email.value.endsWith('@student.university.ac.id')) {
    alert('Gunakan email mahasiswa!')
    return
  }

  // 🔥 SIMPAN ROLE
  localStorage.setItem('role', role)

  // 🔥 REDIRECT SESUAI ROLE
  if (role === 'lecturer') {
    router.push('/beranda-dosen')
  } else {
    router.push('/beranda')
  }
}
</script>

<style scoped>
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
}

button {
  width: 100%;
  padding: 14px;
  background: #2f80ed;
  color: white;
  border: none;
  border-radius: 12px;
  font-weight: bold;
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