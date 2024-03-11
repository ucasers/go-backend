package bootstrap

//
//import (
//	"go-backend/app/global/my_errors"
//	"go-backend/app/global/variable"
//	"go-backend/app/utils/config"
//	"log"
//	"os"
//)
//
//// 检查项目必须的非编译目录是否存在，避免编译后调用的时候缺失相关目录
//func checkRequiredFolders() {
//	//1.检查配置文件是否存在
//	if _, err := os.Stat(variable.BasePath + "/config/config.yml"); err != nil {
//		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
//	}
//	if _, err := os.Stat(variable.BasePath + "/config/gorm.yml"); err != nil {
//		log.Fatal(my_errors.ErrorsConfigGormNotExists + err.Error())
//	}
//}
//
//func init() {
//	// 1. 初始化 项目根路径，参见 variable 常量包，相关路径：app\global\variable\variable.go
//
//	//2.检查配置文件以及日志目录等非编译性的必要条件
//	checkRequiredFolders()
//
//	// 4.启动针对配置文件(config.yml、gorm.yml)变化的监听， 配置文件操作指针，初始化为全局变量
//	variable.ConfigYml = config.CreateYamlFactory("config")
//	variable.ConfigGormYml = config.CreateYamlFactory("gorm")
//
//	// 6.根据配置初始化 gorm mysql 全局 *gorm.Db
//	if variable.ConfigGormv2Yml.GetInt("Gormv2.Mysql.IsInitGlobalGormMysql") == 1 {
//		if dbMysql, err := gorm_v2.GetOneMysqlClient(); err != nil {
//			log.Fatal(my_errors.ErrorsGormInitFail + err.Error())
//		} else {
//			variable.GormDbMysql = dbMysql
//		}
//	}
//	// 根据配置初始化 gorm sqlserver 全局 *gorm.Db
//	if variable.ConfigGormv2Yml.GetInt("Gormv2.Sqlserver.IsInitGlobalGormSqlserver") == 1 {
//		if dbSqlserver, err := gorm_v2.GetOneSqlserverClient(); err != nil {
//			log.Fatal(my_errors.ErrorsGormInitFail + err.Error())
//		} else {
//			variable.GormDbSqlserver = dbSqlserver
//		}
//	}
//	// 根据配置初始化 gorm postgresql 全局 *gorm.Db
//	if variable.ConfigGormv2Yml.GetInt("Gormv2.PostgreSql.IsInitGlobalGormPostgreSql") == 1 {
//		if dbPostgre, err := gorm_v2.GetOnePostgreSqlClient(); err != nil {
//			log.Fatal(my_errors.ErrorsGormInitFail + err.Error())
//		} else {
//			variable.GormDbPostgreSql = dbPostgre
//		}
//	}
//
//	// 7.雪花算法全局变量
//	variable.SnowFlake = snow_flake.CreateSnowflakeFactory()
//
//	// 8.websocket Hub中心启动
//	if variable.ConfigYml.GetInt("Websocket.Start") == 1 {
//		// websocket 管理中心hub全局初始化一份
//		variable.WebsocketHub = core.CreateHubFactory()
//		if Wh, ok := variable.WebsocketHub.(*core.Hub); ok {
//			go Wh.Run()
//		}
//	}
//
//	// 9.casbin 依据配置文件设置参数(IsInit=1)初始化
//	if variable.ConfigYml.GetInt("Casbin.IsInit") == 1 {
//		var err error
//		if variable.Enforcer, err = casbin_v2.InitCasbinEnforcer(); err != nil {
//			log.Fatal(err.Error())
//		}
//	}
//	//10.全局注册 validator 错误翻译器,zh 代表中文，en 代表英语
//	if err := validator_translation.InitTrans("zh"); err != nil {
//		log.Fatal(my_errors.ErrorsValidatorTransInitFail + err.Error())
//	}
//}
