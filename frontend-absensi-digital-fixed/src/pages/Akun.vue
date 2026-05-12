<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER -->
      <div class="header">
        <span class="back" @click="back">←</span>
        <h3>Akun Saya</h3>
      </div>

      <!-- CONTENT -->
      <div class="content">

        <!-- FOTO -->
        <div class="profile">
          <img src="https://i.pravatar.cc/150?img=12" alt="profile" />
        </div>

        <!-- CARD NAMA -->
        <div class="card nama-card">
          <h2>Sean Benedict</h2>
          <p>220123455</p>
        </div>

        <!-- CARD DETAIL -->
        <div class="card detail-card">

          <div class="row">
            <span>Program Studi</span>
            <span>S1 Informatika</span>
          </div>

          <div class="divider"></div>

          <div class="row">
            <span>Fakultas</span>
            <span>Informatika</span>
          </div>

          <div class="divider"></div>

          <div class="row">
            <span>Angkatan</span>
            <span>2024</span>
          </div>

          <div class="divider"></div>

          <div class="row">
            <span>Email</span>
            <span>sean@student.com</span>
          </div>

          <div class="divider"></div>

          <div class="row">
            <span>No HP</span>
            <span>0812345678910</span>
          </div>

        </div>

        <!-- LOGOUT -->
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

// 1. State awal, disesuaikan karena backend lu belum ngirim data lengkap
const user = ref({
  nama: 'Memuat...',
  nim: 'Menunggu API...', // Harus bikin API di Golang buat ini
  prodi: '-',
  fakultas: '-',
  angkatan: '-',
  email: '-',
  nohp: '-'
})

const back = () => {
  router.push('/Beranda')
}

// 2. Logika Logout Full-Stack (Lokal + Server Azure)
const logout = async () => {
  try {
    const baseURL = import.meta.env.VITE_API_URL || 'https://sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net'
    const cleanBaseURL = baseURL.replace(/\/$/, "")

    // Tembak API Logout di Golang buat ngehancurin Cookie!
    await axios.post(`${cleanBaseURL}/logout`, {}, {
      withCredentials: true // Wajib supaya Azure ngebaca Cookie mana yang mau dihancurin
    })
  } catch (error) {
    console.error('Server error saat logout:', error)
  } finally {
    // 3. Bersihkan sisa data di browser & tendang ke Login
    localStorage.clear()
    router.push('/login?role=mahasiswa')
  }
}

// 4. Ambil data profil dari sistem yang udah kita bangun
onMounted(() => {
  const savedUserNama = localStorage.getItem('user_nama')
  const savedRole = localStorage.getItem('role')

  // Proteksi ganda (terima 'mahasiswa' atau 'student')
  if (savedRole !== 'mahasiswa' && savedRole !== 'student') {
    router.push('/login?role=mahasiswa')
    return
  }

  // Load nama dari storage hasil login tadi
  if (savedUserNama) {
    user.value.nama = savedUserNama
    
    /* 
      NOTE PENTING: 
      Saat ini Golang lu di main.go BELUM PUNYA rute buat ngambil profil lengkap (NIM, Prodi, dll).
      Kalau lu mau data profilnya lengkap, lu harus bikin rute GET "/api/profil" di Golang, 
      terus tembak pakai Axios di sini pakai { withCredentials: true }.
      Untuk sementara, kita tampilin namanya aja dulu biar nggak error!
    */
    
  } else {
    router.push('/login?role=mahasiswa')
  }
})
</script>

<style scoped>

/* BACKGROUND */
.wrapper {
  min-height: 100vh;
  background: #0f1c2e;
  display: flex;
  justify-content: center;
  align-items: center;
}

/* PHONE */
.phone {
  width: 390px;
  height: 800px;
  background: #f6f6f6;
  border-radius: 30px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}
.phone:gradient-style {
  background: linear-gradient(135deg, #0f1c2e, #ffffff);
}

/* HEADER */
.header {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 16px;
  background: #ffffff;
}

.header h3 {
  font-weight: 700;
  font-size: 18px;
  color: black;
}

.back {
  font-size: 20px;
  cursor: pointer;
}

/* CONTENT */
.content {
  flex: 1;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  align-items: center;
}

/* FOTO */
.profile {
  margin-top: 10px;
}

.profile img {
  width: 110px;
  height: 110px;
  border-radius: 50%;
  object-fit: cover;
  border: 4px solid white;
  box-shadow: 0 6px 12px rgba(0,0,0,0.2);
}

/* CARD FIX UTAMA */
.card {
  width: 100%;
  max-width: 330px;
  background: white;
  border-radius: 18px;
  padding: 20px;
  box-shadow: 0 10px 18px rgba(0,0,0,0.12);
  border: 1px solid #ff4d4f;
}

/* NAMA CARD */
.nama-card {
  text-align: center;
}

.nama-card h2 {
  font-size: 20px;
  font-weight: 700;
  color: black;
}

.nama-card p {
  font-size: 16px;
  color: black;
}

/* DETAIL */
.detail-card {
  font-size: 15px;
}

/* ROW */
.row {
  display: flex;
  justify-content: space-between;
  padding: 10px 6px;
  color: black;
}

/* DIVIDER */
.divider {
  height: 1px;
  background: #ddd;
}

/* LOGOUT */
.logout {
  width: 100%;
  max-width: 330px;
  padding: 16px;
  background: red;
  border-radius: 16px;
  border: 1px solid pink;
  color: white;
  font-weight: 700;
  font-size: 20px;
  margin-top: 20px;
  box-shadow: 0 6px 12px rgba(255,0,0,0.3);
}
</style>