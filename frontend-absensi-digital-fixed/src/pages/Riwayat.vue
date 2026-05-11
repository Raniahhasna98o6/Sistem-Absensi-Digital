<template>
  <div class="wrapper">
    <div class="phone">
      
      <!-- TOPBAR -->
      <div class="topbar">
        <span @click="back">←</span>
        <h3>Riwayat Kehadiran</h3>
      </div>

      <div class="content">
        <!-- FILTER TABS -->
        <div class="tabs">
          <button 
            :class="{ active: active === 'semua' }" 
            @click="active = 'semua'"
          >Semua</button>
          <button 
            :class="{ active: active === 'hadir' }" 
            @click="active = 'hadir'"
          >Hadir</button>
          <button 
            :class="{ active: active === 'tidak' }" 
            @click="active = 'tidak'"
          >Tidak Hadir</button>
        </div>

        <!-- DAFTAR RIWAYAT -->
        <div class="list-container">
          <div v-if="filteredData.length === 0" class="empty">
            <p>Belum ada riwayat absen.</p>
          </div>

          <div 
            v-for="(item, index) in filteredData" 
            :key="index" 
            class="history-card"
          >
            <div class="card-info">
              <p class="matkul">{{ item.matkul }}</p>
              <p class="detail">{{ item.jam }} | {{ item.ruang }}</p>
            </div>
            <div :class="['status-badge', item.status]">
              {{ item.status === 'hadir' ? 'Hadir' : 'Tidak Hadir' }}
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

// --- LOGIKA NAVIGASI ---
const back = () => router.push('/beranda')

// --- LOGIKA FILTER DATA REAKTIF ---
const filteredData = computed(() => {
  if (active.value === 'semua') return data.value
  return data.value.filter(item => item.status === active.value)
})

// --- AMBIL DATA RIWAYAT DARI AZURE ---
onMounted(async () => {
  try {
    const baseURL = import.meta.env.VITE_API_URL || 'https://sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net'
    const cleanBaseURL = baseURL.replace(/\/$/, "")

    const response = await axios.get(`${cleanBaseURL}/api/absensi/riwayat`, {
      withCredentials: true // Menggunakan Cookie nim_user sesuai backend lu
    })
    data.value = response.data
  } catch (error) {
    console.error('Gagal mengambil data:', error)
    // Fallback data simulasi
    data.value = [
      { matkul: 'Jaringan Komputer', ruang: 'TULT 0714', jam: '08.00 - 10.00', status: 'hadir' },
      { matkul: 'Basis Data', ruang: 'LAB 02', jam: '10.00 - 12.00', status: 'tidak' },
      { matkul: 'Kecerdasan Buatan', ruang: 'TULT 0501', jam: '13.00 - 15.00', status: 'hadir' }
    ]
  }
})
</script>

<style scoped>
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
  background: #f8fafc;
  border-radius: 30px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.topbar {
  padding: 20px;
  background: white;
  display: flex;
  align-items: center;
  gap: 15px;
  border-bottom: 1px solid #e2e8f0;
}

.topbar span { font-size: 24px; cursor: pointer; color: #0f172a; }
.topbar h3 { font-weight: 700; font-size: 18px; color: #0f172a; margin: 0; }

.content {
  flex: 1;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 20px;
  overflow-y: auto;
}

/* TABS */
.tabs {
  display: flex;
  background: #e2e8f0;
  padding: 4px;
  border-radius: 12px;
}

.tabs button {
  flex: 1;
  padding: 8px;
  border: none;
  border-radius: 8px;
  font-weight: 600;
  font-size: 13px;
  cursor: pointer;
  background: transparent;
  transition: 0.3s;
}

.tabs button.active {
  background: white;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  color: #dc2626;
}

/* HISTORY CARD */
.list-container {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.history-card {
  background: white;
  padding: 16px;
  border-radius: 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
}

.matkul { font-weight: 700; font-size: 15px; margin-bottom: 4px; }
.detail { font-size: 12px; color: #64748b; }

.status-badge {
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 11px;
  font-weight: 700;
  text-transform: uppercase;
}

.status-badge.hadir { background: #dcfce7; color: #166534; }
.status-badge.tidak { background: #fee2e2; color: #991b1b; }

.empty { text-align: center; color: #94a3b8; margin-top: 40px; }
.bottom-spacer { height: 20px; background: #f8fafc; }
</style>