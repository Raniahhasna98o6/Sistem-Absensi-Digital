<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER & NAVIGASI -->
      <div class="header">
        <span @click="back" style="cursor:pointer">←</span>
        <h3>Laporan Absensi</h3>
      </div>

      <div class="content">

        <!-- FILTER MATA KULIAH & TANGGAL -->
        <div class="filter">
          <select class="select">
            <option>Jaringan Komputer</option>
            <option>Basis Data</option>
          </select>
          <input type="date" class="date" />
        </div>

        <!-- RINGKASAN STATISTIK DARI DATABASE -->
        <div class="summary">
          <div>Total: {{ mahasiswa.length }}</div>
          <div>Hadir: {{ countHadir }}</div>
          <div>Tidak: {{ mahasiswa.length - countHadir }}</div>
        </div>

        <!-- DAFTAR MAHASISWA HASIL ABSENSI -->
        <div class="list">
          <div class="item" v-for="mhs in mahasiswa" :key="mhs.nama">
            <span class="nama">{{ mhs.nama }}</span>
            <span :class="mhs.status === 'Hadir' ? 'hadir' : 'tidak'">
              {{ mhs.status }}
            </span>
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
const mahasiswa = ref([])

// --- LOGIKA HITUNG RINGKASAN OTOMATIS ---
const countHadir = computed(() => {
  return mahasiswa.value.filter(m => m.status === 'Hadir').length
})

// --- LOGIKA NAVIGASI KEMBALI ---
const back = () => {
  router.push('/BerandaDosen')
}

// --- AMBIL DATA LAPORAN DARI AZURE ---
onMounted(async () => {
  try {
    const token = localStorage.getItem('token')
    // Menggunakan domain: sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net
    const response = await axios.get(`${import.meta.env.VITE_API_URL}/attendance/report`, {
      headers: { Authorization: `Bearer ${token}` }
    })
    
    mahasiswa.value = response.data 
  } catch (error) {
    console.error('Gagal mengambil laporan:', error)
    // Fallback data jika backend belum siap
    mahasiswa.value = [
      { nama: 'Andi Pratama', status: 'Hadir' },
      { nama: 'Budi Santoso', status: 'Tidak Hadir' },
      { nama: 'Citra Dewi', status: 'Hadir' }
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
  background: white;
  border-radius: 30px;
  overflow: hidden;
}

/* --- STYLE HEADER & FORM --- */
.header {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 15px;
  background: #f3f3f3;
}

.header h3 {
  font-weight: bold;
  color: black;
}

.content {
  padding: 16px;
}

.filter {
  display: flex;
  gap: 10px;
  margin-bottom: 15px;
}

select, input {
  flex: 1;
  padding: 10px;
  border-radius: 12px;
  border: none;
  background: #333;
  color: white;
}

/* --- STYLE RINGKASAN & ITEM LIST --- */
.summary {
  display: flex;
  justify-content: space-between;
  background: #eaeaea;
  padding: 14px;
  border-radius: 14px;
  margin-bottom: 15px;
  font-weight: bold;
  color: black;
}

.list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.item {
  display: flex;
  justify-content: space-between;
  background: #eaeaea;
  padding: 14px;
  border-radius: 14px;
}

.nama {
  color: black;
  font-weight: 600;
}

.hadir {
  color: green;
  font-weight: bold;
}

.tidak {
  color: red;
  font-weight: bold;
}
</style>