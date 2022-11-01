package config

import (
	"echo-book/entity"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDatabaseConnection berfungsi  untuk koneksi ke database
func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		panic("failed to load env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",

		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to create a connection to database")
	}

	//Membuat model database
	db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}

// CloseDatabaseConnection berfungsi untuk menutup koneksi database kita
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("failed to close connection database")
	}

	dbSQL.Close()

}

func SetupDatabaseConnectionTest() *gorm.DB {
	// err := godotenv.Load()

	// if err != nil {
	// 	panic("failed to load env")
	// }

	dbUser := "root"             //os.Getenv("DB_USER")
	dbPass := "root123"          //os.Getenv("DB_PASS")
	dbHost := "localhost"        //os.Getenv("DB_HOST")
	dbPort := "3306"             //os.Getenv("DB_PORT")
	dbTest := "golang_book_test" //os.Getenv("DB_TEST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",

		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbTest,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to create a connection to database")
	}

	//Membuat model database
	db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}
