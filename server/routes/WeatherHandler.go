// server/routes/WeatherHandler.go
package routes

import (
	"net/http"
	"os"

	"github.com/briandowns/openweathermap"
	"github.com/gin-gonic/gin"
)

type WeatherResponse struct {
	CityName    string  `json:"CityName"`
	Temperature float64 `json:"Temperature"`
}

func WeatherHandler(c *gin.Context) {
	cityName := c.Query("cityName")
	metric := c.Query("metric")

	if cityName == "" || metric == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing cityName or metric parameter"})
		return
	}

	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing OpenWeatherMap API key"})
		return
	}

	wm, err := openweathermap.NewCurrent(metric, "en", apiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = wm.CurrentByName(cityName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := WeatherResponse{
		CityName:    wm.Name,
		Temperature: wm.Main.Temp,
	}

	c.JSON(http.StatusOK, resp)
}
