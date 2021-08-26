package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Person struct {
    Email string `gorm:"primary_key;not null"`
    Nickname string `json:"nickname"`
    Password string `json:"password"`
}
type Dependency struct {
	db *gorm.DB
}
var d *Dependency
func initial() *Dependency{
	sqlDB, _ := gorm.Open("mysql", "root:Cptbtptp1790340626`@tcp(127.0.0.1:3306)/comp")
	return &Dependency{
		db: sqlDB,
	}
}

func main() {
	r := gin.Default()
	r.POST("/login", login) // http://localhost:8084/register
	r.POST("/register", register)
    d = initial()
	defer d.db.Close()
    d.db.AutoMigrate(&Person{})
	r.Run(":8084")
}

func login(c *gin.Context) {
	//db, _ := gorm.Open("mysql", "root:Cptbtptp1790340626@tcp(127.0.0.1:3306)/comp")
    //defer db.Close()
	email, ok1 := c.GetPostForm("email")
	password, ok2 := c.GetPostForm("password")
	if !ok1 || !ok2 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email or password is missing",
		})
		return
	}
	var user Person
	if err := d.db.Where("email = ? AND password = ?", string(email), string(password)).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "email or password is wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "login successfully",
	})
}

func register(c *gin.Context) {
	email, ok1 := c.GetPostForm("email")
	password, ok2 := c.GetPostForm("password")
	nickname, ok3 := c.GetPostForm("nickname")
	if !ok1 || !ok2 || !ok3 || len(email) == 0 || len(password) == 0 || len(nickname) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email, nickname or password is missing",
		})
		return
	}
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	nickname = strings.TrimSpace(nickname)
	var p Person
	if err := d.db.Where("email = ?", string(email)).First(&p).Error; err == nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email has been used",
		})
		return
	}
	p = Person{
		Email: email,
		Password: password,
		Nickname: nickname,
	}
	// &取地址，*（&p）=p
	if err := d.db.Create(&p).Error; err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "can not register",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "register successfully",
	})
}