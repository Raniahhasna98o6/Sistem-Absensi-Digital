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

	// 1. INI RUTE REGISTER YANG HILANG (Biar Frontend Nggak 404)
	r.POST("/register", func(c *gin.Context) {
		// Logika insert database untuk registrasi ditaruh di sini nanti
		c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil masuk backend!"})
	})

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
			c.SetCookie("nim_user", m.NIM, 3600, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{"message": "Login Mahasiswa Berhasil", "nama": m.Nama})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau Password salah"})
		}
	})

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

	r.GET("/api/absensi/riwayat", func(c *gin.Context) {
		nim, _ := c.Cookie("nim_user")
		m := models.Mahasiswa{NIM: nim}
		c.JSON(http.StatusOK, m.BukaRiwayat())
	})

	r.POST("/login/dosen", func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		c.BindJSON(&req)

		d := models.Dosen{}
		if d.Login(req.Email, req.Password) {
			c.SetCookie("nidn_user", d.NIDN, 3600, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{"message": "Login Dosen Berhasil", "nama": d.Nama})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Akses Dosen ditolak"})
		}
	})

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

	r.POST("/logout", func(c *gin.Context) {
		c.SetCookie("nim_user", "", -1, "/", "localhost", false, true)
		c.SetCookie("nidn_user", "", -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "Logout berhasil"})
	})

	// 2. PERBAIKAN PORT UNTUK AZURE
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run("0.0.0.0:" + port)
}
