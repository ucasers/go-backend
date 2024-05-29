package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ucasers/go-backend/backend/utils"
	"net/http"
)

func (server *Server) GetUser(c *gin.Context) {

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "token验证失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   utils.ResponseData(user, "User"),
	})
}
