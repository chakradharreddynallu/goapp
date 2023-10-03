package config

import (
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnectionString() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3307"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}

	password := os.Getenv("DB_PASS")
	if password == "" {
		password = "@root"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "moviestore"
	}
	// d, err := gorm.Open("mysql", "manda:supraja@tcp(%s:%s)/movie?charset=utf8&parseTime=True&loc=Local")

	return fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
}
