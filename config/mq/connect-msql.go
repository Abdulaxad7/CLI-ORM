package mq

import (
	"Cli-Orm/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(con *config.DB) (*gorm.DB, error) {
	db, err := getConnection(&config.DB{
		DBUser:     con.DBUser,
		DBPassword: con.DBPassword,
		Port:       con.Port,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getConnection(con *config.DB) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(query(con)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func query(con *config.DB) string {

	return fmt.Sprintf("%s:%s@tcp(localhost:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		con.DBUser, con.DBPassword, con.Port)
}
