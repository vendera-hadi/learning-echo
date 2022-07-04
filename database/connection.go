package database

import (
	"fmt"
	"log"
	"sync"

	"learnecho/config"

	"gorm.io/driver/mysql"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var onceDb sync.Once

var instance *gorm.DB

func GetInstance() *gorm.DB {
	onceDb.Do(func() {
		databaseConfig := config.DatabaseNew().(*config.DatabaseConfig)
		// MYSQL
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			databaseConfig.Mysql.DbUsername,
			databaseConfig.Mysql.DbPassword,
			databaseConfig.Mysql.DbHost,
			databaseConfig.Mysql.DbPort,
			databaseConfig.Mysql.DbDatabase,
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		// PostgreSQL
		// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		// 	databaseConfig.Psql.DbHost,
		// 	databaseConfig.Psql.DbUsername,
		// 	databaseConfig.Psql.DbPassword,
		// 	databaseConfig.Psql.DbDatabase,
		// 	databaseConfig.Psql.DbPort,
		// )
		// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("Could not connect to database :%v", err)
		}
		instance = db
	})
	return instance
}
