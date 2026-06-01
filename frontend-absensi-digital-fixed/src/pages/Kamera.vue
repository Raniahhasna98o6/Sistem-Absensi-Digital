<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER NAVIGASI -->
      <div class="header">
        <span class="back" @click="goBack">←</span>
        <h3>Ambil Foto Absensi</h3>
      </div>

      <!-- AREA VIEWPORT KAMERA -->
      <div class="camera-container">
        <video ref="video" autoplay playsinline></video>
        <!-- Overlay Bingkai Wajah -->
        <div class="frame"></div>
        <div class="hint">Pastikan wajah terlihat jelas</div>
      </div>

      <!-- TOMBOL SHUTTER -->
      <div class="bottom">
        <button class="shutter" @click="takePhoto"></button>
      </div>

      <!-- CANVAS TERSEMBUNYI UNTUK PROSES FOTO -->
      <canvas ref="canvas" style="display:none;"></canvas>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'

const video = ref(null)
const canvas = ref(null)
const router = useRouter()
let streamInstance = null 

// --- NAVIGASI KEMBALI ---
const goBack = () => {
  router.push('/absensi')
}

// --- INISIALISASI KAMERA SAAT HALAMAN DIBUKA ---
onMounted(async () => {
  try {
    const stream = await navigator.mediaDevices.getUserMedia({
      video: { facingMode: "user" } // Menggunakan kamera depan
    })
    streamInstance = stream 
    video.value.srcObject = stream
  } catch (err) {
    alert("Kamera tidak diizinkan / tidak tersedia")
  }
})

// --- MEMATIKAN STREAM KAMERA SAAT PINDAH HALAMAN ---
onUnmounted(() => {
  if (streamInstance) {
    streamInstance.getTracks().forEach(track => track.stop()) // Stop semua track video
  }
})

// --- PROSES PENGAMBILAN GAMBAR ---
const takePhoto = () => {
  const ctx = canvas.value.getContext('2d')
  
  // Kecilin resolusi kanvasnya biar Base64-nya nggak kegedean
  // (Misal dikecilin jadi setengahnya)
  canvas.value.width = video.value.videoWidth / 2
  canvas.value.height = video.value.videoHeight / 2

  ctx.drawImage(video.value, 0, 0, canvas.value.width, canvas.value.height)

  // Konversi pake kualitas yang agak diturunin (0.7)
  const image = canvas.value.toDataURL("image/png", 1.0)

  // Simpan foto
  localStorage.setItem('captured_photo', image)

  // Pastikan nama rutenya SAMA PERSIS dengan yang di router/index.js
  router.push('/preview') 
}
</script>

<style scoped>
/* --- KONFIGURASI LAYOUT UTAMA --- */
.wrapper {
  min-height: 100vh;
  background: #0f1c2e;
  display: flex;
  justify-content: center;
  align-items: center;
}

.phone {
  width: 390px;
  height: 780px;
  background: #f3f3f3;
  border-radius: 30px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

/* Efek Cermin buat Preview Video */
.camera-container video {
  transform: scaleX(-1);
}

/* --- TAMPILAN HEADER --- */
.header {
  margin-top: 25px;
  padding: 15px;
  background: #f3f3f3;
  align-items: center;
  display: flex;
  gap: 10px;
  border-bottom: 1px solid lightgray;
}

.header h3 {
  margin: 0;
  font-weight: bold;
  font-size: 18px;
  color: #000;
}

.back {
  font-size: 22px;
  cursor: pointer;
  color: #000;
}

/* --- TAMPILAN KAMERA & OVERLAY --- */
.camera-container {
  position: relative;
  flex: 1;
  background: black;
}

video {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.frame {
  position: absolute;
  top: 18%;
  left: 10%;
  width: 80%;
  height: 60%;
  border: 3px solid white;
  border-radius: 10px;
}

.hint {
  position: absolute;
  bottom: 10px;
  left: 50%;
  transform: translateX(-50%);
  background: rgba(255,255,255,0.9);
  padding: 10px 18px;
  border-radius: 20px;
  border: 1px solid #ea3236;
  font-size: 14px;
  color: #ea3236;
  white-space: nowrap;
  text-align: center;
}

/* --- KONTROL TOMBOL --- */
.bottom {
  border-top: 1px solid lightgray;
  padding: 30px;
  display: flex;
  justify-content: center;
}

.shutter {
  width: 70px;
  height: 70px;
  border-radius: 50%;
  background: lightgray;
  border: 6px solid #ccc;
  cursor: pointer;
}
</style>