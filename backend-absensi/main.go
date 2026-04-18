package main
import (
	"backend-absensi/config"
	"backend-absensi/models"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// endpoint test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// login API
	r.POST("/login", func(c *gin.Context) {
		type LoginRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var req LoginRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		user, err := models.GetUserByEmail(req.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Email atau password salah",
			})
			return
		}

		if user.CheckPassword(req.Password) {
			c.SetCookie("login", "true", 3600, "/", "localhost", false, true)
			c.SetCookie("id_user", user.ID, 3600, "/", "localhost", false, true)

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
		c.SetCookie("id_user", "", -1, "/", "", false, true)
		c.Redirect(http.StatusFound, "/")
	})

	r.GET("/api/absensi", func(c *gin.Context) {
		if !isLogin(c) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		idUser, _ := c.Cookie("id_user")
		data, err := models.GetAbsensiByUser(idUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, data)
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

func isLogin(c *gin.Context) bool {
	login, err := c.Cookie("login")
	if err != nil || login != "true" {
		return false
	}
	return true
}
