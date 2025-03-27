package infrastructure

import (
	"fmt"
	"github.com/rodrigoferra/ports-blog/internal/infrastructure/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitPostgresDatabase(
	host string,
	port int,
	user string,
	password string,
	dbname string,
) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate domain models
	err = db.AutoMigrate(&entities.PostEntity{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	err = db.AutoMigrate(&entities.CommentEntity{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
