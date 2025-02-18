package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var msDB *gorm.DB

func InitDatabase(dsn string) error {
	db, err := gorm.Open(
		sqlserver.Open(dsn),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             time.Second * 10,
					LogLevel:                  logger.Silent,
					IgnoreRecordNotFoundError: false,
					ParameterizedQueries:      true,
					Colorful:                  false,
				},
			),
			CreateBatchSize: 300,
		},
	)

	if err != nil {
		return fmt.Errorf("sqlserver 数据库链接失败: %v", err)
	}

	msDB = db

	return nil
}

// func generateTimestampID() int64 {
// 	return time.Now().UnixMicro() - time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC).UnixMicro()
// }
