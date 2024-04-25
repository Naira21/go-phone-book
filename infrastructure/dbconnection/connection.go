package dbconnection

import (
	"fmt"
	"go-phone-book/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDBConnection(config config.Config) (*gorm.DB, error) {
	fmt.Println(config.DBUsername, config.DBHost, config.DBPort, config.DBName)

	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBHost, config.DBPort, config.DBName)
	// dsn := fmt.Sprintf("root:@tcp(localhost:3306)/go-phone-book?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("connected to DB")
	return db, nil
}
