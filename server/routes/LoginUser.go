// server\routes\LoginUser.go
package routes

import (
	"database/sql"
	"net/http"

	"server/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context, db *sql.DB) {
	email := c.Query("email")
	password := c.Query("password")

	var user models.User
	err := db.QueryRow("SELECT id, firstname, password, email FROM users WHERE email = $1", email).Scan(&user.ID, &user.FirstName, &user.Password, &user.Email)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wrong email"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Updated At": user.UpdatedAt, "Created At": user.CreatedAt, "firstname": user.FirstName, "email": user.Email, "id": user.ID})
}
