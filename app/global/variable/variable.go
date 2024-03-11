package variable

import (
	"gorm.io/gorm"
)

var (
	BasePath           = "/Users/mainjay/GolandProjects/go-backend" // 定义项目的根目录，用户根据自己的位置修改
	EventDestroyPrefix = "Destroy_"                                 //  程序退出时需要销毁的事件前缀
	ConfigKeyPrefix    = "Config_"                                  //  配置文件键值缓存时，键的前缀
	DateFormat         = "2006-01-02 15:04:05"                      //  设置全局日期时间格式

	// ConfigYml 全局日志指针
	//ZapLog *zap.Logger

	// ConfigYml 全局配置文件
	ConfigYml     map[string]any // 全局配置文件指针
	ConfigGormYml map[string]any // 全局配置文件指针

	// GormDbPostgreSql GormDbMysql gorm 数据库客户端，在 bootstrap>init 文件，进行初始化即可使用
	GormDbPostgreSql *gorm.DB // 全局gorm的客户端连接

	////雪花算法全局变量
	//SnowFlake snowflake_interf.InterfaceSnowFlake

	//websocket
	WebsocketHub              interface{}
	WebsocketHandshakeSuccess = `{"code":200,"msg":"ws连接成功","data":""}`
	WebsocketServerPingMsg    = "Server->Ping->Client"

	////casbin 全局操作指针
	//Enforcer *casbin.SyncedEnforcer

	//  用户自行定义其他全局变量 ↓

)
