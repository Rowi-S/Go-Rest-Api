package handlers

import (
	"music/web-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum models.Album

	// Bind the JSON request body to the newAlbum object
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if all required fields are present (non-empty)
	if newAlbum.Title == "" || newAlbum.Artist == "" || newAlbum.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// Generate a new ID for the album
	if len(albums) > 0 {
		lastAlbum := albums[len(albums)-1]

		// Convert the last ID to an integer, increment it, and convert it back to string
		lastID, err := strconv.Atoi(lastAlbum.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID format"})
			return
		}

		newAlbum.ID = strconv.Itoa(lastID + 1)
	} else {
		// If no albums exist yet, start from ID 1
		newAlbum.ID = "1"
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)

	// Return the newly created album
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func PatchAlbum(c *gin.Context) {
	id := c.Param("id")

	// Find the album by ID
	for i, a := range albums {
		if a.ID == id {
			var updatedFields map[string]interface{}
			if err := c.BindJSON(&updatedFields); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
				return
			}

			// Update only the fields that are present in the request
			if title, exists := updatedFields["Title"].(string); exists {
				albums[i].Title = title
			}
			if artist, exists := updatedFields["Artist"].(string); exists {
				albums[i].Artist = artist
			}
			if price, exists := updatedFields["Price"].(float64); exists {
				albums[i].Price = price
			}

			c.JSON(http.StatusOK, albums[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	// Find the album and remove it
	for i, a := range albums {
		if a.ID == id {
			// Remove the album from the slice by appending the items before and after it
			albums = append(albums[:i], albums[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Album deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}
