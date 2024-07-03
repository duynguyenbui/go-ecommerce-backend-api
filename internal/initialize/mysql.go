package initialize

import (
	"fmt"
	"time"

	"github.com/duynguyenbui/go-ecommerce-backend-api/global"
	"github.com/duynguyenbui/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErroPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.MySQL
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DbName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	checkErroPanic(err, "Failed to connect to mysql")

	global.Logger.Info("Success to connect to mysql")
	global.Mdb = db

	// Set Pool
	SetPool()
	MigrateTables()
}

func SetPool() {
	m := global.Config.MySQL
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Println("Failed to set pool")
	}

	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func MigrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
	}
}
