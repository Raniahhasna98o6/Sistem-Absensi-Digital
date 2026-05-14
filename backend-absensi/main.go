package main

import (
	"backend-absensi/config"
	"backend-absensi/models"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	// 1. CORS
	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		allowedOrigins := map[string]bool{
			"http://localhost:5173":            true,
			"https://absensisistem.vercel.app": true,
		}

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

	// 2. REGISTER (placeholder)
	r.POST("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil masuk backend!"})
	})

	// 3. LOGIN MAHASISWA
	r.POST("/login/mahasiswa", func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah!"})
			return
		}

		m := models.Mahasiswa{}
		if m.Login(req.Email, req.Password) {
			c.SetCookie("nim_user", m.NIM, 3600, "/", "", false, true)
			c.JSON(http.StatusOK, gin.H{
				"message": "Login Mahasiswa Berhasil",
				"nama":    m.Nama,
				"nim":     m.NIM, // FIX: kirim NIM ke frontend untuk disimpan di localStorage
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau Password salah"})
		}
	})

	// 4. INPUT ABSENSI — baca NIM dari body (bukan cookie, karena cross-domain)
	r.POST("/api/absensi", func(c *gin.Context) {
		var data models.Absensi
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data absensi tidak valid"})
			return
		}

		if data.Nim == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "NIM tidak ditemukan"})
			return
		}

		m := models.Mahasiswa{NIM: data.Nim}
		success, pesan := m.InputKehadiran(data)

		if success {
			c.JSON(http.StatusOK, gin.H{"message": pesan})
		} else {
			c.JSON(http.StatusForbidden, gin.H{"message": pesan})
		}
	})

	// 5. RIWAYAT ABSENSI — FIX: baca NIM dari query param, bukan cookie
	r.GET("/api/absensi/riwayat", func(c *gin.Context) {
		// Coba dari query param dulu (untuk frontend cross-domain)
		nim := c.Query("nim")
		if nim == "" {
			// Fallback ke cookie (untuk testing Postman)
			nim, _ = c.Cookie("nim_user")
		}
		if nim == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "NIM tidak ditemukan"})
			return
		}

		m := models.Mahasiswa{NIM: nim}
		c.JSON(http.StatusOK, m.BukaRiwayat())
	})

	// 6. LOGIN DOSEN
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
			c.JSON(http.StatusOK, gin.H{
				"message": "Login Dosen Berhasil",
				"nama":    d.Nama,
				"nidn":    d.NIDN, // FIX: kirim NIDN ke frontend
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Akses Dosen ditolak"})
		}
	})

	// 6.5. AMBIL PROFIL DOSEN
	r.GET("/api/dosen/profil", func(c *gin.Context) {
		// Ambil NIDN dari query param (biar aman buat frontend Vercel) atau fallback ke cookie
		nidn := c.Query("nidn")
		if nidn == "" {
			nidn, _ = c.Cookie("nidn_user")
		}

		if nidn == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "NIDN tidak ditemukan, sesi habis"})
			return
		}

		d := models.Dosen{}
		profil, err := d.AmbilProfil(nidn)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profil tidak ditemukan"})
			return
		}
		c.JSON(http.StatusOK, profil)
	})

	// 7. LAPORAN DOSEN
	r.GET("/api/dosen/laporan", func(c *gin.Context) {
		// Coba cookie dulu, fallback ke query param
		nidn, err := c.Cookie("nidn_user")
		if err != nil || nidn == "" {
			nidn = c.Query("nidn")
		}
		if nidn == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sesi telah berakhir"})
			return
		}

		periode := c.DefaultQuery("periode", "2026-05")
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

	// 8. LOGOUT
	r.POST("/logout", func(c *gin.Context) {
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie("nim_user", "", -1, "/", "", true, true)
		c.SetCookie("nidn_user", "", -1, "/", "", true, true)
		c.JSON(http.StatusOK, gin.H{"message": "Logout berhasil"})
	})

	// AMBIL PROFIL MAHASISWA
	r.GET("/api/mahasiswa/profil", func(c *gin.Context) {
		nim := c.Query("nim")
		if nim == "" {
			nim, _ = c.Cookie("nim_user")
		}

		if nim == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "NIM tidak ditemukan, sesi habis"})
			return
		}

		m := models.Mahasiswa{}
		profil, err := m.GetAttribute(nim)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profil tidak ditemukan"})
			return
		}
		c.JSON(http.StatusOK, profil)
	})

	// AMBIL DAFTAR KELAS DOSEN
	r.GET("/api/dosen/daftar-kelas", func(c *gin.Context) {
		nidn := c.Query("nidn")
		if nidn == "" {
			nidn, _ = c.Cookie("nidn_user")
		}

		d := models.Dosen{NIDN: nidn}
		list, err := d.AmbilDaftarKelas()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil daftar kelas"})
			return
		}
		c.JSON(http.StatusOK, list)
	})

	// 9. PORT AZURE
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
