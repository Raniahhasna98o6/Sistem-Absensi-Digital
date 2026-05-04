<template>
    <div class="wrapper">
      <div class="phone">
        <div class="header">
          <h2>Choose Role</h2>
          <p>Please select your role to continue</p>
          <p style="font-size: smaller;">Use the format below on your email</p>
        </div>

        <!-- LOGO -->
        <div class="logo-wrapper">
          <img :src="logo" class="logo" />
        </div>
        
        <div class="content">

          <!-- Lecturer -->
          <div class="card" :class="{ 'selected': role === 'lecturer' }" @click="selectRole('lecturer')">
            <h2 style="color: cadetblue;">Lecturer</h2>
            <p>@lecturer.university.ac.id</p>
          </div>

          <!-- Student -->
          <div class="card" :class="{ 'selected': role === 'student' }" @click="selectRole('student')">
            <h2 style="color: cadetblue;">Student</h2>
            <p>@student.university.ac.id</p>
          </div>

          <button :disabled="!role" @click="goToLogin">
            Continue
          </button>
        </div>
      </div>
    </div>
</template>

<script>
import logo from '../assets/logo.png'
export default {
  data() {
    return {
      role: null,
      logo
    }
  },
  methods: {
    selectRole(role) {
      this.role = role
    },
    goToLogin() {
      this.$router.push({ 
        path: '/login',
        query: { role: this.role }
    });
    }
  }
}
</script>

<style scoped>
/* LOGO */
.logo-wrapper {
  display: flex;
  justify-content: center;
  margin-top: -70px;
}

.logo {
  width: 95px;
}

.header p {
  font-style: italic;
}

.wrapper {
  background: #0f172a;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.phone {
  width: 390px;
  min-height: 780px;
  background: white;
  border-radius: 30px;
  overflow: hidden;
}

.header {
  background-color: #ff2d2d;
  color: white;
  text-align: center;
  padding: 50px 20px 100px;
  border-bottom-left-radius: 50% 25%;
  border-bottom-right-radius: 50% 25%;
}

.header h2 {
  font-size: 30px;
}

.content {
  padding: 25px;
  margin-top: 10px;
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.card {
  padding: 18px;
  border-radius: 15px;
  border: 1px solid #ddd;
  background: #f5f5f5;
  text-align: center;
  cursor: pointer;
  transition: 0.2s;
  box-shadow: 0 2px 8px rgba(0,0,0,0.4);
}

.card h2 {
  margin: 0;
  font-size: 18px;
  color: #2f80ed;
}

.card p {
  margin-top: 5px;
  font-size: 14px;
  color: #555;
}

.card:hover {
  background: #fff;
  border: 1px solid #ff2d2d;
}

.card.selected {
  background: lightcyan;
  border: 1px solid #ff2d2d;
  box-shadow: 0 4px 12px rgba(255, 45, 45, 0.2);
}

button {
  margin-top: 15px;
  padding: 14px;
  border: none;
  border-radius: 12px;
  background: #ff2d2d;
  color: white;
  font-weight: bold;
  font-size: 16px;
  cursor: pointer;
}

button:disabled {
  background: gray;
  cursor: not-allowed;
}
</style>