package initialize

import (
	"github.com/duynguyenbui/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	// Load Configuration
	LoadConfig()
	// Init Logger
	InitLogger()
	// Init Mysql
	global.Logger.Info("Config Log ok!!", zap.String("ok", "success"))
	InitMysql()
	// Init Redis
	InitRedis()
	// Init Router
	r := InitRouter()

	r.Run(":8002")
}
