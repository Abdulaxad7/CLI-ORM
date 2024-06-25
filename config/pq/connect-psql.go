package pq

import (
	"Cli-Orm/config"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(con *config.DB) (*gorm.DB, error) {
	db, err := getConnection(&config.DB{
		DBUser:     con.DBUser,
		DBName:     con.DBName,
		DBPassword: con.DBPassword,
		Port:       con.Port,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getConnection(con *config.DB) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(query(con)), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}

func query(con *config.DB) string {
	return fmt.Sprintf("user=%s password=%s database=%s host=localhost port=%s sslmode=disable",
		con.DBUser, con.DBPassword, con.DBName, con.Port)
}
