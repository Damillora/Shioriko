package database

import (
	"log"

	"github.com/Damillora/Shioriko/pkg/config"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func Initialize() {
	dsn := config.CurrentConfig.PostgresDatabase
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{})
	db.AutoMigrate(&TagType{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Blob{})
	db.AutoMigrate(&Post{})

	var general TagType
	result := db.Where("name = ?", "general").First(&general)
	if result.Error != nil {
		db.Create(&TagType{
			Name: "general",
		})
	}

	if err != nil {
		log.Fatal("Unable to connect to database" + err.Error())
	}

	DB = db
}
