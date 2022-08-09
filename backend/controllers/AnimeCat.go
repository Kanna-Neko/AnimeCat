package controllers

import (
	"AnimeCat/models/mongodb"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	} else {
		cat, err := mongodb.GetAnimeCatObj(paths)
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

func PostAnimeCatDir(c *gin.Context) {
	var info gin.H
	err := c.BindJSON(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	_, exist := info["_id"]
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "_id field is not founded",
		})
		return
	}
	_, exist = info["name"]
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "name field is not founded",
		})
		return
	}
	id, err := primitive.ObjectIDFromHex(info["_id"].(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	cat, err := mongodb.GetAnimeCat(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	err = cat.CreateDir(info["name"].(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "dir create process is success",
			"data":    cat,
		})
	}
}

func AnimeCatIDHandler(c *gin.Context) {
	var info map[string]string
	c.BindJSON(&info)
	val, exist := info["_id"]
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "want _id field, but null",
		})
		return
	}
	id, err := primitive.ObjectIDFromHex(val)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "id format wrong, error: " + err.Error(),
		})
		return
	}
	cat, err := mongodb.GetAnimeCat(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "error: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    cat,
	})
}
