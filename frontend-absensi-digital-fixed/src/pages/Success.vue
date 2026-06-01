<template>
  <div class="wrapper">
    <div class="phone">
      <div class="content">
        <!-- IKON SUKSES -->
        <div class="check-circle">✓</div>

        <!-- TEKS KONFIRMASI -->
        <h2>Absensi Berhasil</h2>

        <!-- TOMBOL DINAMIS -->
        <button @click="kembali">
          Kembali ke Beranda
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

// --- INISIALISASI ROUTER & STATE ---
const router = useRouter()
const userRole = ref('student')

// --- LOGIKA AMBIL ROLE SAAT MOUNT ---
onMounted(() => {
  // Mengambil role dari storage untuk menentukan tujuan navigasi
  userRole.value = localStorage.getItem('role') || 'student'
})

// --- NAVIGASI DINAMIS BERDASARKAN ROLE ---
const kembali = () => {
  // Redirect ke endpoint beranda yang sesuai dengan identitas user[cite: 9, 10]
  if (userRole.value === 'lecturer') {
    router.push('/beranda-dosen')
  } else {
    router.push('/beranda')
  }
}
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

/* --- KONTEN TENGAH --- */
.content {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 25px;
  padding: 20px;
}

/* --- ELEMEN VISUAL --- */
.check-circle {
  width: 120px;
  height: 120px;
  background: #27ae60;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  color: white;
  font-size: 60px;
  font-weight: bold;
}

h2 {
  font-size: 24px;
  font-weight: 700;
  color: black;
}

button {
  background: #ff3b30;
  color: white;
  border: none;
  padding: 14px 20px;
  border-radius: 12px;
  font-weight: 700;
  font-size: 16px;
  cursor: pointer;
  width: 70%;
}
</style>