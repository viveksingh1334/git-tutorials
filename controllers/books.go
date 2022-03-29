package controllers

import (
	
	"github.com/gin-gonic/gin"
	"github.com/viveksingh1334/viveksingh1334/models"
)

type CreateUserInput struct {
    Name    string  `json:"name" binding:"required"`
	Recent_apps  string  `json:"recent_apps" binding:"required"`
	Favourite_apps  string  `json:"favourite_apps" binding:"required"`
}

type UpdateUserInput struct {
    Name    string  `json:"name"`
	Recent_apps  string  `json:"recent_apps"`
	Favourite_apps  string `json:"favourite_apps"`
}

// GET /users
// Find all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /users/:id
// Find a user
func Finduser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /users
// Create new user
func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create User
	user := models.User{Name: input.Name, Recent_apps: input.Recent_apps,Favourite_apps: input.Favourite_apps}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /users/:id
// Update a user
func UpdateUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id
// Delete a User
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
