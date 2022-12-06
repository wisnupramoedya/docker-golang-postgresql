package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBUser     = "user_pg"
	DBPassword = "user_pg_123"
	DBName     = "movies"
	DBHost     = "postgres_container"
	DBPort     = 5432
)

var DB *gorm.DB

func OpenConnection() *gorm.DB {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBName)

	fmt.Println(psqlInfo)
	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connect to %s successfully\n", DBName)

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
