// server\routes\UpdateUser.go

package routes

import (
	"database/sql"
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UpdateUser(c *gin.Context, db *sql.DB) {
	email := c.Query("email")
	password := c.Query("password")
	newFirstname := c.Query("new_firstname")
	newPassword := c.Query("new_password")

	var user models.User
	err := db.QueryRow("SELECT id, firstname, password FROM users WHERE email ILIKE $1", email).Scan(&user.ID, &user.FirstName, &user.Password)
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec("UPDATE users SET password = $1, firstname = $2 WHERE email = $3", hashedPassword, newFirstname, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User information updated successfully"})
}
