<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER -->
      <div class="header">
        <span @click="back">←</span>
        <h3>Laporan Absensi</h3>
      </div>

      <!-- CONTENT -->
      <div class="content">

        <!-- FILTER -->
        <div class="filter">
          <select class="select">
            <option>Jaringan Komputer</option>
            <option>Basis Data</option>
          </select>

          <input type="date" class="date" />
        </div>

        <!-- SUMMARY -->
        <div class="summary">
          <div>Total: 30</div>
          <div>Hadir: 25</div>
          <div>Tidak Hadir: 5</div>
        </div>

        <!-- LIST -->
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
import { useRouter } from 'vue-router'
const router = useRouter()

const back = () => {
  router.push('/beranda-dosen')
}

const mahasiswa = [
  { nama: 'Andi Pratama', status: 'Hadir' },
  { nama: 'Budi Santoso', status: 'Tidak Hadir' },
  { nama: 'Citra Dewi', status: 'Hadir' },
  { nama: 'Dewi Lestari', status: 'Hadir' },
  { nama: 'Eko Saputra', status: 'Tidak Hadir' }
]
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