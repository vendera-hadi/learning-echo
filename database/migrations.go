package database

import (
	"learnecho/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func GetMigrations(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create table
		{
			ID: "2020080201",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		{
			ID: "2020080202",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Book{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("books")
			},
		},
		// add column
		{
			ID: "2020080203",
			Migrate: func(tx *gorm.DB) error {
				// when table already exists, it just adds fields as columns
				type User struct {
					Age int
				}
				return tx.AutoMigrate(&User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn("users", "age")
			},
		},
	})
}
