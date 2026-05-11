<template>
  <div class="wrapper">
    <div class="phone">

      <!-- TOPBAR -->
      <div class="topbar">
        <span class="back" @click="goBack">←</span>
        <h3>Absensi Kuliah</h3>
      </div>

      <!-- RED AREA -->
      <div class="red">

        <!-- CARD 1 -->
        <div class="card">

          <h2 class="judul">Mata Kuliah Hari Ini</h2>
          <div class="divider"></div>

          <!-- MATKUL -->
          <div class="row">
            <div class="icon">📅</div>
            <div class="info">
              <p class="matkul">Jaringan Komputer</p>
              <p class="jam">08.00 - 10.00</p>
            </div>
          </div>

          <div class="divider"></div>

          <!-- LOKASI -->
          <div class="row">
            <div class="icon">📍</div>
            <p class="lokasi">TULT 0714</p>
          </div>

        </div>

        <!-- CARD 2 (MAPS) -->
        <div class="card">
          <!-- Ganti iframe dengan div ini -->
          <div id="map" class="map"></div>
        </div>

        <!-- CARD 3 -->
        <div class="card bottom-card">

          <div class="lokasi-box">
            📍 Lokasi Terdeteksi Dalam Area Kampus
          </div>

          <button @click="ambilFoto">AMBIL FOTO</button>

        </div>

      </div>

      <div class="bottom"></div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'

const goBack = () => {
  router.push('/beranda')
}

const router = useRouter()
const koordinat = ref({ lat: -6.974, lng: 107.630 }) // Default Telkom
let map = null
let blueDot = null
let watchId = null

const ambilFoto = () => {
  localStorage.setItem('active_class', JSON.stringify({
    matkul: 'Jaringan Komputer', // Data dummy sesuai tugas kamu
    lat: koordinat.value.lat,
    lng: koordinat.value.lng
  }))
  router.push('/kamera')
}

onMounted(() => {
  // 1. Inisialisasi Peta Leaflet
  map = L.map('map').setView([koordinat.value.lat, koordinat.value.lng], 17);
  
  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '© OpenStreetMap'
  }).addTo(map);

  // 2. Buat "Titik Biru" Kustom
  blueDot = L.circleMarker([koordinat.value.lat, koordinat.value.lng], {
    radius: 8,
    fillColor: '#2196F3', // Warna biru Google Maps
    color: 'white',
    weight: 2,
    fillOpacity: 1
  }).addTo(map);

  // 3. Pantau Lokasi Mahasiswa secara Real-Time
  if (navigator.geolocation) {
    watchId = navigator.geolocation.watchPosition((pos) => {
      const { latitude, longitude } = pos.coords
      koordinat.value = { lat: latitude, lng: longitude }

      // Update posisi titik biru dan fokus kamera peta
      blueDot.setLatLng([latitude, longitude])
      map.setView([latitude, longitude])
    }, (err) => {
      console.error("GPS Error:", err)
    }, { enableHighAccuracy: true });
  }
})

// Bersihkan sensor GPS saat pindah halaman agar baterai HP tidak boros
onUnmounted(() => {
  if (watchId) navigator.geolocation.clearWatch(watchId)
})
</script>

<style scoped>

/* BACKGROUND */
.wrapper {
  background: #0f172a;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

/* PHONE */
.phone {
  width: 390px;
  height: 780px;
  background: white;
  border-radius: 30px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.back {
  font-size: 22px;
  cursor: pointer;
  color: #000;
}

/* TOPBAR */
.topbar {
  padding: 15px;
  background: #f3f3f3;
  display: flex;
  align-items: center;
  gap: 10px;
}

.topbar span {
  font-size: 20px;
  cursor: pointer;
}

.topbar h3 {
  font-weight: bold;
  color: black;
  font-size: 18px;
}

/* RED AREA */
.red {
  background: #ff2d2d;
  padding: 20px;
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

/* CARD */
.card {
  background: #f1f1f1;
  border-radius: 20px;
  padding: 16px;
  box-shadow: 0 10px 20px rgba(0,0,0,0.25);
}

/* FIX JUDUL (INI YANG DIPERBAIKI) */
.judul {
  font-size: 22px;
  font-weight: 700;
  text-align: left;
  color: #000; /* hitam pekat */
  opacity: 1;
}

/* GARIS */
.divider {
  height: 1px;
  background: #ccc;
  margin: 10px 0;
}

/* ROW */
.row {
  display: flex;
  align-items: center;
  gap: 12px;
}

/* ICON */
.icon {
  font-size: 22px;
}

/* TEXT */
.matkul {
  font-weight: bold;
  font-size: 16px;
  color: black;
}

.jam {
  font-size: 14px;
  color: black;
}

.lokasi {
  font-size: 15px;
  font-weight: bold;
  color: black;
}

/* MAP */
.map {
  width: 100%;
  height: 200px;
  border-radius: 15px;
  z-index: 1;
  border: none;
}

/* CARD BAWAH */
.bottom-card {
  margin-top: -5px;
}

/* LOKASI BOX */
.lokasi-box {
  background: #c8e6c9;
  color: #1b5e20;
  padding: 12px;
  border-radius: 12px;
  text-align: center;
  font-weight: bold;
  margin-bottom: 12px;
}

/* BUTTON */
button {
  width: 100%;
  padding: 14px;
  background: #2f80ed;
  color: white;
  border: none;
  border-radius: 12px;
  font-weight: bold;
  cursor: pointer;
  position: relative;
  z-index: 10;
}

/* BOTTOM */
.bottom {
  height: 25px;
  background: #f3f3f3;
}

</style>