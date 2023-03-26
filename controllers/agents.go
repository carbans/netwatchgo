package controllers

import (
	"net/http"

	"github.com/carbans/netdebug/models"
	"github.com/carbans/netdebug/utils/token"
	"github.com/gin-gonic/gin"
)

type CreateAgentInput struct {
	Name    string `json:"name" binding:"required"`
	Version string `json:"version"`
}

type UpdateAgentInput struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// GET /agents
// Find all agents
func FindAgents(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": models.GetAgents(user_id)})
}

// GET /agents/:id
// Find a agent
func FindAgent(c *gin.Context) {
	// Get model if exist
	user_id, err := token.ExtractTokenID(c)
	println("Get Agent: ", user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	println("Test")

	var agent models.Agent

	if err := models.DB.Where("id = ? AND user_id = ?", c.Param("id"), user_id).First(&agent).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": agent})
}

// POST /agents
// Create new agent
func CreateAgent(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	var input CreateAgentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create agent
	agent := models.Agent{Name: input.Name, Version: input.Version, UserID: user_id}
	models.DB.Create(&agent)

	c.JSON(http.StatusOK, gin.H{"data": agent})
}

// PATCH /agents/:id
// Update a agent
func UpdateAgent(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Get model if exist
	var agent models.Agent
	if err := models.DB.Where("id = ? AND user_id = ?", c.Param("id"), user_id).First(&agent).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateAgentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&agent).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": agent})
}

// DELETE /agent/:id
// Delete a agent
func DeleteAgent(c *gin.Context) {
	// Get model if exist
	var agent models.Agent
	if err := models.DB.Where("id = ?", c.Param("id")).First(&agent).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&agent)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
