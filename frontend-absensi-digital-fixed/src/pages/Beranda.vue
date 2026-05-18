<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER TOPBAR -->
      <div class="topbar">
        <div class="top-left">
          <div class="profile">
            <img src="https://i.pravatar.cc/150?img=12" alt="profile" width="40" style="border-radius: 50%;"/>
          </div>
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
                <p><b>Jaringan Komputer</b></p>
                <p>08.00 - 10.00</p>
              </div>
            </div>
            <button @click="absen">ABSEN SEKARANG</button>
          </div>
        </div>

        <!-- MENU FITUR UTAMA -->
        <div class="card">
          <h3 class="fitur-title">Fitur Website</h3>
          <div class="menu">
            <div class="menu-card" @click="keAkun">
              <img :src="akun" class="menu-icon"/>
              <p>Akun</p>
            </div>
            
            <div class="menu-card" @click="keRiwayat">
              <img :src="riwayat" class="menu-icon"/>
              <p>Riwayat Absensi</p>
            </div>
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
  margin-top: 15px;
  margin-bottom: -15px;
}

.top-left {
  display: flex;
  align-items: center;
  gap: 15px;
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
  margin-bottom: 20px;
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
  margin-top: 0;
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
  width: 60px;
  height: 60px;
  stroke: black;
  fill: none;
  stroke-width: 3;
}

.bottom {
  height: 35px;
  background: #f3f3f3;
}
</style>