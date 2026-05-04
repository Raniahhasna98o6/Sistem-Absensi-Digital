<template>
  <div class="wrapper">
    <div class="phone">

      <!-- CONTENT CENTER -->
      <div class="content">

        <!-- ICON -->
        <div class="check-circle">
          ✓
        </div>

        <!-- TEXT -->
        <h2>Absensi Berhasil</h2>

        <!-- BUTTON -->
        <button @click="kembali">
          Kembali ke Beranda
        </button>

      </div>

    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const kembali = () => {
  router.push('/beranda')
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

/* PHONE */
.phone {
  width: 390px;
  height: 800px;
  background: white;
  border-radius: 30px;
  overflow: hidden;
  display: flex;
}

/* CONTENT CENTER (INI KUNCI BIAR MIRIP MOCKUP) */
.content {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;  /* tengah vertikal */
  align-items: center;      /* tengah horizontal */
  gap: 25px;
  padding: 20px;
}

/* ICON BULAT */
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

/* TEXT */
h2 {
  font-size: 24px;
  font-weight: 700;
  color: black;
}

/* BUTTON */
button {
  background: #ff3b30;
  color: white;
  border: none;
  padding: 14px 20px;
  border-radius: 12px;
  font-weight: 700;
  font-size: 16px;
  cursor: pointer;

  /* biar lebar kayak mockup */
  width: 70%;
}

</style>