package controllers

import (
	"net/http"

	"github.com/carbans/netdebug/models"
	"github.com/carbans/netdebug/types"
	"github.com/carbans/netdebug/utils/token"
	"github.com/gin-gonic/gin"
)

func CreateApiKey(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input types.CreateApiKeyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apikey, err := models.CreateApiKey(input, user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"apikey": apikey})

}
