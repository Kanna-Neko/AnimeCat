package mongodb

import "go.mongodb.org/mongo-driver/bson/primitive"

type Setting struct {
	BucketName   string `json:"bucketName" bson:"bucketName"`
	BucketRegion string `json:"bucketRegion" bson:"bucketRegion"`
	EndPoint     string `json:"endPoint" bson:"endPoint"`
	Footer       string `json:"footer" bson:"footer"`
	GlobalCSS    string `json:"globalCss" bson:"globalCss"`
	GlobalJS     string `json:"globalJs" bson:"globalJs"`
	Language     string `json:"language" bson:"language"`
	Logo         string `json:"logo" bson:"logo"`
	PageSize     int64  `json:"pageSize" bson:"pageSize"`
	SecretId     string `json:"secretId" bson:"secretId"`
	SecretKey    string `json:"secretKey" bson:"secretKey"`
	Theme        string `json:"theme" bson:"theme"`
	Wallpaper    string `json:"wallpaper" bson:"wallpaper"`
	WebsiteTitle string `json:"websiteTitle" bson:"websiteTitle"`
	Password     string `json:"password" bson:"password"`
	Aim          string `json:"aim" bson:"aim"`
}

type AnimeCat struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	UpdateTime int64              `json:"updateTime" bson:"updateTime"`
	Size       int64              `json:"size" bson:"size"`
	IsDir      bool               `json:"isDir" bson:"isDir"`
	DirChild   []Cat              `json:"dirChild" bson:"dirChild"`
	ObjChild   []Cat              `json:"objChild" bson:"objChild"`
}

type Cat struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}
