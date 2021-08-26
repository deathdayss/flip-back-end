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
	sqlDB, _ := gorm.Open("mysql", "root:Cptbtptp1790340626@tcp(127.0.0.1:3306)/comp")
	return sqlDB
}

func (db *Db) Close() {
	db.MsClient.Close()
}

func (db *Db) AutoCreateTable() {
	db.MsClient.AutoMigrate(&Person{})
}