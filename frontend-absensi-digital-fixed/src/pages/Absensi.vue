<template>
  <div class="wrapper">
    <div class="phone">

      <div class="topbar">
        <span class="back" @click="goBack">←</span>
        <h3>Absensi Kuliah</h3>
      </div>

      <div class="red">
        <div class="card">
          <h2 class="judul">Mata Kuliah Hari Ini</h2>
          <div class="divider"></div>

          <div class="row">
            <img :src="kalender" class="calendar-icon" />
            <div class="info">
              <p class="matkul">Jaringan Komputer</p>
              <p class="jam">08.00 - 10.00</p>
            </div>
          </div>

          <div class="divider"></div>

          <div class="row">
            <img :src="location" class="location-icon" />
            <p class="lokasi">TULT 0714</p>
          </div>

        </div>

        <div class="card">
          <div id="map" class="map"></div>
        </div>

        <div class="card bottom-card">
          <div :class="['lokasi-box', diLuarJangkauan ? 'luar' : 'dalam']">
            <img
              :src="diLuarJangkauan ? warning : location"
              class="lokasi-icon"
            />
            <span>
              {{ diLuarJangkauan 
              ? `Di Luar Jangkauan (${jarakMeter}m)` 
              : `Dalam Area Kampus (${jarakMeter}m)` }}
            </span>
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
import kalender from '../assets/marketeq_date.svg'
import location from '../assets/streamline-plump-color--location-pin-3.svg'
import warning from '../assets/fluent-emoji-flat--warning.svg'

const router = useRouter()
const goBack = () => router.push('/beranda')

// Titik Koordinat TULT (Pusat Kampus)
const TELYU_LAT = -6.974001
const TELYU_LNG = 107.630339
const RADIUS_MAKSIMAL = 800 // Ganti jadi 100 meter aja biar ketat

const koordinat = ref({ lat: TELYU_LAT, lng: TELYU_LNG }) 
const jarakMeter = ref(0)
const diLuarJangkauan = ref(false)

let map = null
let blueDot = null
let watchId = null

// RUMUS MENGHITUNG JARAK ANTARA 2 TITIK KOORDINAT BUMI (HAVERSINE)
const hitungJarak = (lat1, lon1, lat2, lon2) => {
  const R = 6371e3; // Radius bumi dalam meter
  const p1 = lat1 * Math.PI/180;
  const p2 = lat2 * Math.PI/180;
  const deltaP = (lat2 - lat1) * Math.PI/180;
  const deltaLon = (lon2 - lon1) * Math.PI/180;

  const a = Math.sin(deltaP/2) * Math.sin(deltaP/2) +
            Math.cos(p1) * Math.cos(p2) *
            Math.sin(deltaLon/2) * Math.sin(deltaLon/2);
  const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1-a));

  return R * c; // Hasil dalam meter
}

const ambilFoto = () => {
  localStorage.setItem('active_class', JSON.stringify({
    matkul: 'Jaringan Komputer',
    lat: koordinat.value.lat,
    lng: koordinat.value.lng,
    // FIX: Kirim status ini ke halaman Preview biar tombol Kirim ke-lock kalau dari rumah!
    diLuarJangkauan: diLuarJangkauan.value 
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
    fillColor: '#2196F3',
    color: 'white',
    weight: 2,
    fillOpacity: 1
  }).addTo(map);

  // 3. Pantau Lokasi Mahasiswa secara Real-Time
  if (navigator.geolocation) {
    watchId = navigator.geolocation.watchPosition((pos) => {
      const { latitude, longitude } = pos.coords
      koordinat.value = { lat: latitude, lng: longitude }

      // HITUNG JARAK KE KAMPUS SECARA REALTIME
      const jarak = hitungJarak(latitude, longitude, TELYU_LAT, TELYU_LNG)
      jarakMeter.value = Math.round(jarak)

      // UPDATE STATUS UI
      if (jarak > RADIUS_MAKSIMAL) {
        diLuarJangkauan.value = true
      } else {
        diLuarJangkauan.value = false
      }

      // Update posisi titik biru dan fokus kamera peta
      blueDot.setLatLng([latitude, longitude])
      map.setView([latitude, longitude])
    }, (err) => {
      console.error("GPS Error:", err)
    }, { enableHighAccuracy: true });
  }
})

onUnmounted(() => {
  if (watchId) navigator.geolocation.clearWatch(watchId)
})
</script>

<style scoped>
/* BACKGROUND */
.wrapper { background: #0f172a; min-height: 100vh; display: flex; justify-content: center; align-items: center; }
/* PHONE */
.phone { width: 390px; height: 780px; background: #f3f3f3; border-radius: 30px; overflow: hidden; display: flex; flex-direction: column; }
.back { font-size: 22px; cursor: pointer; color: #000; }
/* TOPBAR */
.topbar { 
  margin-top: 25px; 
  padding: 15px; 
  background: #f3f3f3; 
  display: flex; 
  align-items: center; 
  gap: 10px; }
.topbar span { font-size: 20px; cursor: pointer; }
.topbar h3 { 
  font-weight: bold; 
  color: black; 
  font-size: 18px; 
  margin: 0; }
/* RED AREA */
.red { background: #ff2d2d; padding: 20px; flex: 1; display: flex; flex-direction: column; gap: 14px; }
/* CARD */
.card { background: #f1f1f1; border-radius: 20px; padding: 16px; box-shadow: 0 10px 20px rgba(0,0,0,0.25); }
.judul { font-size: 20px; font-weight: 700; text-align: left; color: #000; margin: 0;}
/* GARIS */
.divider { height: 1px; background: #ccc; margin: 10px 0; }
/* ROW */
.row { display: flex; align-items: center; gap: 12px; }
.icon { font-size: 22px; }
.calendar-icon { width: 54px; height: 54px; }
.matkul { font-weight: bold; font-size: 16px; color: black; margin: 0; text-align: left;}
.jam { font-size: 14px; color: black; margin: 0; text-align: left;}
.lokasi { font-size: 15px; font-weight: bold; color: black; margin: 0;}
.lokasi-icon { width: 22px; height: 22px; object-fit: contain;}
.location-icon { width: 30px; height: 30px; }
/* MAP */
.map { width: 100%; height: 200px; border-radius: 15px; z-index: 1; border: none; }
/* CARD BAWAH */
.bottom-card { margin: 0; }

/* FIX: STYLE LOKASI BOX DINAMIS */
.lokasi-box {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px;
  border-radius: 12px;
  text-align: center;
  justify-content: center;
  font-weight: bold;
  margin-bottom: 12px;
  transition: background-color 0.3s;
}
.lokasi-box.dalam {
  background: #c8e6c9;
  color: #1b5e20;
}
.lokasi-box.luar {
  background: #ffcdd2;
  color: #b71c1c;
}

/* BUTTON */
button { width: 100%; padding: 14px; background: #2f80ed; color: white; border: none; border-radius: 12px; font-weight: bold; cursor: pointer; }
/* BOTTOM */
.bottom { height: 25px; background: #f3f3f3; }
</style>