package mongodb

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
	SecretId     string `json:"SecretId" bson:"secretId"`
	SecretKey    string `json:"SecretKey" bson:"secretKey"`
	Theme        string `json:"theme" bson:"theme"`
	Wallpaper    string `json:"wallpaper" bson:"wallpaper"`
	WebsiteTitle string `json:"websiteTitle" bson:"websiteTitle"`
	Aim          string `json:"aim" bson:"aim"`
}