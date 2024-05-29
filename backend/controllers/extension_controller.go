package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ucasers/go-backend/backend/models"
	"github.com/ucasers/go-backend/dao"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
)

func (server *Server) UploadExtension(c *gin.Context) {
	user, _ := c.Get("user")
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "请求参数错误"})
		return
	}
	var extension models.Extension
	if err := json.Unmarshal(body, &extension); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "请求参数错误"})
		return
	}
	// 确保类型转换成功
	userModel, _ := user.(*models.User)

	// 绑定 user 到 extension
	extension.OwnerID = userModel.ID

	err = dao.Q.Extension.
		WithContext(c.Request.Context()).Create(&extension)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "上传失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (server *Server) ModifyExtension(c *gin.Context) {
	// 获取当前用户
	user, _ := c.Get("user")
	userModel, _ := user.(*models.User)

	// 读取请求体
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "请求参数错误"})
		return
	}

	// 解析请求体中的新数据
	var newExtensionData models.Extension
	if err := json.Unmarshal(body, &newExtensionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "请求参数错误"})
		return
	}

	extension, err := dao.Q.Extension.
		WithContext(c.Request.Context()).
		Where(dao.Extension.ID.Eq(newExtensionData.ID)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "找不到扩展"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "查询扩展时出错"})
		}
		return
	}

	// 验证当前用户是否为扩展的拥有者
	if extension.OwnerID != userModel.ID {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "没有权限修改此扩展"})
		return
	}

	// 更新 extension 的字段
	extension.Title = newExtensionData.Title
	extension.Description = newExtensionData.Description
	extension.Content = newExtensionData.Content
	extension.Tag = newExtensionData.Tag
	extension.UpdatedAt = time.Now()

	// 将更新后的数据保存到数据库
	err = dao.Q.Extension.WithContext(c.Request.Context()).Save(extension)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "保存扩展时出错"})
		return
	}

	// 返回更新后的数据
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": extension})
}
