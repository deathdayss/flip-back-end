package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Db struct {
	MsClient *gorm.DB
}

var DbClient *Db

func (db *Db) Init() {
	DbClient = &Db{
		MsClient: InitMySql(),
	}
	DbClient.AutoCreateTable()
}

func InitMySql() *gorm.DB {
	// This is used to connect the MySQL on server
	//sqlDB, _ := gorm.Open("mysql", "root:Cptbtptp1790340626.@tcp(127.0.0.1:3306)/comp")

	sqlDB, _ := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/version2")
	return sqlDB
}

func (db *Db) Close() {
	db.MsClient.Close()
}

func (db *Db) AutoCreateTable() {
	db.MsClient.AutoMigrate(&Person{})
	db.MsClient.AutoMigrate(&Game{})
	db.MsClient.AutoMigrate(&ProductInfo{})
	db.MsClient.AutoMigrate(&Code{})

	db.MsClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Like{})
	db.MsClient.Model(&Like{}).AddIndex("user_id", "game_id")

	db.MsClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Share{})
	db.MsClient.Model(&Share{}).AddIndex("game_id", "share_num")
}
