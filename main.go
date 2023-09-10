package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a handler function for the endpoint
	router.GET("/info", func(c *gin.Context) {
		// Get query parameters
		slackName := c.DefaultQuery("slack_name", "")
		track := c.DefaultQuery("track", "")

		// Debugging: Print the received parameters
		fmt.Println("Received Slack Name:", slackName)
		fmt.Println("Received Track:", track)

		// Check if slackName and track are not empty
		if slackName == "" || track == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request. Please provide valid parameters.",
			})
			return
		}

		// Get the current day of the week
		currentDay := time.Now().UTC().Weekday().String()

		// Get the current UTC time with validation of +/-2 hours
		currentTime := time.Now().UTC()
		// Debugging: Print the current time
		fmt.Println("Current UTC Time:", currentTime)

		// Check if the current time is within +/-2 hours
		if !isValidTime(currentTime) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid UTC time. It should be within +/-2 hours of the current time.",
			})
			return
		}

		// Construct the GitHub URLs (replace with your actual URLs)
		fileURL := "https://github.com/haywhyz/task1/blob/main/main.go"
		sourceCodeURL := "https://github.com/haywhyz/task1"

		// Prepare the response JSON
		response := gin.H{
			"slack_name":      slackName,
			"current_day":     currentDay,
			"utc_time":        currentTime.Format(time.RFC3339),
			"track":           track,
			"github_file_url": fileURL,
			"github_repo_url": sourceCodeURL,
		}

		// Set the appropriate status code
		c.JSON(http.StatusOK, response)
	})

	// Run the server on port 8080
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// Function to validate UTC time within +/-2 hours
func isValidTime(t time.Time) bool {
	currentTime := time.Now().UTC()
	minValidTime := currentTime.Add(-2 * time.Hour)
	maxValidTime := currentTime.Add(2 * time.Hour)

	return t.After(minValidTime) && t.Before(maxValidTime)
}
