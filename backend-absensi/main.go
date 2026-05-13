package main

import (
	"backend-absensi/config"
	"backend-absensi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	// --- 1. JEMBATAN CORS (VERSI DINAMIS UNTUK VERCEL & LOCALHOST) ---
	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Daftar URL yang lu kasih izin buat nembak ke Azure
		allowedOrigins := map[string]bool{
			"http://localhost:5173":            true, // Untuk ngetes di laptop lu
			"https://absensisistem.vercel.app": true, // Link Vercel (PASTIKAN TANPA GARIS MIRING DI AKHIR)
		}

		// Kalau Origin dari request ada di daftar atas, kasih izin
		if allowedOrigins[origin] {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 2. RUTE REGISTER
	r.POST("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil masuk backend!"})
	})

	// 3. RUTE LOGIN MAHASISWA
	r.POST("/login/mahasiswa", func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah"})
			return
		}

		m := models.Mahasiswa{}
		if m.Login(req.Email, req.Password) {
			c.SetSameSite(http.SameSiteNoneMode)
			c.SetCookie("nim_user", m.NIM, 3600, "/", "", true, true)
			c.JSON(http.StatusOK, gin.H{"message": "Login Mahasiswa Berhasil", "nama": m.Nama})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau Password salah"})
		}
	})

	// 4. RUTE INPUT ABSENSI
	r.POST("/api/absensi", func(c *gin.Context) {
		var data models.Absensi
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data absensi tidak valid"})
			return
		}

		nim, err := c.Cookie("nim_user")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Silakan login terlebih dahulu"})
			return
		}

		m := models.Mahasiswa{NIM: nim}
		success, pesan := m.InputKehadiran(data)

		if success {
			c.JSON(http.StatusOK, gin.H{"message": pesan})
		} else {
			c.JSON(http.StatusForbidden, gin.H{"message": pesan})
		}
	})

	// 5. RUTE RIWAYAT ABSENSI
	r.GET("/api/absensi/riwayat", func(c *gin.Context) {
		nim, _ := c.Cookie("nim_user")
		m := models.Mahasiswa{NIM: nim}
		c.JSON(http.StatusOK, m.BukaRiwayat())
	})

	// 6. RUTE LOGIN DOSEN
	r.POST("/login/dosen", func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		c.BindJSON(&req)

		d := models.Dosen{}
		if d.Login(req.Email, req.Password) {
			c.SetSameSite(http.SameSiteNoneMode)
			c.SetCookie("nidn_user", d.NIDN, 3600, "/", "", true, true)
			c.JSON(http.StatusOK, gin.H{"message": "Login Dosen Berhasil", "nama": d.Nama})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Akses Dosen ditolak"})
		}
	})

	// 7. RUTE LAPORAN DOSEN
	r.GET("/api/dosen/laporan", func(c *gin.Context) {
		nidn, err := c.Cookie("nidn_user")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sesi telah berakhir"})
			return
		}

		periode := c.DefaultQuery("periode", "2026-04")
		idKelasStr := c.Query("id_kelas")
		idKelas, _ := strconv.Atoi(idKelasStr)

		d := models.Dosen{NIDN: nidn}
		laporan, err := d.MintaLaporan(periode, idKelas)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, laporan)
	})

	// 8. RUTE LOGOUT
	r.POST("/logout", func(c *gin.Context) {
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie("nim_user", "", -1, "/", "", true, true)
		c.SetCookie("nidn_user", "", -1, "/", "", true, true)
		c.JSON(http.StatusOK, gin.H{"message": "Logout berhasil"})
	})

	// 9. PERBAIKAN PORT UNTUK AZURE
	r.Run("0.0.0.0:8080")
}
