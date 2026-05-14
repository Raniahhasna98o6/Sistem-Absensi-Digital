<template>
  <div class="wrapper">
    <div class="phone">

      <div class="header">
        <span @click="back" style="cursor:pointer">←</span>
        <h3>Laporan Absensi</h3>
      </div>

      <div class="content">

        <div class="filter">
          <select class="select" v-model="periode">
            <option value="2026-05">Mei 2026</option>
            <option value="2026-04">April 2026</option>
            <option value="2026-03">Maret 2026</option>
          </select>
          <select class="select" v-model="idKelas">
            <option value="1">IF-46-01</option>
            <option value="2">IF-46-02</option>
          </select>
        </div>

        <button class="btn-generate" @click="ambilLaporan">Tampilkan Laporan</button>

        <div class="summary">
          <div>Total: {{ laporan.length }}</div>
          <div>Hadir: {{ countHadir }}</div>
          <div>Tidak: {{ laporan.length - countHadir }}</div>
        </div>

        <div class="list">
          <div v-if="laporan.length === 0" class="empty">Belum ada data laporan.</div>
          <div class="item" v-for="(item, i) in laporan" :key="i">
            <span class="nama">{{ item.tanggal_abs }}</span>
            <span :class="item.status_abs === 'Hadir' ? 'hadir' : 'tidak'">
              {{ item.status_abs }}
            </span>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const laporan = ref([])
const periode = ref('2026-05')
const idKelas = ref('1')

const countHadir = computed(() => laporan.value.filter(m => m.status_abs === 'Hadir').length)

const back = () => router.push('/BerandaDosen')

// FIX: Pakai endpoint yang benar sesuai main.go, kirim nidn via query param
const ambilLaporan = async () => {
  const nidn = localStorage.getItem('user_nidn')
  if (!nidn) {
    alert('Sesi habis, silakan login ulang!')
    router.push('/login?role=lecturer')
    return
  }

  try {
    const baseURL = import.meta.env.VITE_API_URL || 'https://sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net'
    const cleanBaseURL = baseURL.replace(/\/$/, "")

    const response = await axios.get(`${cleanBaseURL}/api/dosen/laporan`, {
      params: { periode: periode.value, id_kelas: idKelas.value },
      withCredentials: true
    })

    laporan.value = response.data || []
  } catch (error) {
    console.error('Gagal mengambil laporan:', error)
    if (error.response?.status === 404) {
      laporan.value = []
      alert('Tidak ada data absensi untuk periode ini.')
    } else {
      alert('Gagal mengambil laporan: ' + (error.response?.data?.message || error.message))
    }
  }
}
</script>

<style scoped>
.wrapper { min-height: 100vh; background: #0f1c2e; display: flex; justify-content: center; align-items: center; }
.phone { width: 390px; height: 800px; background: white; border-radius: 30px; overflow: hidden; display: flex; flex-direction: column; }
.header { display: flex; align-items: center; gap: 10px; padding: 15px; background: #f3f3f3; }
.header h3 { font-weight: bold; color: black; }
.content { padding: 16px; display: flex; flex-direction: column; gap: 12px; overflow-y: auto; flex: 1; }
.filter { display: flex; gap: 10px; }
.select { flex: 1; padding: 10px; border-radius: 12px; border: none; background: #333; color: white; }
.btn-generate { padding: 12px; background: #2f80ed; color: white; border: none; border-radius: 12px; font-weight: bold; cursor: pointer; }
.summary { display: flex; justify-content: space-between; background: #eaeaea; padding: 14px; border-radius: 14px; font-weight: bold; color: black; }
.list { display: flex; flex-direction: column; gap: 12px; }
.item { display: flex; justify-content: space-between; background: #eaeaea; padding: 14px; border-radius: 14px; }
.nama { color: black; font-weight: 600; font-size: 13px; }
.hadir { color: green; font-weight: bold; }
.tidak { color: red; font-weight: bold; }
.empty { text-align: center; color: #94a3b8; margin-top: 20px; }
</style>
