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
	sqlDB, _ := gorm.Open("mysql", "root:Cptbtptp1790340626.@tcp(localhost:3306)/comp")
	return sqlDB
}

func (db *Db) Close() {
	db.MsClient.Close()
}

func (db *Db) AutoCreateTable() {
	db.MsClient.AutoMigrate(&Person{})
	db.MsClient.AutoMigrate(&Game{})
	db.MsClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Like{})
	db.MsClient.Model(&Like{}).AddIndex("user_id", "game_id")
	db.MsClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Collect{})
	db.MsClient.Model(&Collect{}).AddIndex("user_id", "game_id")
}
