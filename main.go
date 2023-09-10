package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/api", func(c *gin.Context) {

		slackName := c.DefaultQuery("slack_name", "")
		track := c.DefaultQuery("track", "")

		if slackName == "" || track == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request. Please provide valid parameters.",
			})
			return
		}

		currentDay := time.Now().UTC().Weekday().String()

		currentTime := time.Now().UTC()
		fmt.Println("Current UTC Time:", currentTime)

		if !isValidTime(currentTime) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid UTC time. It should be within +/-2 hours of the current time.",
			})
			return
		}

		fileURL := "https://github.com/haywhyz/task1/blob/main/main.go"
		sourceCodeURL := "https://github.com/haywhyz/task1"

		response := gin.H{
			"slack_name":      slackName,
			"current_day":     currentDay,
			"utc_time":        currentTime.Format(time.RFC3339),
			"track":           track,
			"github_file_url": fileURL,
			"github_repo_url": sourceCodeURL,
		}

		c.JSON(http.StatusOK, response)
	})

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func isValidTime(t time.Time) bool {
	currentTime := time.Now().UTC()
	minValidTime := currentTime.Add(-2 * time.Hour)
	maxValidTime := currentTime.Add(2 * time.Hour)

	return t.After(minValidTime) && t.Before(maxValidTime)
}
