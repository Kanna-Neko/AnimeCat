package controllers

import (
	"AnimeCat/models/mongodb"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSettingHandler(c *gin.Context) {
	res, err := mongodb.GetSetting()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": fmt.Sprintf("server error: %s", err.Error()),
		})
		log.Printf("server error: %s", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    res,
	})
}

func GetWallPaperHandler(c *gin.Context) {
	setting, err := mongodb.GetWallPaper()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("server error: %s", err.Error()),
		})
		log.Printf("server error: %s", err.Error())
		return
	}
	c.Redirect(http.StatusFound, setting)
}
