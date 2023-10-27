// server\routes\DeleteUser.go
package routes

import (
	"database/sql"
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func DeleteUser(c *gin.Context, db *sql.DB) {
	email := c.Query("email")
	password := c.Query("password")

	var user models.User
	err := db.QueryRow("SELECT id, firstname, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.FirstName, &user.Password)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
		return
	}

	_, err = db.Exec("DELETE FROM users WHERE id = $1", user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}
