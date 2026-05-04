<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER -->
      <div class="header">
        <span class="back" @click="back">←</span>
        <h3>Verifikasi Foto</h3>
      </div>

      <!-- CONTENT -->
      <div class="content">

        <!-- FOTO -->
        <img :src="image" class="photo" />

        <!-- STATUS -->
        <div class="status">
          📍 Dalam Area kampus
        </div>

        <!-- INFO -->
        <div class="info-box">
          <div class="row">
            <div class="icon">📅</div>
            <div class="text">
              <p class="matkul">Jaringan Komputer</p>
              <p class="jam">08.10 WIB</p>
            </div>
          </div>
        </div>

        <!-- BUTTON -->
        <div class="buttons">
          <button class="ulang" @click="ulang">Ambil Ulang</button>
          <button class="kirim" @click="kirim">Kirim Absensi</button>
        </div>

      </div>

    </div>
  </div>
</template>

<script setup>
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

const image = route.query.img

const back = () => router.push('/kamera')
const ulang = () => router.push('/kamera')
const kirim = () => router.push('/success')
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

/* PHONE */
.phone {
  width: 390px;
  height: 800px;
  background: white;
  border-radius: 30px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

/* HEADER */
.header {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 16px;
  background: #f3f3f3;
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

/* CONTENT (SUPER RAPAT & PROPORSIONAL) */
.content {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 16px;
  gap: 12px; /* 🔥 lebih rapat */
}

/* FOTO (LEBIH PROPORSIONAL) */
.photo {
  width: 100%;
  height: 260px; /* 🔥 diturunin biar ga makan space */
  object-fit: cover;
  border-radius: 20px;
}

/* STATUS */
.status {
  background: #e9e9e9;
  padding: 12px;
  border-radius: 20px;
  font-weight: 600;
  color: #555;
  text-align: center;
}

/* INFO BOX (LEBIH BESAR & RAPI) */
.info-box {
  background: #f2f2f2;
  border-radius: 20px;
  padding: 16px;
  display: flex;
  align-items: center;
}

/* ROW */
.row {
  display: flex;
  align-items: center;
  gap: 14px;
}

/* ICON */
.icon {
  font-size: 30px;
}

/* TEXT */
.text {
  display: flex;
  flex-direction: column;
}

/* MATKUL */
.matkul {
  font-size: 18px;
  font-weight: 700;
  color: black;
}

/* JAM */
.jam {
  font-size: 15px;
  font-weight: 600;
  color: #444;
  margin-top: 2px;
}

/* BUTTON (NEMPEL BAWAH) */
.buttons {
  margin-top: auto; /* 🔥 kunci biar ga ada space kosong */
  display: flex;
  gap: 12px;
}

/* BUTTON ULANG */
.ulang {
  flex: 1;
  background: #ff3b30;
  color: white;
  border: none;
  padding: 14px;
  border-radius: 12px;
  font-weight: 700;
  font-size: 15px;
}

/* BUTTON KIRIM */
.kirim {
  flex: 1;
  background: #2f80ed;
  color: white;
  border: none;
  padding: 14px;
  border-radius: 12px;
  font-weight: 700;
  font-size: 15px;
}

</style>