package main

import (
	"backend-absensi/config"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	// 1. TAMBAHIN RUTE REGISTER BIAR NGGAK 404
	r.POST("/register", func(c *gin.Context) {
		// Nanti logika insert ke database lu taruh di sini
		// Ini contoh balasan sukses sementara biar Frontend lu bisa jalan dulu
		c.JSON(http.StatusOK, gin.H{"message": "Registrasi Berhasil Diterima Backend!"})
	})

	// Rute Login (Ingat, di Vue lu harus nembak ke /login/mahasiswa atau dosen, bukan cuma /login)
	r.POST("/login/mahasiswa", func(c *gin.Context) {
		// ... (kode lu tetap sama) ...
	})

	// ... (rute absensi dan dosen tetap sama) ...

	// Hapus duplicate /logout yang ada di kode lama lu
	r.POST("/logout", func(c *gin.Context) {
		c.SetCookie("nim_user", "", -1, "/", "localhost", false, true)
		c.SetCookie("nidn_user", "", -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "Logout berhasil"})
	})

	// 2. PERBAIKAN CARA RUN SERVER BUAT AZURE
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback buat ngetes di laptop
	}

	// Cukup SATU kali r.Run() di paling akhir file
	r.Run(":" + port)
}
