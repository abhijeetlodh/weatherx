// server\main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/jwt"
	"server/models"
	"server/routes"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	gin.SetMode(gin.ReleaseMode)
	errENV := godotenv.Load()
	if errENV != nil {
		log.Fatal("Error loading .env file: ", errENV)
	}

	var err error
	db, err := sql.Open("postgres", os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	c := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	router.Use(c)

	router.POST("/register", func(c *gin.Context) {
		routes.RegisterUser(c, db)
	})

	router.GET("/login", func(c *gin.Context) {
		routes.LoginUser(c, db)
		userIDStr := c.Query("user_id")
		userID, err := strconv.Atoi(userIDStr)
		authToken, err := jwt.GenerateJWTToken(userID, 30*time.Minute)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
			return
		}
		jwt.SetCookieWithJWTToken(c.Writer, authToken, 30*time.Minute)
	})

	router.PUT("/update", func(c *gin.Context) {
		routes.UpdateUser(c, db)
	})

	router.DELETE("/delete", func(c *gin.Context) {
		routes.DeleteUser(c, db)
	})

	router.GET("/weather", routes.WeatherHandler)

	router.GET("/fetchhistorytable", func(c *gin.Context) {
		models.FetchHistoryTableHandler(db, c)
	})

	router.POST("/logout", func(c *gin.Context) {
		tokenCookie := jwt.RemoveJWTTokenFromCookie()
		http.SetCookie(c.Writer, tokenCookie)
		c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
	})

	router.POST("/saveSearchData", func(c *gin.Context) {
		var data models.Weather
		if err := c.BindJSON(&data); err != nil {
			log.Printf("Data to insert: %+v", data)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data.CreatedAt = time.Now()
		fmt.Println(data.CreatedAt)
		err := models.SaveSearchData(db, data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Search data saved successfully"})
	})

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
