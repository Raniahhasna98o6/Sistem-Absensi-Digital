<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER TOPBAR -->
      <div class="topbar">
        <div class="top-left">
          <svg class="icon-user" viewBox="0 0 24 24">
            <path d="M12 12c2.7 0 5-2.3 5-5s-2.3-5-5-5-5 2.3-5 5 2.3 5 5 5z"/>
            <path d="M2 22c0-4 4-7 10-7s10 3 10 7"/>
          </svg>
          <h3>Beranda</h3>
        </div>
      </div>

      <!-- AREA KONTEN (TEMA MERAH) -->
      <div class="red">
        <!-- Nama user dinamis dari database Azure -->
        <h2 class="title">Halo, {{ user.nama || 'Mahasiswa' }}</h2>

        <!-- KARTU ABSENSI HARI INI -->
        <div class="card">
          <h3 class="card-title">Absensi Hari Ini</h3>
          <div class="card-inner">
            <div class="row">
              <svg class="icon-calendar" viewBox="0 0 24 24">
                <rect x="3" y="5" width="18" height="16" rx="2"/>
                <line x1="3" y1="10" x2="21" y2="10"/>
              </svg>
              <div class="text-group">
                <p class="matkul">Jaringan Komputer</p>
                <p class="jam">08.00 - 10.00</p>
              </div>
            </div>
            <button @click="absen">ABSEN SEKARANG</button>
          </div>
        </div>

        <!-- MENU FITUR UTAMA -->
        <h3 class="fitur-title">Fitur Website</h3>
        <div class="menu">
          <div class="menu-card" @click="keAkun">
            <svg class="menu-icon" viewBox="0 0 24 24">
              <path d="M12 12c2.7 0 5-2.3 5-5s-2.3-5-5-5-5 2.3-5 5 2.3 5 5 5z"/>
              <path d="M2 22c0-4 4-7 10-7s10 3 10 7"/>
            </svg>
            <p>Akun</p>
          </div>

          <div class="menu-card" @click="keRiwayat">
            <svg class="menu-icon" viewBox="0 0 24 24">
              <circle cx="12" cy="12" r="9"/>
              <path d="M12 7v5l3 2"/>
            </svg>
            <p>Riwayat Absensi</p>
          </div>
        </div>
      </div>

      <!-- FOOTER AREA -->
      <div class="bottom"></div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

// --- INISIALISASI STATE & ROUTER ---
const router = useRouter()
const user = ref({ nama: 'Mahasiswa' })

// --- LOGIKA NAVIGASI ---
const absen = () => router.push('/absensi')
const keAkun = () => router.push('/akun')
const keRiwayat = () => router.push('/riwayat')

// --- AMBIL DATA USER DARI STORAGE SAAT HALAMAN DIMUAT ---
onMounted(() => {
  const savedUser = localStorage.getItem('user')
  const savedRole = localStorage.getItem('role')

  // Proteksi: Pastikan hanya mahasiswa yang bisa akses
  if (savedRole !== 'student') {
    router.push('/login')
    return
  }

  // Load data hasil sukses login dari database Azure
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
  background: #0f172a;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.phone {
  width: 390px;
  height: 780px;
  background: white;
  border-radius: 30px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

/* --- STYLE HEADER & TEMA --- */
.topbar {
  padding: 15px;
  background: #f3f3f3;
}

.top-left {
  display: flex;
  align-items: center;
  gap: 10px;
}

.topbar h3 {
  font-weight: bold;
  color: black;
}

.icon-user {
  width: 30px;
  height: 30px;
  fill: #6b46c1;
}

.red {
  background: #ff2d2d;
  padding: 20px;
  color: white;
  flex: 1;
}

.title {
  text-align: left;
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 15px;
}

/* --- KARTU ABSENSI & TOMBOL --- */
.card {
  background: #f1f1f1;
  border-radius: 20px;
  padding: 15px;
  color: black;
}

.card-title {
  text-align: left;
  font-size: 18px;
  margin-bottom: 10px;
}

.card-inner {
  background: white;
  border-radius: 15px;
  padding: 15px;
  box-shadow: 0 10px 20px rgba(0,0,0,0.25);
}

.row {
  display: flex;
  align-items: center;
  gap: 14px;
}

.text-group {
  display: flex;
  flex-direction: column;
}

.icon-calendar {
  width: 32px;
  height: 32px;
  stroke: black;
  fill: none;
  stroke-width: 2.5;
}

button {
  width: 100%;
  margin-top: 14px;
  padding: 14px;
  background: #2f80ed;
  color: white;
  border: none;
  border-radius: 12px;
  font-weight: bold;
  cursor: pointer;
}

/* --- MENU NAVIGATION CARDS --- */
.fitur-title {
  text-align: left;
  margin-top: 25px;
  font-size: 18px;
  font-weight: bold;
}

.menu {
  display: flex;
  gap: 12px;
  margin-top: 10px;
}

.menu-card {
  flex: 1;
  background: white;
  color: black;
  padding: 20px;
  border-radius: 15px;
  text-align: center;
  box-shadow: 0 8px 16px rgba(0,0,0,0.25);
  cursor: pointer;
}

.menu-icon {
  width: 40px;
  height: 40px;
  stroke: black;
  fill: none;
  stroke-width: 3;
  margin-bottom: 10px;
}

.bottom {
  height: 35px;
  background: #f3f3f3;
}
</style>