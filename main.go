package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type baseballPlayer struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	Team               string  `json:"team"`
	BattingAverage     float32 `json:"batting_average"`
	OnBasePercentage   float32 `json:"on_base_percentage"`
	OnBasePlusSlugging float32 `json:"on_base_plus_slugging"`
}

var baseballPlayers = []baseballPlayer{
	{ID: "1", Name: "Bryce Harper", Team: "Philadelphia Phillies", BattingAverage: 0.286, OnBasePercentage: 0.374, OnBasePlusSlugging: 0.895},
	{ID: "2", Name: "Shohei Ohtani", Team: "Los Angeles Dodgers", BattingAverage: 0.292, OnBasePercentage: 0.376, OnBasePlusSlugging: 0.993},
	{ID: "3", Name: "Trea Turner", Team: "Philadelphia Phillies", BattingAverage: 0.299, OnBasePercentage: 0.345, OnBasePlusSlugging: 0.819},
}

func main() {
	router := gin.Default()
	router.GET("/baseballPlayers", getBaseballPlayers)
	router.GET("/baseballPlayers/:id", getBaseballPlayerByID)
	router.POST("/baseballPlayers", postBaseballPlayers)

	router.Run("localhost:8080")
}

func getBaseballPlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, baseballPlayers)
}

func postBaseballPlayers(c *gin.Context) {
	var newBaseballPlayer baseballPlayer

	if err := c.BindJSON(&newBaseballPlayer); err != nil {
		return
	}

	baseballPlayers = append(baseballPlayers, newBaseballPlayer)
	c.IndentedJSON(http.StatusCreated, newBaseballPlayer)
}

func getBaseballPlayerByID(c *gin.Context) {
	id := c.Param("id")

	for _, player := range baseballPlayers {
		if player.ID == id {
			c.IndentedJSON(http.StatusOK, player)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "player not found"})
}
