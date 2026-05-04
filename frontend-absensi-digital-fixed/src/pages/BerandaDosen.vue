<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER -->
      <div class="header">
        <div class="user-icon">👤</div>
        <h3 class="beranda">Beranda</h3>
      </div>

      <!-- RED AREA -->
      <div class="red">

        <h2 class="title">Halo, Dosen</h2>

        <!-- PROFILE CARD -->
        <div class="card">

          <div class="profile">
            <img src="https://i.pravatar.cc/100" class="avatar" />

            <div>
              <p class="nama">Kevin Sandy</p>
              <p class="nim">20113242298</p>
            </div>
          </div>

          <div class="info">
            <div class="row">
              <span>Program Studi</span>
              <span>S1 Informatika</span>
            </div>

            <div class="row">
              <span>Fakultas</span>
              <span>Informatika</span>
            </div>

            <div class="row">
              <span>Email</span>
              <span>sandy@dosen.com</span>
            </div>

            <div class="row">
              <span>No HP</span>
              <span>0812345678910</span>
            </div>
          </div>

        </div>

        <!-- FITUR -->
        <h3 class="fitur">Fitur Website</h3>

        <div class="laporan" @click="generate">
          ⏱️ Generate Laporan
        </div>

        <!-- LOGOUT -->
        <button class="logout" @click="logout">Logout</button>

      </div>

    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
const router = useRouter()

const logout = () => {
  router.push('/login')
}

const generate = () => {
  router.push('/laporan')
}
</script>

<style scoped>

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

/* HEADER */
.header {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 15px;
  background: #f3f3f3;
}

.user-icon {
  font-size: 22px;
}

/* 🔥 FIX WARNA BERANDA */
.beranda {
  color: black;
  font-weight: 700;
}

/* RED */
.red {
  background: #ff2d2d;
  padding: 20px;
  height: 100%;
  color: white;
}

/* 🔥 RATA KIRI */
.title {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 15px;
  text-align: left;
}

/* CARD */
.card {
  background: #f3f3f3;
  color: black;
  border-radius: 20px;
  padding: 16px;
  margin-bottom: 15px;
}

/* PROFILE */
.profile {
  display: flex;
  align-items: center;
  gap: 12px;
}

.avatar {
  width: 55px;
  height: 55px;
  border-radius: 50%;
}

.nama {
  font-weight: bold;
}

.nim {
  font-size: 13px;
}

/* INFO */
.info {
  margin-top: 10px;
}

.row {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid #ddd;
  padding: 6px 0;
}

/* 🔥 RATA KIRI */
.fitur {
  margin-top: 10px;
  text-align: left;
  font-weight: 600;
}

/* LAPORAN */
.laporan {
  background: white;
  color: black;
  padding: 14px;
  border-radius: 14px;
  margin-top: 10px;
  font-weight: 600;
  cursor: pointer;
}

/* LOGOUT */
.logout {
  margin-top: 15px;
  width: 100%;
  padding: 14px;
  border-radius: 14px;
  border: none;
  background: white;
  color: red;
  font-weight: bold;
  cursor: pointer;
}

</style>