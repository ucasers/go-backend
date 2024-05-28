package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
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
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "error": "Unable to get request"})
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "error": "Cannot unmarshal body"})
		return
	}

	if errors := user.Validate("login"); len(errors) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "error": errors})
		return
	}

	userData, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "response": userData})
}

// SignIn 处理用户身份验证
func (server *Server) SignIn(email, password string) (map[string]interface{}, error) {
	// 使用生成的 User 模型
	user, err := query.User.
		WithContext(server.DB.Statement.Context).
		Where(query.User.Email.Eq(email)).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("invalid email or password")
		}
		server.logError("Error executing the query", err)
		return nil, fmt.Errorf("internal server error")
	}

	if user.Password != password {
		server.logError("Error verifying the password", err)
		return nil, fmt.Errorf("invalid email or password")
	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		server.logError("Error creating the token", err)
		return nil, fmt.Errorf("internal server error")
	}

	userData := map[string]interface{}{
		"token":    token,
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
	}
	return userData, nil
}

func (server *Server) Register(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "error": "Unable to get request"})
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "error": "Cannot unmarshal body"})
		return
	}

	err = query.User.
		WithContext(server.DB.Statement.Context).Create(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "error": "Email already exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "response": "successfully register"})

}

// logError 记录错误日志
func (server *Server) logError(message string, err error) {
	fmt.Printf("%s: %v\n", message, err)
}
