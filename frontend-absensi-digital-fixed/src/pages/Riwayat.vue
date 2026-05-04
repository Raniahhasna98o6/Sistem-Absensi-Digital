<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER & NAVIGASI -->
      <div class="header">
        <span class="back" @click="back">←</span>
        <h3>Riwayat Absensi</h3>
      </div>

      <!-- FILTER STATUS (SEMUA, HADIR, TIDAK) -->
      <div class="filter">
        <button 
          :class="['tab', active === 'semua' && 'active']"
          @click="active = 'semua'"
        >
          Semua
        </button>
        <button 
          :class="['tab', active === 'hadir' && 'active']"
          @click="active = 'hadir'"
        >
          Hadir
        </button>
        <button 
          :class="['tab', active === 'tidak' && 'active']"
          @click="active = 'tidak'"
        >
          Tidak Hadir
        </button>
      </div>

      <!-- LIST RIWAYAT HASIL FILTER -->
      <div class="content">
        <div 
          v-for="(item, index) in filteredData" 
          :key="index"
          class="card"
        >
          <div class="row">
            <div class="icon">📅</div>
            <div class="text">
              <p class="matkul">{{ item.matkul }}</p>
              <p class="ruang">{{ item.ruang }}</p>
              <p class="jam">{{ item.jam }}</p>
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
    const token = localStorage.getItem('token')
    // Menggunakan domain: sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net
    const response = await axios.get(`${import.meta.env.VITE_API_URL}/attendance/history`, {
      headers: { Authorization: `Bearer ${token}` }
    })
    data.value = response.data
  } catch (error) {
    console.error('Gagal mengambil data:', error)
    // Fallback data jika server offline
    data.value = [
      { matkul: 'Jaringan Komputer', ruang: 'TULT 0714', jam: '08.00 - 10.00', status: 'hadir' },
      { matkul: 'Basis Data', ruang: 'LAB 02', jam: '10.00 - 12.00', status: 'tidak' }
    ]
  }
})
</script>

<style scoped>
/* --- LAYOUT UTAMA --- */
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
  background: #f5f5f5;
  border-radius: 30px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

/* --- TAMPILAN HEADER --- */
.header {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 16px;
  background: white;
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

/* --- TAB FILTER --- */
.filter {
  display: flex;
  gap: 10px;
  padding: 14px 16px;
}

.tab {
  flex: 1;
  padding: 10px;
  border-radius: 12px;
  border: 1.5px solid #ff3b30;
  background: white;
  color: #ff3b30;
  font-weight: 600;
  cursor: pointer;
}

.tab.active {
  background: #ff3b30;
  color: white;
}

/* --- KONTEN & KARTU --- */
.content {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.card {
  background: #ff3b30;
  color: white;
  padding: 16px;
  border-radius: 18px;
  box-shadow: 0 6px 14px rgba(0,0,0,0.2);
}

.row {
  display: flex;
  align-items: flex-start;
  gap: 12px;
}

.icon {
  font-size: 30px;
}

.text {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.matkul {
  font-size: 16px;
  font-weight: 700;
}

.ruang, .jam {
  font-size: 14px;
  margin-top: 2px;
}
</style>