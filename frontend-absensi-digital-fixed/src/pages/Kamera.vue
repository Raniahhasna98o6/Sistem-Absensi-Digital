<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER -->
      <div class="header">
        <span class="back" @click="goBack">←</span>
        <h3>Ambil Foto Absensi</h3>
      </div>

      <!-- CAMERA AREA -->
      <div class="camera-container">

        <!-- VIDEO -->
        <video ref="video" autoplay playsinline></video>

        <!-- FRAME -->
        <div class="frame"></div>

        <!-- TEXT FIX -->
        <div class="hint">
          Pastikan wajah terlihat jelas
        </div>

      </div>

      <!-- BUTTON -->
      <div class="bottom">
        <button class="shutter" @click="takePhoto"></button>
      </div>

      <canvas ref="canvas" style="display:none;"></canvas>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const video = ref(null)
const canvas = ref(null)
const router = useRouter()

const goBack = () => {
  router.push('/absensi')
}

onMounted(async () => {
  try {
    const stream = await navigator.mediaDevices.getUserMedia({
      video: { facingMode: "user" }
    })
    video.value.srcObject = stream
  } catch (err) {
    alert("Kamera tidak diizinkan / tidak tersedia")
  }
})

const takePhoto = () => {
  const ctx = canvas.value.getContext('2d')
  canvas.value.width = video.value.videoWidth
  canvas.value.height = video.value.videoHeight

  ctx.drawImage(video.value, 0, 0)

  const image = canvas.value.toDataURL("image/png")

  // 🔥 kirim ke preview
  router.push({
    path: '/preview',
    query: {
      img: image
    }
  })
}
</script>

<style scoped>

/* BACKGROUND */
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
  background: #fff;
  border-radius: 30px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

/* HEADER FIX */
.header {
  display: flex;
  align-items: center;
  padding: 18px;
  background: #f5f5f5;
}

.header h3 {
  margin-left: 10px;
  font-size: 18px;
  font-weight: 700; /* BOLD */
  color: #000; /* HITAM */
}

.back {
  font-size: 22px;
  cursor: pointer;
  color: #000;
}

/* CAMERA */
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

/* FRAME */
.frame {
  position: absolute;
  top: 20%;
  left: 10%;
  width: 80%;
  height: 40%;
  border: 3px solid white;
  border-radius: 10px;
}

/* 🔥 TEXT FIX (INI YANG PENTING) */
.hint {
  position: absolute;
  bottom: 120px;
  left: 50%;
  transform: translateX(-50%);
  background: rgba(255,255,255,0.9);
  padding: 10px 18px;
  border-radius: 20px;
  font-size: 14px;

  /* FIX UTAMA */
  white-space: nowrap; /* supaya tidak turun */
  text-align: center;
}

/* BUTTON */
.bottom {
  padding: 20px;
  display: flex;
  justify-content: center;
}

.shutter {
  width: 70px;
  height: 70px;
  border-radius: 50%;
  background: #e0e0e0;
  border: 6px solid #ccc;
  cursor: pointer;
}
</style>