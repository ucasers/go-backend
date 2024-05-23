package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (server *Server) Initialize(DbUser, DbPassword, DbPort, DbHost, DbName string) error {
	// 创建连接字符串
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	// 连接数据库
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Cannot connect to postgres database: %v", err)
		return err
	}

	// 检查连接是否正常
	if err := db.DB().Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
		return err
	}

	// 设置 DB 字段
	server.DB = db

	// 初始化路由
	server.Router = gin.Default()

	return nil
}

func (server *Server) Run(addr string) {
	// 在退出时关闭数据库连接
	defer server.DB.Close()

	// 启动服务器
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
