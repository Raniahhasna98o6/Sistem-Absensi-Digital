<template>
  <div class="wrapper">
    <div class="phone">

      <div class="header">
        <span class="back" @click="ulang">←</span>
        <h3>Preview Foto</h3>
      </div>
      <div class="body">
        <div class="content">
          <div class = "border">
            <img :src="image" class="photo" />
          </div>
          
          <div :class="['lokasi-box', activeClass.diLuarJangkauan ? 'luar' : 'dalam']">
            <img
            :src="activeClass.diLuarJangkauan ? lokasiMerah : lokasiHijau"
            class="lokasi-icon"
            />
            <span>
              {{ activeClass.diLuarJangkauan
              ? 'Di Luar Jangkauan'
              : 'Dalam Area Kampus'
            }}
          </span>
        </div>
        
        <div class="info-box">
          <div class="row">
            <img :src="kalender" class="calendar-icon" />
            <div class="text">
              <p class="matkul">{{ activeClass.nama_kelas || activeClass.matkul || 'Mata Kuliah' }} ({{ activeClass.kode_mk || '-' }})</p>
              <p class="jam">{{ activeClass.jam_mulai || '08:00' }} - {{ activeClass.jam_selesai || '10:00' }}</p>
            </div>
          </div>
        </div>
        
        <div class="buttons">
          <button class="ulang" @click="ulang">Ambil Ulang</button>
          <button class="kirim" @click="kirim" :disabled="isSubmitting || activeClass.diLuarJangkauan">
            {{ isSubmitting ? 'Mengirim...' : 'Kirim Absensi' }}
          </button>
        </div>
        
      </div>
    </div>
  </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import kalender from '../assets/marketeq_date.svg'
import lokasiHijau from '../assets/streamline-plump-color--location-pin-3.svg'
import lokasiMerah from '../assets/fluent-emoji-flat--warning.svg'

const router = useRouter()
const image = ref('')
const activeClass = ref({})
const isSubmitting = ref(false)

const ulang = () => router.push('/kamera')

onMounted(() => {
  const capturedPhoto = localStorage.getItem('captured_photo')
  const savedClass = localStorage.getItem('active_class')

  if (capturedPhoto) {
    image.value = capturedPhoto
  } else {
    router.push('/kamera')
  }

  if (savedClass) {
    activeClass.value = JSON.parse(savedClass)
    if (!('diLuarJangkauan' in activeClass.value)) {
       activeClass.value.diLuarJangkauan = false
    }
  }
})

const kirim = async () => {
  if (!navigator.geolocation) {
    alert("Browser tidak support GPS!")
    return
  }

  // FIX: Ambil NIM dari localStorage yang disimpan saat login
  const nimUser = localStorage.getItem('user_nim')
  if (!nimUser) {
    alert("Sesi habis, silakan login ulang!")
    router.push('/login?role=student')
    return
  }

  isSubmitting.value = true

  navigator.geolocation.getCurrentPosition(
    async (position) => {
      try {
        const latAsli = position.coords.latitude
        const lngAsli = position.coords.longitude

        const baseURL = import.meta.env.VITE_API_URL || 'https://sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net'
        const cleanBaseURL = baseURL.replace(/\/$/, "")

        const payload = {
          nim: nimUser,                              // FIX: dari localStorage, bukan response yang tidak ada
          foto_abs: image.value,
          lokasi_abs: activeClass.value.ruangan || 'TULT',
          status_abs: 'Hadir',
          latitude: latAsli,
          longitude: lngAsli
        }

        await axios.post(`${cleanBaseURL}/api/absensi`, payload, {
          withCredentials: true
        })

        localStorage.removeItem('captured_photo')
        localStorage.removeItem('active_class')
        router.push('/success')

      } catch (error) {
        console.error('Gagal kirim absen:', error)
        alert("Gagal: " + (error.response?.data?.message || error.message))
      } finally {
        isSubmitting.value = false
      }
    },
    (error) => {
      alert("Gagal ambil lokasi! Pastikan GPS nyala.")
      isSubmitting.value = false
    },
    { enableHighAccuracy: true, timeout: 10000, maximumAge: 0 }
  )
}
</script>

<style scoped>
.border{
  transform: scaleX(-1);
}
.wrapper { min-height: 100vh; background: #0f1c2e; display: flex; justify-content: center; align-items: center; }
.header { margin-top: 25px;
  padding: 15px;
  background: #f3f3f3;
  align-items: center;
  display: flex;
  gap: 10px;
  border-bottom: 1px solid lightgray;
}
.header h3 { margin: 0;
  font-weight: bold;
  font-size: 18px;
  color: #000; }
.body {
  background: #fff;
}
.back { font-size: 20px; cursor: pointer; color: black; }
.content { flex: 1; display: flex; flex-direction: column; padding: 16px; gap: 12px; }
.photo { width: 100%; height: 400px; object-fit: cover; border-radius: 20px; }
.status { padding: 12px; border-radius: 20px; font-weight: 600; text-align: center; }
.status-aman { background: #e9e9e9; color: #555; }
.status-bahaya { background: #ffebee; color: #c62828; }
.info-box { margin-top: -10px; background: #f1f1f1; border-radius: 20px; padding: 4px; display: flex; align-items: center; }
.row { display: flex; align-items: center; gap: 14px;}
.icon { font-size: 30px; }
.text { display: flex; flex-direction: column; }
.matkul { font-size: 18px; font-weight: 700; color: black; margin: 0; text-align: left;}
.jam { font-size: 15px; font-weight: 600; color: #444; margin-top: 2px; margin-bottom: 0; text-align: left;}
.buttons { margin-top: auto; display: flex; gap: 12px; }
.ulang { flex: 1; background: #ea3236; color: white; border: none; padding: 14px; border-radius: 12px; font-weight: 700; font-size: 15px; cursor: pointer; }
.kirim { flex: 1; background: #2f80ed; color: white; border: none; padding: 14px; border-radius: 12px; font-weight: 700; font-size: 15px; cursor: pointer; transition: 0.2s; }
.kirim:disabled { background: #9e9e9e; cursor: not-allowed; }
.lokasi-icon { width: 24px; height: 24px; }
.calendar-icon { width: 54px; height: 54px; }

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
</style>
