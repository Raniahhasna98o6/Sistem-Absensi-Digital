<template>
  <div class="wrapper">
    <div class="phone">

      <div class="header">
        <span @click="back" class="back-btn">←</span>
        <h3>Laporan Absensi</h3>
      </div>

      <div class="content">

        <div class="filter">
          <input type="date" class="select date-picker" v-model="tanggalVal" @change="ambilLaporan" />
          
          <select class="select" v-model="idKelas" @change="ambilLaporan">
            <option v-for="kelas in listKelas" :key="kelas.id" :value="kelas.id">
              {{ kelas.nama }}
            </option>
          </select>
        </div>

        <button class="btn-generate" @click="ambilLaporan" :disabled="isLoading">
          {{ isLoading ? 'Memuat...' : 'Tampilkan Laporan' }}
        </button>

        <div class="summary">
          <div class="stat-box">
            <span class="stat-label">Total</span>
            <span class="stat-value">{{ laporan.length }}</span>
          </div>
          <div class="stat-box">
            <span class="stat-label text-green">Hadir</span>
            <span class="stat-value">{{ countHadir }}</span>
          </div>
          <div class="stat-box">
            <span class="stat-label text-red">Tidak</span>
            <span class="stat-value">{{ laporan.length - countHadir }}</span>
          </div>
        </div>

        <div class="list">
          <div v-if="laporan.length === 0" class="empty">
            <div class="empty-icon">📅</div>
            Belum ada data absensi di tanggal ini.
          </div>
          
          <div class="item" v-for="(item, i) in laporan" :key="i">
            
            <div class="photo-container">
              <img v-if="item.foto_abs" :src="formatImageBase64(item.foto_abs)" alt="Foto Absen" class="absensi-img" />
              <div v-else class="no-photo">📷</div>
            </div>

            <div class="info-kiri">
              <span class="waktu">{{ item.nama_mhs }}</span>
              <span class="lokasi">📍 {{ item.lokasi_abs || 'TULT' }} | ⏰ {{ formatJam(item.tanggal_abs) }}</span>
            </div>
            
            <div class="info-kanan">
              <span :class="['badge', item.status_abs === 'Hadir' ? 'hadir' : 'tidak']">
                {{ item.status_abs }}
              </span>
            </div>
            
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const laporan = ref([])
const listKelas = ref([]) 
const idKelas = ref('') 
const isLoading = ref(false)

const today = new Date().toISOString().split('T')[0]
const tanggalVal = ref(today)

const countHadir = computed(() => laporan.value.filter(m => m.status_abs === 'Hadir').length)

const back = () => router.push('/BerandaDosen')

const formatJam = (datetime) => {
  if (!datetime) return '-'
  return datetime.split(' ')[1] || datetime
}

const formatImageBase64 = (base64String) => {
  if (!base64String) return '';
  const cleanString = base64String.trim();
  if (cleanString.startsWith('data:image')) {
    return cleanString;
  }
  return `data:image/jpeg;base64,${cleanString}`;
}

const ambilLaporan = async () => {
  if (!idKelas.value) return; // Jangan nembak API kalau kelas belum ada

  const nidn = localStorage.getItem('user_nidn')
  if (!nidn) {
    alert('Sesi habis, silakan login ulang!')
    router.push('/login?role=lecturer')
    return
  }

  isLoading.value = true
  try {
    const baseURL = import.meta.env.VITE_API_URL || 'https://sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net'
    const cleanBaseURL = baseURL.replace(/\/$/, "")

    const response = await axios.get(`${cleanBaseURL}/api/dosen/laporan`, {
      params: { 
        periode: tanggalVal.value, 
        id_kelas: idKelas.value,
        nidn: nidn // Cadangan NIDN buat di HP
      },
      withCredentials: true
    })

    laporan.value = response.data || []
  } catch (error) {
    console.error('Gagal mengambil laporan:', error)
    if (error.response?.status === 404) {
      laporan.value = [] 
    } else {
      alert('Gagal mengambil laporan: ' + (error.response?.data?.message || error.message))
    }
  } finally {
    isLoading.value = false
  }
}

