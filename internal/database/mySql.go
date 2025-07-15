package database

import (
	"fmt"
	"github.com/spf13/viper"
	mySqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func NewMySqlGorm() (*gorm.DB, error) {
	Host := viper.GetString("db.host")
	Port := viper.GetUint16("db.port")
	Username := viper.GetString("db.username")
	Password := os.Getenv("DB_PASSWORD")
	DBName := viper.GetString("db.dbname")
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FDushanbe",
		Username, Password, Host, Port, DBName)

	conn, err := gorm.Open(mySqlDriver.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Printf("❌ MySQL -> Open error: %s", err.Error())
		return nil, err
	}

	log.Println("✅ MySQL Connection success:", Host)
	return conn, nil
}
