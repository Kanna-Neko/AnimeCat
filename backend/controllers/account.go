package controllers

import (
	"AnimeCat/models/jwt"
	"AnimeCat/models/mongodb"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	password := c.PostForm("password")
	jaxleof, err := mongodb.GetPassword()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("server error: %s", err),
		})
		return
	}
	if jaxleof == password {
		token, err := jwt.GenRegisteredClaims()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": fmt.Sprintf("server error: %s", err),
			})
		} else {
			c.JSON(200, gin.H{
				"status":  200,
				"message": "success",
				"data": gin.H{
					"token": token,
				},
			})
		}
	} else {
		c.JSON(http.StatusForbidden,gin.H{
			"status": 403,
			"message": "password error",
		})
	}
}
