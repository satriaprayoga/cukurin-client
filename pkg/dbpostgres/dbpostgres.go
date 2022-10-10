package dbpostgres

import (
	"cukurin-client/pkg/setting"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Conn *gorm.DB

func Setup() {
	now := time.Now()
	var err error

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable Timezone=Asia/Jakarta",
		setting.FileConfigSetting.Database.Host,
		setting.FileConfigSetting.Database.User,
		setting.FileConfigSetting.Database.Password,
		setting.FileConfigSetting.Database.Name,
		setting.FileConfigSetting.Database.Port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      false,
		},
	)

	Conn, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.FileConfigSetting.Database.TablePrefix,
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		log.Printf("connection setup error: %v", err)
		panic(err)
	}

	sqlDB, err := Conn.DB()
	if err != nil {
		log.Printf("connection.setup DB err : %v", err)
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	timeSpent := time.Since(now)
	log.Printf("Config database is ready in %v", timeSpent)
}
