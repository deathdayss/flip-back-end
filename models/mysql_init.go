package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	//config := "root:19960822@tcp(localhost:3306)/comp"
	config := "root:Cptbtptp1790340626.@tcp(localhost:3306)/comp?parseTime=true&loc=Local"
	sqlDB, _ := gorm.Open(mysql.Open(config), &gorm.Config{})

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
	//db.MsClient.Model(&PersonImg{}).AddIndex("uid")
	db.MsClient.AutoMigrate(&Code{})
	db.MsClient.AutoMigrate(&Comment{})
	db.MsClient.AutoMigrate(&GameRescale{})
	db.MsClient.AutoMigrate(&Zone{})
}
