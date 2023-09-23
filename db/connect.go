package db

import (
	"log"
	"os" // 環境変数などのサポート

	"gorm.io/driver/postgres" // GORMのPostgreSQL用のドライバ
	"gorm.io/gorm"            // GORMのメインパッケージ
	"gorm.io/gorm/logger"
)

func ConnectGORM(logFile *os.File) *gorm.DB {
	// データベースへの接続
	dsn := os.Getenv("DATABASE_URL")
	// 接続の初期化
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(logFile, "\r\n", log.LstdFlags), // ログをファイルに出力するLoggerを設定
			logger.Config{
				LogLevel: logger.Info, // ログレベルを指定
			},
		),
	})

	if err != nil {
		panic(err.Error())
	}

	return db
}
