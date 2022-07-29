package controllers

import (
	"AnimeCat/models/mongodb"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AnimeCatHandler(c *gin.Context) {
	path := c.Param("path")
	paths := strings.Split(path[1:], "/")
	if path[len(path)-1] == '/' {
		cat, err := mongodb.GetAnimeCatDir(paths[:len(paths)-1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"status":  200,
				"message": "success",
				"data": gin.H{
					"cat": cat,
				},
			})
		}
	}

}
