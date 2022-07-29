package controllers

import (
	"AnimeCat/models/mongodb"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSettingHandler(c *gin.Context) {
	res, err := mongodb.GetSetting()
	res.Password = "就不告诉你"
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
	wallpaper, err := mongodb.GetWallPaper()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("server error: %s", err.Error()),
		})
		log.Printf("server error: %s", err.Error())
		return
	}
	c.Redirect(http.StatusFound, wallpaper)
}
func GetLogoHandler(c *gin.Context) {
	logoPath, err := mongodb.GetLogo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("server error: %s", err.Error()),
		})
		log.Printf("server error: %s", err.Error())
		return
	}
	c.Redirect(http.StatusFound, logoPath)
}

func PutSettingHandler(c *gin.Context) {
	setting, err := mongodb.GetSetting()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("server error: %s", err.Error()),
		})
		log.Printf("server error: %s", err.Error())
		return
	}
	pageSize, err := strconv.Atoi(c.PostForm("pageSize"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": fmt.Sprintf("pageSize must be int, error is %s", err),
		})
		return
	}
	if pageSize == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  http.StatusForbidden,
			"message": fmt.Sprintf("pageSize must greater then zero, you input %d", pageSize),
		})
		return
	}
	setting.PageSize = int64(pageSize)
	setting.Footer = c.PostForm("footer")
	setting.Theme = c.PostForm("theme")
	setting.Logo = c.PostForm("logo")
	setting.Wallpaper = c.PostForm("wallpaper")
	setting.BucketName = c.PostForm("bucketName")
	setting.BucketRegion = c.PostForm("bucketRegion")
	setting.EndPoint = c.PostForm("endPoint")
	setting.SecretId = c.PostForm("secretId")
	setting.SecretKey = c.PostForm("secretKey")
	setting.GlobalCSS = c.PostForm("globalCss")
	setting.GlobalJS = c.PostForm("globalJs")
	setting.Language = c.PostForm("language")
	err = mongodb.ModifySetting(setting)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("modify setting error: %s", err.Error()),
		})
		log.Println(err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "success",
		})
	}
}
