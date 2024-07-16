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
	"strconv"
)

func (server *Server) AddCipherPair(c *gin.Context) {
	user, _ := c.Get("user")
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "请求参数错误"})
		return
	}
	var pair models.CipherPair
	if err := json.Unmarshal(body, &pair); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "请求参数错误"})
		return
	}
	// 确保类型转换成功
	userModel, _ := user.(*models.User)

	// 绑定 user 到 extension
	pair.OwnerID = userModel.ID

	err = dao.Q.CipherPair.
		WithContext(c.Request.Context()).Create(&pair)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "添加密钥对失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (server *Server) ModifyCipherPair(c *gin.Context) {
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
	var cipherPairData models.CipherPair
	if err := json.Unmarshal(body, &cipherPairData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "请求参数错误"})
		return
	}

	cipherPair, err := dao.Q.CipherPair.
		WithContext(c.Request.Context()).
		Where(dao.CipherPair.ID.Eq(cipherPairData.ID)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "找不到密钥对"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "查询密钥对时出错"})
		}
		return
	}

	// 验证当前用户是否为扩展的拥有者
	if cipherPair.OwnerID != userModel.ID {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "没有权限修改此密钥对"})
		return
	}

	// 更新 extension 的字段
	cipherPair.Key = cipherPairData.Key
	cipherPair.Name = cipherPairData.Name
	cipherPair.Pwd = cipherPairData.Pwd

	// 将更新后的数据保存到数据库
	err = dao.Q.CipherPair.WithContext(c.Request.Context()).Save(cipherPair)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "保存密钥对时出错"})
		return
	}

	// 返回更新后的数据
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": cipherPair})
}

// 删除密钥对

func (server *Server) DeleteCipherPair(c *gin.Context) {
	// 获取当前用户
	user, _ := c.Get("user")
	userModel, _ := user.(*models.User)

	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将数据转换为字符串
	idStr := string(data)

	// 尝试将字符串转换为 uint32
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// 将 id 转换为 uint32 类型
	uint32ID := uint32(id)

	cipherPair, err := dao.Q.CipherPair.
		WithContext(c.Request.Context()).
		Where(dao.CipherPair.ID.Eq(uint32ID)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "找不到密钥对"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "删除密钥对时出错"})
		}
		return
	}

	// 验证当前用户是否为扩展的拥有者
	if cipherPair.OwnerID != userModel.ID {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "没有权限修改此密钥对"})
		return
	}

	if err := server.DB.WithContext(c.Request.Context()).Delete(&cipherPair).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "删除密钥对时出错"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "密钥对已删除"})

}

// list add cipherpairs
func (server *Server) ListCipherPairs(c *gin.Context) {
	// 获取当前用户
	user, _ := c.Get("user")
	userModel, _ := user.(*models.User)

	// 查询当前用户的密钥对
	var cipherPairs []models.CipherPair
	err := server.DB.
		WithContext(c.Request.Context()).
		Where("owner_id = ?", userModel.ID).
		Find(&cipherPairs).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "查询密钥对时出错"})
		return
	}

	// 返回密钥对
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": cipherPairs})
}
