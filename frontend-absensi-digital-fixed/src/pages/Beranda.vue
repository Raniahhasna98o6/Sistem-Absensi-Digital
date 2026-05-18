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
              <img :src="kalender" class="logo" width="60"/>
              <div class="text-group">
                <p>Jaringan Komputer</p>
                <p>08.00 - 10.00</p>
              </div>
            </div>
            <button @click="absen">ABSEN SEKARANG</button>
          </div>
        </div>

        <!-- MENU FITUR UTAMA -->
        <h3 class="fitur-title">Fitur Website</h3>
        <div class="menu">
          <div class="menu-card" @click="keAkun">
            <img :src="akun" class="menu-icon"/>
            <p><b>Akun</b></p>
          </div>

          <div class="menu-card" @click="keRiwayat">
            <img :src="riwayat" class="menu-icon"/>
            <p><b>Riwayat Absensi</b></p>
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
import kalender from '../assets/marketeq_date.svg'
import akun from '../assets/streamline-plump-color_user-multiple-accounts-flat.svg'
import riwayat from '../assets/fluent-color_history-28.svg'

// --- INISIALISASI STATE & ROUTER ---
const router = useRouter()
const user = ref({ nama: 'Mahasiswa' })

// --- LOGIKA NAVIGASI ---
const absen = () => router.push('/absensi')
const keAkun = () => router.push('/akun')
const keRiwayat = () => router.push('/riwayat')

// --- AMBIL DATA USER DARI STORAGE SAAT HALAMAN DIMUAT ---
onMounted(() => {
  // 1. Ambil key yang BENAR sesuai kodingan Login lu tadi!
  const savedUserNama = localStorage.getItem('user_nama')
  const savedRole = localStorage.getItem('role')

  // 2. Proteksi: Kita tadi udah ubah role jadi 'mahasiswa', bukan 'student'
  if (savedRole !== 'student') {
    alert('Akses ditolak! Anda bukan mahasiswa.')
    router.push('/login?role=mahasiswa')
    return
  }

  // 3. Load data hasil sukses login dari database Azure
  if (savedUserNama) {
    // Nggak perlu JSON.parse karena wujudnya udah pure text string
    user.value.nama = savedUserNama
  } else {
    // Kalau nggak ada nama di storage, lempar balik ke login
    router.push('/login?role=mahasiswa')
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
  background: #ea3236;
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
  font-size: 20px;
  margin-top: -3px;
  margin-bottom: 15px;
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
  text-align: left;
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
  font-size: 18px;
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
  width: 70px;
  height: 70px;
  stroke: black;
  fill: none;
  stroke-width: 3;
  margin-bottom: 5px;
}

.bottom {
  height: 35px;
  background: #f3f3f3;
}
</style>