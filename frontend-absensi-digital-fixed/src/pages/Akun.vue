<template>
  <div class="wrapper">
    <div class="phone">

      <div class="header">
        <span class="back" @click="back">←</span>
        <h3>Akun Saya</h3>
      </div>

      <div class="content">

        <div class="profile">
          <img src="https://i.pravatar.cc/150?img=12" alt="profile" />
        </div>

        <div class="card nama-card">
          <h2>{{ user.nama }}</h2>
          <p>{{ user.nim }}</p>
        </div>

        <div class="card detail-card">

          <div class="row">
            <span>Program Studi</span>
            <span>{{ user.prodi }}</span>
          </div>

          <div class="divider"></div>

          <div class="row">
            <span>Fakultas</span>
            <span>{{ user.fakultas }}</span>
          </div>

          <div class="divider"></div>

          <div class="row">
            <span>Angkatan</span>
            <span>{{ user.angkatan }}</span>
          </div>

          <div class="divider"></div>

          <div class="row">
            <span>Email</span>
            <span style="text-align: right; max-width: 170px; display: block; overflow-wrap: break-word; word-break: break-word;">{{ user.email }}</span>
          </div>

          <div class="divider"></div>

          <div class="row">
            <span>No HP</span>
            <span>{{ user.nohp }}</span>
          </div>

        </div>

        <button class="logout" @click="logout">
          Logout
        </button>

      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()

const user = ref({
  nama: 'Memuat...',
  nim: '-', 
  prodi: '-',
  fakultas: '-',
  angkatan: '-',
  email: '-',
  nohp: '-'
})

const back = () => {
  router.push('/Beranda')
}

const logout = async () => {
  try {
    const baseURL = import.meta.env.VITE_API_URL || 'https://sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net'
    const cleanBaseURL = baseURL.replace(/\/$/, "")

    await axios.post(`${cleanBaseURL}/logout`, {}, {
      withCredentials: true 
    })
  } catch (error) {
    console.error('Server error saat logout:', error)
  } finally {
    localStorage.clear()
    router.push('/role')
  }
}

onMounted(async () => {
  const savedRole = localStorage.getItem('role')
  const savedUserNim = localStorage.getItem('user_nim')

  if (savedRole !== 'mahasiswa' && savedRole !== 'student') {
    router.push('/login?role=mahasiswa')
    return
  }

  if (savedUserNim) {
    try {
      const baseURL = import.meta.env.VITE_API_URL || 'https://sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net'
      const cleanBaseURL = baseURL.replace(/\/$/, "")

      // Vue nanya data lengkap ke Golang
      const response = await axios.get(`${cleanBaseURL}/api/mahasiswa/profil?nim=${savedUserNim}`, {
        withCredentials: true
      })

      // Timpa state user dengan data dari database!
      if (response.data) {
        user.value = {
          nama: response.data.nama,
          nim: response.data.nim,
          prodi: response.data.prodi,
          fakultas: response.data.fakultas,
          angkatan: response.data.angkatan,
          email: response.data.email,
          nohp: response.data.nohp
        }
      }
    } catch (error) {
      console.error('Gagal mengambil profil:', error)
      user.value.nama = localStorage.getItem('user_nama') || 'Error memuat'
    }
  } else {
    router.push('/login?role=mahasiswa')
  }
})
</script>

<style scoped>
/* BACKGROUND */
.wrapper { min-height: 100vh; background: #0f1c2e; display: flex; justify-content: center; align-items: center; }
/* HEADER */
.header { 
  padding: 15px; 
  background: #f3f3f3; 
  display: flex; 
  align-items:end; 
  gap: 10px;
  border-bottom: 1px solid lightgray;
}
.header h3 { font-weight: bold; font-size: 18px; color: black; margin: 25px 0 0 0; }
.back { font-size: 20px; cursor: pointer; color: #000; }
/* CONTENT */
.content { flex: 1; padding: 20px; display: flex; flex-direction: column; gap: 16px; align-items: center; }
/* FOTO */
.profile { margin-top: 10px; }
.profile img { width: 110px; height: 110px; border-radius: 50%; object-fit: cover; border: 4px solid white; box-shadow: 0 6px 12px rgba(0,0,0,0.2); }
/* CARD FIX UTAMA */
.card { width: 100%; max-width: 330px; background: white; border-radius: 18px; padding: 20px; box-shadow: 0 10px 18px rgba(0,0,0,0.12); border: 1px solid #ea3236; }
/* NAMA CARD */
.nama-card { text-align: center; }
.nama-card h2 { font-size: 20px; font-weight: 700; color: black; }
.nama-card p { font-size: 16px; color: black; }
/* DETAIL */
.detail-card { font-size: 15px; }
.row { display: flex; justify-content: space-between; padding: 10px 6px; color: black; }
.divider { height: 1px; background: #ddd; }
/* LOGOUT */
.logout { width: 100%; max-width: 330px; padding: 16px; background: #d32f2f; border-radius: 16px; border: 1px solid pink; color: white; font-weight: 700; font-size: 20px; margin-top: 20px; box-shadow: 0 6px 12px rgba(255,0,0,0.3); }
</style>