package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    "hello world",
	})
}

func (server *Server) GetUser(c *gin.Context) {

	//clear previous error if any
	//errList := map[string]string{}
	//
	//userID := c.Param("id")
	//
	//uid, err := strconv.ParseUint(userID, 10, 32)
	//if err != nil {
	//	errList["Invalid_request"] = "Invalid Request"
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"status": http.StatusBadRequest,
	//		"error":  errList,
	//	})
	//	return
	//}
	//user := models.User{}
	//
	//userGotten, err := user.FindUserByID(server.DB, uint32(uid))
	//if err != nil {
	//	errList["No_user"] = "No User Found"
	//	c.JSON(http.StatusNotFound, gin.H{
	//		"status": http.StatusNotFound,
	//		"error":  errList,
	//	})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": "hello world",
	})
}
