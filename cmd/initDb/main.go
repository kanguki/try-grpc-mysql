package main

import (
	"log"
	"os"

	"github.com/kanguki/go-grpc-mysql/internal/core/movie"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	mysqlConnectStr := os.Getenv("MYSQL_URL")
	if mysqlConnectStr == "" {
		log.Fatalln("Empty mysql url")
	}
	db, err := gorm.Open(mysql.Open(mysqlConnectStr))
	if err != nil {
		log.Fatalf("Error connect mysql: %v", err)
	}
	err = db.AutoMigrate(&movie.Movie{})
	if err != nil {
		log.Fatalf("Error autoMigrate Movie: %v", err)
	}
}
