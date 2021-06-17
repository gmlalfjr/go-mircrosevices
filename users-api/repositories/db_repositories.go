package repositories

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	Client *gorm.DB
)



func init()  {
	dsn := "root:root1234@tcp(127.0.0.1:3306)/users_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Client, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}


	log.Println("database successfully configured")
	// Get generic database object sql.DB to use its functions
	sqlDB, err := Client.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour);
}