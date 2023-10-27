// server/models/weather.go

package models

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Weather struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	TempF     string    `json:"temp_f"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	Metric    string    `json:"metric"`
}

func SaveSearchData(db *sql.DB, data Weather) error {
	sqlStatement := `
        INSERT INTO weather (user_id, location, temp_f, created_at, metric)
        VALUES ($1, $2, $3::text, $4, $5)
    `
	_, err := db.Exec(sqlStatement, data.UserID, data.Location, data.TempF, data.CreatedAt, data.Metric)

	if err != nil {
		log.Printf("Error saving search data: %v", err)
		return err
	}

	log.Printf("Search data saved successfully")
	return nil
}

func FetchHistoryTableHandler(db *sql.DB, c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	weatherData, err := RetrieveWeatherData(db, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve weather data"})
		return
	}

	var formattedData []map[string]interface{}
	for _, data := range weatherData {
		formattedData = append(formattedData, map[string]interface{}{
			"userID":      data.UserID,
			"CityName":    data.Location,
			"Temperature": data.TempF,
			"Metric":      data.Metric,
			"Time(ist)":   data.CreatedAt.Format("15:04 PM 02/01/2006"),
		})
	}

	c.JSON(http.StatusOK, formattedData)
}

func RetrieveWeatherData(db *sql.DB, userID int) ([]Weather, error) {
	rows, err := db.Query("SELECT * FROM weather WHERE user_id = $1 ORDER BY Created_At DESC LIMIT 5", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var weatherData []Weather
	for rows.Next() {
		var data Weather
		err := rows.Scan(&data.ID, &data.UserID, &data.TempF, &data.Location, &data.CreatedAt, &data.Metric)
		if err != nil {
			return nil, err
		}
		weatherData = append(weatherData, data)
	}

	return weatherData, nil
}
