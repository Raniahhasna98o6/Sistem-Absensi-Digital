<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER -->
      <div class="header">
        <span class="back" @click="back">←</span>
        <h3>Akun Saya</h3>
      </div>

      <!-- CONTENT -->
      <div class="content">

        <!-- FOTO -->
        <div class="profile">
          <img src="https://i.pravatar.cc/150?img=12" alt="profile" />
        </div>

        <!-- CARD NAMA -->
        <div class="card nama-card">
          <h2>Sean Benedict</h2>
          <p>220123455</p>
        </div>

        <!-- CARD DETAIL -->
        <div class="card detail-card">

          <div class="row">
            <span>Program Studi</span>
            <span>S1 Informatika</span>
          </div>

          <div class="divider"></div>

          <div class="row">
            <span>Fakultas</span>
            <span>Informatika</span>
          </div>

          <div class="divider"></div>

          <div class="row">
            <span>Angkatan</span>
            <span>2024</span>
          </div>

          <div class="divider"></div>

          <div class="row">
            <span>Email</span>
            <span>sean@student.com</span>
          </div>

          <div class="divider"></div>

          <div class="row">
            <span>No HP</span>
            <span>0812345678910</span>
          </div>

        </div>

        <!-- LOGOUT -->
        <button class="logout" @click="logout">
          Logout
        </button>

      </div>

    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const back = () => {
  router.push('/beranda')
}

const logout = () => {
  router.push('/login')
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
  background: #f5f5f5;
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
  background: #ffffff;
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

/* CONTENT */
.content {
  flex: 1;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  align-items: center;
}

/* FOTO */
.profile {
  margin-top: 10px;
}

.profile img {
  width: 110px;
  height: 110px;
  border-radius: 50%;
  object-fit: cover;
  border: 4px solid white;
  box-shadow: 0 6px 12px rgba(0,0,0,0.2);
}

/* 🔥 CARD FIX UTAMA */
.card {
  width: 100%;
  max-width: 330px; /* 🔥 BATASIN BIAR GA NEMPEL */
  background: white;
  border-radius: 18px;
  padding: 20px;
  box-shadow: 0 10px 18px rgba(0,0,0,0.12);
  border: 1px solid #ff4d4f;
}

/* NAMA CARD */
.nama-card {
  text-align: center;
}

.nama-card h2 {
  font-size: 20px;
  font-weight: 700;
  color: black;
}

.nama-card p {
  font-size: 16px;
  color: black;
}

/* DETAIL */
.detail-card {
  font-size: 15px;
}

/* ROW */
.row {
  display: flex;
  justify-content: space-between;
  padding: 10px 6px;
  color: black;
}

/* DIVIDER */
.divider {
  height: 1px;
  background: #ddd;
}

/* LOGOUT */
.logout {
  width: 100%;
  max-width: 330px; /* 🔥 BIAR SEJAJAR CARD */
  padding: 16px;
  background: white;
  border-radius: 16px;
  border: 2px solid #ff3b30;
  color: #ff3b30;
  font-weight: 700;
  font-size: 20px;
  margin-top: 20px;
}

</style>