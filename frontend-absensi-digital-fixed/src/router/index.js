import { createRouter, createWebHistory } from 'vue-router'

// PAGES
import Login from '../pages/Login.vue'
import Beranda from '../pages/Beranda.vue'
import BerandaDosen from '../pages/BerandaDosen.vue'
import Absensi from '../pages/Absensi.vue'
import Kamera from '../pages/Kamera.vue'
import Preview from '../pages/Preview.vue'
import Success from '../pages/Success.vue'
import Akun from '../pages/Akun.vue'
import Riwayat from '../pages/Riwayat.vue'
import GenerateLaporan from '../pages/GenerateLaporan.vue'
import Role from '../pages/Role.vue'
import Register from '../pages/Register.vue'
const routes = [
  {
    path: '/',
    redirect: '/role'
  },
  {
    path: '/role',
    name: 'Role',
    component: Role
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/beranda',
    name: 'Beranda',
    component: Beranda
  },
  {
  path: '/beranda-dosen',
  component: BerandaDosen
  },
  {
    path: '/absensi',
    name: 'Absensi',
    component: Absensi
  },
  {
    path: '/kamera',
    name: 'Kamera',
    component: Kamera
  },
  {
    path: '/preview',
    name: 'Preview',
    component: Preview
  },
  {
    path: '/success',
    name: 'Success',
    component: Success
  },
  {
    path: '/akun',
    name: 'Akun',
    component: Akun
  },
  {
    path: '/riwayat',
    name: 'Riwayat',
    component: Riwayat
  },
  {
  path: '/laporan',
  component: GenerateLaporan
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router