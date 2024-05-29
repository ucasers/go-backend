package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ucasers/go-backend/backend/auth"
	"github.com/ucasers/go-backend/backend/models"
	"github.com/ucasers/go-backend/query"
	"gorm.io/gorm"
	"io"
	"net/http"
)

// Login 处理用户登录请求
func (server *Server) Login(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "请求参数错误"})
		return
	}

	var responseUser models.User
	if err := json.Unmarshal(body, &responseUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "请求参数错误"})
		return
	}

	user, err := query.Q.User.
		WithContext(c).
		Where(query.User.Email.Eq(responseUser.Email)).
		First()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "邮箱或用户名不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "服务器内部错误"})
		return
	}

	if user.Password != responseUser.Password {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "账号或密码错误"})
		return
	}

	token, _ := auth.CreateToken(user.ID)
	userData := map[string]interface{}{
		"token":    token,
		"email":    user.Email,
		"username": user.Username,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": userData})
	return
}

func (server *Server) Register(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "请求参数错误"})
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "请求参数错误"})
		return
	}

	err = query.Q.User.
		WithContext(c).Create(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "邮箱已经存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})

}

func (server *Server) HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   "hello world",
	})
}
