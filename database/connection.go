package database

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/rahadirafsanjani/notes_app/config"
	"github.com/rahadirafsanjani/notes_app/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Catched Error: ", err)
	}

	// Create the database connection string (DSN)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Models to migrate
	models := []interface{}{
		&model.Notes{}, // Add all the models you want to migrate here
		&model.Labels{},
		&model.LabelNotes{},
		// Add other models here as needed
	}

	// Automatically migrate each model
	for _, model := range models {
		if err := DB.AutoMigrate(model); err != nil {
			log.Fatalf("failed to migrate model %v: %v", reflect.TypeOf(model).Elem(), err)
		}
	}

	fmt.Println("Database Migrated")
}
