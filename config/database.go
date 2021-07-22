package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb(host string, user string, password string, dbName string, port string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable " +
		"TimeZone=Asia/Shanghai", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func CloseConnection(db *gorm.DB)  {
	dbSQL, err := db.DB()
	if err != nil{
		panic(err)
	}

	dbSQL.Close()
}
