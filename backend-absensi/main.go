package main

import (
	"backend-absensi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	// login page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// dashboard page
	r.GET("/dashboard", func(c *gin.Context) {
		login, _ := c.Cookie("login")
		if login != "true" {
			c.Redirect(http.StatusFound, "/")
			return
		}
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// absensi page
	r.GET("/absensi", func(c *gin.Context) {
		login, _ := c.Cookie("login")
		if login != "true" {
			c.Redirect(http.StatusFound, "/")
			return
		}
		c.HTML(http.StatusOK, "absensi.html", nil)
	})

	// login API
	r.POST("/login", func(c *gin.Context) {
		type LoginRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		var req LoginRequest
		c.BindJSON(&req)

		user := models.NewUser("1", "Hasna", "admin@gmail.com", "admin123")

		if user.Login(req.Email, req.Password) {
			c.SetCookie("login", "true", 3600, "/", "", false, true)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Login Berhasil",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Email atau password salah",
			})
		}
	})

	// logout
	r.GET("/logout", func(c *gin.Context) {
		c.SetCookie("login", "", -1, "/", "", false, true)
		c.Redirect(http.StatusFound, "/")
	})

	// API absensi
	r.GET("/api/absensi", func(c *gin.Context) {
		login, _ := c.Cookie("login")
		if login != "true" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "Hadir",
		})
	})

	// API jadwal
	r.GET("/get/jadwal", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"matkul": "Jaringan Komputer",
			"jam":    "08.30 - 10.30",
			"ruang":  "TULT 0714",
		})
	})

	r.Run(":8080")
}