const ambilDaftarKelas = async () => {
  const nidn = localStorage.getItem('user_nidn')
  try {
    const baseURL = import.meta.env.VITE_API_URL || 'https://sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net'
    const cleanBaseURL = baseURL.replace(/\/$/, "")
    
    const response = await axios.get(`${cleanBaseURL}/api/dosen/daftar-kelas?nidn=${nidn}`)
    
    listKelas.value = response.data
    
    if (listKelas.value.length > 0) {
      idKelas.value = listKelas.value[0].id
      ambilLaporan() 
    }
  } catch (error) {
    console.error("Gagal load daftar kelas:", error)
  }
}

onMounted(() => {
  ambilDaftarKelas() 
})
</script>

<style scoped>
.wrapper { min-height: 100vh; background: #0f1c2e; display: flex; justify-content: center; align-items: center; font-family: sans-serif; }
.phone { width: 390px; height: 800px; background: #f8f9fa; border-radius: 30px; overflow: hidden; display: flex; flex-direction: column; box-shadow: 0 10px 25px rgba(0,0,0,0.5); }
.header { display: flex; align-items: center; gap: 15px; padding: 20px; background: white; border-bottom: 1px solid #eee; }
.header h3 { font-weight: 700; color: #1e293b; margin: 0; font-size: 18px; }
.back-btn { font-size: 24px; cursor: pointer; color: #64748b; font-weight: bold; }
.content { padding: 20px; display: flex; flex-direction: column; gap: 16px; overflow-y: auto; flex: 1; }

.filter { display: flex; gap: 10px; }
.select { flex: 1; padding: 12px; border-radius: 12px; border: 1px solid #cbd5e1; background: white; color: #334155; font-size: 14px; font-weight: 500; outline: none; }
.date-picker { font-family: inherit; }

.btn-generate { padding: 14px; background: #2f80ed; color: white; border: none; border-radius: 12px; font-weight: bold; font-size: 15px; cursor: pointer; transition: 0.2s; }
.btn-generate:disabled { background: #94a3b8; cursor: not-allowed; }
.btn-generate:active { transform: scale(0.98); }

.summary { display: flex; justify-content: space-between; background: white; padding: 16px; border-radius: 16px; border: 1px solid #e2e8f0; }
.stat-box { display: flex; flex-direction: column; align-items: center; gap: 4px; width: 30%; }
.stat-label { font-size: 12px; color: #64748b; font-weight: 600; }
.stat-value { font-size: 20px; font-weight: 800; color: #1e293b; }
.text-green { color: #10b981; }
.text-red { color: #ef4444; }

.list { display: flex; flex-direction: column; gap: 12px; padding-bottom: 20px; }
.empty { text-align: center; color: #94a3b8; margin-top: 40px; font-weight: 500; display: flex; flex-direction: column; gap: 10px; }
.empty-icon { font-size: 40px; }

.item { display: flex; justify-content: space-between; align-items: center; background: white; padding: 16px; border-radius: 16px; border: 1px solid #e2e8f0; transition: 0.2s; }
.item:hover { border-color: #cbd5e1; box-shadow: 0 4px 6px rgba(0,0,0,0.05); }

.badge { padding: 6px 12px; border-radius: 20px; font-size: 12px; font-weight: 700; }
.hadir { background: #dcfce7; color: #166534; }
.tidak { background: #fee2e2; color: #991b1b; }

.photo-container {
  width: 50px;
  height: 50px;
  border-radius: 12px;
  overflow: hidden;
  background: #e2e8f0;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  margin-right: 12px;
}

.absensi-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.no-photo {
  font-size: 20px;
  opacity: 0.5;
}

.info-kiri {
  display: flex;
  flex-direction: column;
  gap: 4px;
  flex: 1; 
}
.waktu { color: #1e293b; font-weight: 700; font-size: 15px; }
.lokasi { color: #64748b; font-size: 12px; font-weight: 500; }
</style>