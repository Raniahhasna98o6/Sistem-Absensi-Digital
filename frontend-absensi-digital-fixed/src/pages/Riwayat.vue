<template>
  <div class="wrapper">
    <div class="phone">
      
      <div class="topbar">
        <span @click="back">←</span>
        <h3>Riwayat Kehadiran</h3>
      </div>

      <div class="content">
        <div class="tabs">
          <button :class="{ active: active === 'semua' }" @click="active = 'semua'">Semua</button>
          <button :class="{ active: active === 'hadir' }" @click="active = 'hadir'">Hadir</button>
          <button :class="{ active: active === 'tidak' }" @click="active = 'tidak'">Tidak Hadir</button>
        </div>

        <div class="list-container">
          <div v-if="filteredData.length === 0" class="empty">
            <p>Belum ada riwayat absen.</p>
          </div>

          <div v-for="(item, index) in filteredData" :key="item.id_absensi || index" class="history-card">
            
            <div class="photo-container">
              <img v-if="item.foto_abs" :src="formatImageBase64(item.foto_abs)" alt="Foto Absen" class="absensi-img" />
              <div v-else class="no-photo">📷</div>
            </div>

            <div class="card-info">
              <p class="matkul">{{ item.nama_kelas || item.matkul }} ({{ item.kode_mk || '-' }})</p>
              <p class="detail">{{ formatTanggal(item.tanggal_abs || item.jam) }}</p>
              <p class="detail-loc">📍 {{ item.lokasi_abs || item.ruang }}</p>
            </div>

            <div :class="['status-badge', (item.status_abs || item.status).toLowerCase()]">
              {{ item.status_abs || (item.status === 'hadir' ? 'Hadir' : 'Tidak Hadir') }}
            </div>
          </div>
        </div>

      </div>

      <div class="bottom-spacer"></div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const active = ref('semua')
const data = ref([])

const back = () => router.push('/beranda')

// 1. FILTER DATA (Cukup satu kali aja)
const filteredData = computed(() => {
  if (active.value === 'semua') return data.value
  return data.value.filter(item => {
    const status = (item.status_abs || item.status).toLowerCase()
    return status === active.value
  })
})

// 2. FUNGSI RENDER GAMBAR (Biar gambarnya meledak muncul di layar)
const formatImageBase64 = (base64String) => {
  if (!base64String) return '';
  const cleanString = base64String.trim();
  if (cleanString.startsWith('data:image')) {
    return cleanString;
  }
  return `data:image/jpeg;base64,${cleanString}`;
}

// 3. FUNGSI FORMAT TANGGAL (Cukup satu kali aja)
const formatTanggal = (dateString) => {
  if (!dateString || (dateString.includes('-') && !dateString.includes('T') && !dateString.includes(':'))) return dateString
  try {
    const date = new Date(dateString)
    return new Intl.DateTimeFormat('id-ID', {
      weekday: 'short', day: 'numeric', month: 'short', hour: '2-digit', minute: '2-digit'
    }).format(date)
  } catch(e) { return dateString }
}

// 4. MENGAMBIL DATA DARI AZURE
onMounted(async () => {
  // FIX: ambil NIM dari localStorage, kirim via query param
  const nim = localStorage.getItem('user_nim')
  if (!nim) {
    router.push('/login?role=student')
    return
  }

  try {
    const baseURL = import.meta.env.VITE_API_URL || 'https://sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net'
    const cleanBaseURL = baseURL.replace(/\/$/, "")
    const response = await axios.get(`${cleanBaseURL}/api/absensi/riwayat?nim=${nim}`, {
      withCredentials: true
    })
    data.value = response.data || []
  } catch (error) {
    console.error('Gagal mengambil data:', error)
    data.value = []
  }
})
</script>

<style scoped>
/* CSS Tambahan untuk Foto */
.photo-container {
  width: 60px;
  height: 60px;
  border-radius: 12px;
  overflow: hidden;
  background: #e2e8f0;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
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

/* Penyesuaian Card Info */
.card-info {
  flex: 1;
  margin-left: 15px;
}

.detail-loc {
  font-size: 11px;
  color: #64748b;
  margin-top: 2px;
}

/* SISA CSS LU TETAP SAMA */
.wrapper { background: #0f172a; min-height: 100vh; display: flex; justify-content: center; align-items: center; }
.phone { width: 390px; height: 780px; background: #f8fafc; border-radius: 30px; overflow: hidden; display: flex; flex-direction: column; }
.topbar { padding: 20px; background: white; display: flex; align-items: center; gap: 15px; border-bottom: 1px solid #e2e8f0; }
.topbar span { font-size: 24px; cursor: pointer; color: #0f172a; }
.topbar h3 { font-weight: 700; font-size: 18px; color: #0f172a; margin: 0; }
.content { flex: 1; padding: 20px; display: flex; flex-direction: column; gap: 20px; overflow-y: auto; }
.tabs { display: flex; background: #e2e8f0; padding: 4px; border-radius: 12px; }
.tabs button { flex: 1; padding: 8px; border: none; border-radius: 8px; font-weight: 600; font-size: 13px; cursor: pointer; background: transparent; transition: 0.3s; }
.tabs button.active { background: white; box-shadow: 0 2px 4px rgba(0,0,0,0.1); color: #dc2626; }
.list-container { display: flex; flex-direction: column; gap: 12px; }
.history-card { background: white; padding: 12px; border-radius: 16px; display: flex; align-items: center; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05); }
.matkul { font-weight: 700; font-size: 14px; margin-bottom: 2px; }
.detail { font-size: 11px; color: #64748b; }
.status-badge { padding: 4px 8px; border-radius: 20px; font-size: 10px; font-weight: 700; text-transform: uppercase; }
.status-badge.hadir { background: #dcfce7; color: #166534; }
.status-badge.tidak { background: #fee2e2; color: #991b1b; }
.empty { text-align: center; color: #94a3b8; margin-top: 40px; }
.bottom-spacer { height: 20px; background: #f8fafc; }
</style>