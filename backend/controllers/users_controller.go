package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) GetUser(c *gin.Context) {

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve user from context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"data": ,
	})
}
