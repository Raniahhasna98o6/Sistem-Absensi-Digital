<template>
  <div class="wrapper">
    <div class="phone">

      <div class="header">
        <div class="user-icon">👤</div>
        <h3 class="beranda">Beranda Dosen</h3>
      </div>

      <div class="red">
        <h2 class="title">Halo,<br> <b>{{ user.nama || 'Dosen' }}</b></h2>

        <div class="card">
          <div class="profile">
            <img src="https://i.pravatar.cc/100" class="avatar" />
            <div class="identitas">
              <p class="nama">{{ user.nama }}</p>
              <p class="nim">NIDN: {{ user.nidn }}</p>
            </div>
          </div>

          <div class="info">
            <div class="row">
              <span>Program Studi</span>
              <span>{{ user.prodi }}</span>
            </div>
            <div class="row">
              <span>Fakultas</span>
              <span>{{ user.fakultas }}</span>
            </div>
            <div class="row">
              <span>Email</span>
              <span>{{ user.email }}</span>
            </div>
            <div class="row">
              <span>No HP</span>
              <span>{{ user.nohp }}</span>
            </div>
          </div>
        </div>

        <h3 class="fitur">Fitur Website</h3>
        <div class="laporan" @click="generate">
          ⏱️ Generate Laporan Absensi
        </div>

        <button class="logout" @click="logout" :disabled="isLoggingOut">
          {{ isLoggingOut ? 'Keluar...' : 'Logout' }}
        </button>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const isLoggingOut = ref(false)

const user = ref({
  nama: 'Memuat...',
  nidn: 'Memuat API...',
  prodi: '-',
  fakultas: '-',
  email: '-',
  nohp: '-'
})

const logout = async () => {
  isLoggingOut.value = true
  try {
    const baseURL = import.meta.env.VITE_API_URL || 'https://sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net'
    const cleanBaseURL = baseURL.replace(/\/$/, "")

    await axios.post(`${cleanBaseURL}/logout`, {}, {
      withCredentials: true 
    })
  } catch (error) {
    console.error('Server error saat logout:', error)
  } finally {
    localStorage.clear()
    router.push('/role')
  }
}

const generate = () => {
  router.push('/laporan')
}

// --- FUNGSI AMBIL PROFIL DARI API ---
const ambilProfilDB = async () => {
  // Ambil NIDN dari localStorage (Disimpan saat dosen login)
  const nidnUser = localStorage.getItem('user_nidn') || ''

  try {
    const baseURL = import.meta.env.VITE_API_URL || 'https://sistemabsensi-emcyfabpgpcuhaf5.indonesiacentral-01.azurewebsites.net'
    const cleanBaseURL = baseURL.replace(/\/$/, "")

    // Tembak API yang baru dibikin di main.go, kirim nidn via query
    const response = await axios.get(`${cleanBaseURL}/api/dosen/profil?nidn=${nidnUser}`, {
      withCredentials: true
    })

    if(response.data) {
      user.value = {
        nama: response.data.nama,
        nidn: response.data.nidn,
        prodi: response.data.prodi || '-',
        fakultas: response.data.fakultas || '-',
        email: response.data.email || '-',
        nohp: response.data.nohp || '-'
      }
    }
  } catch (error) {
    console.error('Gagal mengambil profil dari database:', error)
    // Fallback kalau gagal nembak API
    const savedUserNama = localStorage.getItem('user_nama')
    if (savedUserNama) user.value.nama = savedUserNama
    user.value.prodi = 'Gagal memuat'
  }
}

onMounted(() => {
  const savedRole = localStorage.getItem('role')

  if (savedRole !== 'dosen' && savedRole !== 'lecturer') {
    alert('Akses ditolak! Halaman ini khusus Dosen.')
    router.push('/login?role=lecturer')
    return
  }

  // Panggil fungsi API pas halaman kebuka
  ambilProfilDB()
})
</script>

<style scoped>
.wrapper { min-height: 100vh; background: #0f1c2e; display: flex; justify-content: center; align-items: center; }
.phone { width: 390px; height: 780px; background: white; border-radius: 30px; overflow: hidden; display: flex; flex-direction: column; }
.header { display: flex; align-items: center; gap: 10px; padding: 15px; background: #f3f3f3; }
.beranda { color: black; font-weight: 700; margin: 0; }
.red { background: #ea3236; padding: 20px; flex: 1; color: white; display: flex; flex-direction: column;}
.title { font-size: 24px; font-weight:400; margin-bottom: 15px; text-align: left; }

.card { background: #f3f3f3; color: black; border-radius: 20px; padding: 16px; margin-bottom: 15px; box-shadow: 0 4px 10px rgba(0,0,0,0.1); }
.profile { display: flex; align-items: center; gap: 12px; margin-bottom: 15px; }
.avatar { width: 55px; height: 55px; border-radius: 50%; border: 2px solid #ddd; }
.identitas { display: flex; flex-direction: column; align-items: flex-start; }
.nama { font-weight: 800; font-size: 16px; margin: 0; }
.nim { font-size: 13px; color: #666; margin: 0; font-weight: 600; }

.info { display: flex; flex-direction: column; }
.row { display: flex; justify-content: space-between; border-bottom: 1px solid #e0e0e0; padding: 10px 0; font-size: 13px; }
.row:last-child { border-bottom: none; }
.row span:first-child { font-weight: 600; color: #555; }
.row span:last-child { font-weight: 700; text-align: right; color: #111; }

.fitur { margin-top: 10px; text-align: left; font-weight: 700; font-size: 16px;}
.laporan { background: white; color: black; padding: 15px; border-radius: 14px; margin-top: 10px; font-weight: 700; cursor: pointer; text-align: center; box-shadow: 0 4px 6px rgba(0,0,0,0.1); transition: 0.2s;}
.laporan:active { transform: scale(0.98); }

.logout { margin-top: auto; width: 100%; padding: 15px; border-radius: 14px; border: none; background: white; color: #d32f2f; font-weight: 800; font-size: 15px; cursor: pointer; transition: 0.2s; box-shadow: 0 4px 6px rgba(0,0,0,0.1);}
.logout:disabled { background: #ccc; color: #666; cursor: not-allowed; }
.logout:active { transform: scale(0.98); }
</style>