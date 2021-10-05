package service

import (
	"log"
	"net/http"
	"time"

	"github.com/deathdayss/flip-back-end/middleware"
	"github.com/deathdayss/flip-back-end/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func generateToken(c *gin.Context, p models.Person) {
	j := &middleware.JWT{
		[]byte("flip"),
	}
	claims := models.UserClaims{
		Email: p.Email,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,    // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*2, // 过期时间 2h
			Issuer:    "flip",                      // 签名的发行者
		},
	}

	token, err := j.GenerateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 401,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	data := models.LoginResult{
		Person: p,
		Token:  token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "login successfully",
		"data":   data,
	})
	return
}

func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*models.UserClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "token is vilid",
			"data":   claims,
		})
	}
}
