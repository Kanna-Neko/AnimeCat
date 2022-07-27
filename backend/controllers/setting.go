package controllers

import (
	"AnimeCat/models/mongodb"
	"fmt"
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
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    res,
	})
}
