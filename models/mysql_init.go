package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
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

	sqlDB, _ := gorm.Open("mysql", "root:19960822@tcp(localhost:3306)/comp")
	//sqlDB, _ := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/comp")
	return sqlDB
}

func (db *Db) AutoCreateTable() {
	db.MsClient.AutoMigrate(&Person{})
	db.MsClient.AutoMigrate(&Game{})
	db.MsClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Answer{})
	//db.MsClient.Model(&Answer{}).AddIndex("user_id")
	db.MsClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Like{})
	//db.MsClient.Model(&Like{}).AddIndex("user_id", "game_id")
	db.MsClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Collect{})
	//db.MsClient.Model(&Collect{}).AddIndex("user_id", "game_id")
	db.MsClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&PersonImg{})
	db.MsClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&PersonDetail{})
	//db.MsClient.Model(&PersonImg{}).AddIndex("uid")
	db.MsClient.AutoMigrate(&Code{})
}
