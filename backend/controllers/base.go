package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ucasers/go-backend/backend/middlewares"
	"github.com/ucasers/go-backend/dao"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (server *Server) Initialize(DbUser, DbPassword, DbPort, DbHost, DbName string) error {
	// 创建连接字符串
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	// 连接数据库
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot connect to postgres database: %v", err)
		return err
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("Error create Sql Object: %v", err)
		return err
	}

	// 检查连接是否正常
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
		return err
	}

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	//生成文件
	//g := gen.NewGenerator(gen.Config{
	//	OutPath:       "./dao",
	//	Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
	//	FieldNullable: true,
	//})
	//g.UseDB(db)
	//g.ApplyBasic(models.User{}, models.Extension{}, models.CipherPair{})
	//g.Execute()
	//
	//err = db.AutoMigrate(&models.User{}, models.Extension{}, models.CipherPair{})
	//if err != nil {
	//	log.Fatalf("Migrate error: %v", err)
	//	return err
	//}

	dao.SetDefault(db)
	// 设置 DB 字段
	server.DB = db

	// 初始化路由
	server.Router = gin.Default()
	server.Router.Use(middlewares.CORSMiddleware())

	server.initializeRoutes()
	return nil
}

func (server *Server) Run(addr string) {
	// 在退出时关闭数据库连接
	defer func(DB *gorm.DB) {
		sqlDB, _ := DB.DB()
		err := sqlDB.Close()
		if err != nil {

		}
	}(server.DB)

	// 启动服务器
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
