// server\routes\RegisterUser.go
package routes

import (
	"database/sql"
	"net/http"
	"time"

	"server/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context, db *sql.DB) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var existingUser models.User
	err := db.QueryRow("SELECT id FROM users WHERE email = $1", user.Email).Scan(&existingUser.ID)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
		return
	} else if err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var lastID int
	err = db.QueryRow("SELECT COALESCE(MAX(id), 0) FROM users").Scan(&lastID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.ID = lastID + 1
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = string(hashedPassword)
	currentTime := time.Now().In(time.FixedZone("IST", 5*60*60))
	createdTimeStr := currentTime.Format("02-Jan-06 03:04 PM")
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime
	query := `
        INSERT INTO users (id, email, firstname, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = db.Exec(query, user.ID, user.Email, user.FirstName, user.Password, createdTimeStr, user.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
