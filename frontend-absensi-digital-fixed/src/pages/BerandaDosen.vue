<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER NAVIGATION -->
      <div class="header">
        <div class="user-icon">👤</div>
        <h3 class="beranda">Beranda Dosen</h3>
      </div>

      <!-- AREA KONTEN UTAMA (RED THEME) -->
      <div class="red">
        <h2 class="title">Halo, {{ user.nama || 'Dosen' }}</h2>

        <!-- KARTU PROFIL DINAMIS DARI DATABASE -->
        <div class="card">
          <div class="profile">
            <img src="https://i.pravatar.cc/100" class="avatar" />
            <div>
              <p class="nama">{{ user.nama }}</p>
              <p class="nim">{{ user.nidn }}</p>
            </div>
          </div>

          <div class="info">
            <div class="row">
              <span>Program Studi</span>
              <span>{{ user.prodi }}</span>
            </div>
            <div class="row">
              <span>Fakultas</span>
              <span>{{ user.fakultas }}</span>
            </div>
            <div class="row">
              <span>Email</span>
              <span>{{ user.email }}</span>
            </div>
            <div class="row">
              <span>No HP</span>
              <span>{{ user.nohp }}</span>
            </div>
          </div>
        </div>

        <!-- MENU FITUR DOSEN -->
        <h3 class="fitur">Fitur Website</h3>
        <div class="laporan" @click="generate">
          ⏱️ Generate Laporan Absensi
        </div>

        <!-- LOGOUT & PEMBERSIHAN SESI -->
        <button class="logout" @click="logout">Logout</button>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

// --- INISIALISASI STATE & ROUTER ---
const router = useRouter()
const user = ref({
  nama: 'Memuat...',
  nidn: '...',
  prodi: '...',
  fakultas: '...',
  email: '...',
  nohp: '...'
})

// --- LOGIKA PEMBERSIHAN SESI SAAT KELUAR ---
const logout = () => {
  localStorage.clear() // Menghapus token dan data user dari browser
  router.push('/login')
}

const generate = () => {
  router.push('/laporan')
}

// --- AMBIL DATA PROFIL SAAT HALAMAN DIMUAT ---
onMounted(() => {
  const savedUser = localStorage.getItem('user')
  const savedRole = localStorage.getItem('role')

  // Validasi: Pastikan hanya dosen yang bisa akses halaman ini
  if (savedRole !== 'lecturer') {
    router.push('/login')
    return
  }

  // Load data dari storage hasil login sukses di Azure
  if (savedUser) {
    user.value = JSON.parse(savedUser)
  } else {
    router.push('/login')
  }
})
</script>

<style scoped>
/* --- KONFIGURASI LAYOUT TELEPON --- */
.wrapper {
  min-height: 100vh;
  background: #0f1c2e;
  display: flex;
  justify-content: center;
  align-items: center;
}

.phone {
  width: 390px;
  height: 800px;
  background: white;
  border-radius: 30px;
  overflow: hidden;
}

/* --- STYLE HEADER & TEMA MERAH --- */
.header {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 15px;
  background: #f3f3f3;
}

.beranda {
  color: black;
  font-weight: 700;
}

.red {
  background: #ff2d2d;
  padding: 20px;
  height: 100%;
  color: white;
}

.title {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 15px;
  text-align: left;
}

/* --- STYLE KARTU PROFIL & DATA --- */
.card {
  background: #f3f3f3;
  color: black;
  border-radius: 20px;
  padding: 16px;
  margin-bottom: 15px;
}

.profile {
  display: flex;
  align-items: center;
  gap: 12px;
}

.avatar {
  width: 55px;
  height: 55px;
  border-radius: 50%;
}

.nama {
  font-weight: bold;
}

.nim {
  font-size: 13px;
}

.row {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid #ddd;
  padding: 6px 0;
}

/* --- STYLE MENU & TOMBOL KELUAR --- */
.fitur {
  margin-top: 10px;
  text-align: left;
  font-weight: 600;
}

.laporan {
  background: white;
  color: black;
  padding: 14px;
  border-radius: 14px;
  margin-top: 10px;
  font-weight: 600;
  cursor: pointer;
}

.logout {
  margin-top: 15px;
  width: 100%;
  padding: 14px;
  border-radius: 14px;
  border: none;
  background: white;
  color: red;
  font-weight: bold;
  cursor: pointer;
}
</style>