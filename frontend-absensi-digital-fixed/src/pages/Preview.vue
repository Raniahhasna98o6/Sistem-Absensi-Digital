<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER NAVIGASI -->
      <div class="header">
        <span class="back" @click="back">←</span>
        <h3>Verifikasi Foto</h3>
      </div>

      <div class="content">
        <!-- TAMPILAN HASIL FOTO KAMERA -->
        <img :src="image" class="photo" />

        <!-- STATUS GEOLOKASI (GEOFENCING) -->
        <div class="status">
          📍 Dalam Area Kampus
        </div>

        <!-- INFORMASI MATA KULIAH AKTIF -->
        <div class="info-box">
          <div class="row">
            <div class="icon">📅</div>
            <div class="text">
              <!-- Data diambil dinamis dari database Azure -->
              <p class="matkul">{{ activeClass.matkul || 'Mata Kuliah' }}</p>
              <p class="jam">{{ activeClass.jam || '--:--' }} WIB</p>
            </div>
          </div>
        </div>

        <!-- AKSI: ULANG ATAU KIRIM KE SERVER -->
        <div class="buttons">
          <button class="ulang" @click="ulang">Ambil Ulang</button>
          <button class="kirim" @click="kirim">Kirim Absensi</button>
        </div>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const image = ref('')
const activeClass = ref({})

// --- NAVIGASI KEMBALI KE KAMERA ---
const back = () => router.push('/kamera')
const ulang = () => router.push('/kamera')

// --- LOAD DATA DARI STORAGE SAAT HALAMAN DIBUKA ---
onMounted(() => {
  const capturedPhoto = localStorage.getItem('captured_photo')
  const savedClass = localStorage.getItem('active_class')

  if (capturedPhoto) {
    image.value = capturedPhoto
  } else {
    router.push('/kamera') // Proteksi jika tidak ada foto
  }

  if (savedClass) {
    activeClass.value = JSON.parse(savedClass)
  }
})

// --- KIRIM DATA ABSENSI KE BACKEND AZURE ---
// --- KIRIM DATA ABSENSI KE BACKEND AZURE DENGAN GPS ASLI ---
const kirim = async () => {
  // 1. Cek apakah browser support GPS
  if (!navigator.geolocation) {
    alert("Waduh, browser lu nggak support fitur lokasi (GPS), Sean!")
    return
  }

  // Tampilkan loading kalau perlu, karena ambil posisi GPS butuh waktu 1-3 detik
  console.log("Sedang mengambil lokasi presisi...")

  navigator.geolocation.getCurrentPosition(
    async (position) => {
      try {
        // Ambil koordinat asli dari hardware GPS
        const latAsli = position.coords.latitude
        const lngAsli = position.coords.longitude

        const baseURL = import.meta.env.VITE_API_URL || 'https://sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net'
        const cleanBaseURL = baseURL.replace(/\/$/, "")

        // 2. Kirim ke API Golang dengan data GPS yang akurat
        const response = await axios.post(`${cleanBaseURL}/api/absensi`, {
          image: image.value,
          matkul: activeClass.value.matkul,
          lat: latAsli, // Dinamis dari GPS
          lng: lngAsli  // Dinamis dari GPS
        }, {
          withCredentials: true 
        })

        // Bersihkan storage dan pindah ke halaman sukses
        localStorage.removeItem('captured_photo')
        localStorage.removeItem('active_class')
        router.push('/success')

      } catch (error) {
        console.error('Gagal kirim absen:', error)
        // Kalau gagal karena radius, pesan "Gagal! Anda berada di luar radius kampus!" bakal muncul di sini
        alert(error.response?.data?.message || 'Terjadi kesalahan pada server Azure.')
      }
    },
    (error) => {
      // Handle jika user menolak (Deny) izin lokasi
      console.error('Error GPS:', error)
      alert("Gagal ambil lokasi! Pastikan GPS HP nyala dan kasih izin 'Allow' di browser.")
    },
    {
      enableHighAccuracy: true, // Pakai GPS beneran, bukan cuma sinyal tower
      timeout: 5000,            // Maksimal nunggu 5 detik
      maximumAge: 0             // Jangan pakai lokasi lama yang tersimpan di cache
    }
  )
}
</script>

<style scoped>
/* --- KONFIGURASI LAYOUT & WRAPPER --- */
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
  display: flex;
  flex-direction: column;
}

/* --- STYLE KOMPONEN KONTEN --- */
.header {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 16px;
  background: #f3f3f3;
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

.content {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 16px;
  gap: 12px;
}

.photo {
  width: 100%;
  height: 260px;
  object-fit: cover;
  border-radius: 20px;
}

.status {
  background: #e9e9e9;
  padding: 12px;
  border-radius: 20px;
  font-weight: 600;
  color: #555;
  text-align: center;
}

.info-box {
  background: #f2f2f2;
  border-radius: 20px;
  padding: 16px;
  display: flex;
  align-items: center;
}

.row {
  display: flex;
  align-items: center;
  gap: 14px;
}

.icon {
  font-size: 30px;
}

.text {
  display: flex;
  flex-direction: column;
}

.matkul {
  font-size: 18px;
  font-weight: 700;
  color: black;
}

.jam {
  font-size: 15px;
  font-weight: 600;
  color: #444;
  margin-top: 2px;
}

/* --- STYLE TOMBOL AKSI --- */
.buttons {
  margin-top: auto;
  display: flex;
  gap: 12px;
}

.ulang {
  flex: 1;
  background: #ff3b30;
  color: white;
  border: none;
  padding: 14px;
  border-radius: 12px;
  font-weight: 700;
  font-size: 15px;
  cursor: pointer;
}

.kirim {
  flex: 1;
  background: #2f80ed;
  color: white;
  border: none;
  padding: 14px;
  border-radius: 12px;
  font-weight: 700;
  font-size: 15px;
  cursor: pointer;
}
</style>