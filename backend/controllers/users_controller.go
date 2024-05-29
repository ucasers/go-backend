package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ucasers/go-backend/backend/utils"
	"net/http"
)

func (server *Server) GetUser(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   utils.ResponseData(user, "User"),
	})

}
