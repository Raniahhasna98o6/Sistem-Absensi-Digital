<template>
  <div class="wrapper">
    <div class="phone">

      <!-- HEADER -->
      <div class="header">
        <span class="back" @click="back">←</span>
        <h3>Riwayat Absensi</h3>
      </div>

      <!-- FILTER -->
      <div class="filter">
        <button 
          :class="['tab', active === 'semua' && 'active']"
          @click="active = 'semua'"
        >
          Semua
        </button>

        <button 
          :class="['tab', active === 'hadir' && 'active']"
          @click="active = 'hadir'"
        >
          Hadir
        </button>

        <button 
          :class="['tab', active === 'tidak' && 'active']"
          @click="active = 'tidak'"
        >
          Tidak Hadir
        </button>
      </div>

      <!-- CONTENT -->
      <div class="content">

        <div 
          v-for="(item, index) in filteredData" 
          :key="index"
          class="card"
        >
          <div class="row">
            <div class="icon">📅</div>

            <div class="text">
              <p class="matkul">{{ item.matkul }}</p>
              <p class="ruang">{{ item.ruang }}</p>
              <p class="jam">{{ item.jam }}</p>
            </div>
          </div>
        </div>

      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const back = () => router.push('/beranda')

const data = ref([
  {
    matkul: 'Jaringan Komputer',
    ruang: 'TULT 0714',
    jam: '08.00 - 10.00',
    status: 'hadir'
  },
  {
    matkul: 'Basis Data',
    ruang: 'LAB 02',
    jam: '10.00 - 12.00',
    status: 'tidak'
  }
])

const active = ref('semua')

const filteredData = computed(() => {
  if (active.value === 'semua') return data.value
  return data.value.filter(item => item.status === active.value)
})
</script>

<style scoped>

/* WRAPPER */
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
  background: white;
}

.header h3 {
  font-weight: 700;
  font-size: 18px;
  color: black; /* 🔥 FIX WARNA */
}

.back {
  font-size: 20px;
  cursor: pointer;
}

/* FILTER */
.filter {
  display: flex;
  gap: 10px;
  padding: 14px 16px;
}

.tab {
  flex: 1;
  padding: 10px;
  border-radius: 12px;
  border: 1.5px solid #ff3b30;
  background: white;
  color: #ff3b30;
  font-weight: 600;
}

.tab.active {
  background: #ff3b30;
  color: white;
}

/* CONTENT */
.content {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

/* CARD */
.card {
  background: #ff3b30;
  color: white;
  padding: 16px;
  border-radius: 18px;
  box-shadow: 0 6px 14px rgba(0,0,0,0.2);
}

/* ROW */
.row {
  display: flex;
  align-items: flex-start;
  gap: 12px;
}

/* ICON */
.icon {
  font-size: 30px;
}

/* TEXT (RATA KIRI SEMUA 🔥) */
.text {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

/* MATKUL */
.matkul {
  font-size: 16px;
  font-weight: 700;
}

/* RUANG */
.ruang {
  font-size: 14px;
  margin-top: 2px;
}

/* JAM */
.jam {
  font-size: 13px;
  margin-top: 2px;
}

</style>