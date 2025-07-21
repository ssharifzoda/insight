package database

import (
	"fmt"
	"github.com/spf13/viper"
	mySqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"insight/internal/models"
	"insight/pkg/consts"
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
	SqlRunner(conn)
	log.Println("✅ MySQL Connection success:", Host)
	return conn, nil
}

func SqlRunner(conn *gorm.DB) {
	var check *models.Migration
	err := conn.Table("migrations").Last(&check).Error
	if err != nil {
		log.Fatal(err)
	}
	if check.Batch == 0 {
		tx := conn.Begin()
		fileByte, err := os.ReadFile(consts.SqlFilesPath + check.Migration)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		err = tx.Exec(string(fileByte)).Error
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		err = tx.Table("migrations").Where("id", check.Id).UpdateColumn("batch", "1").Error
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}
	return
}
