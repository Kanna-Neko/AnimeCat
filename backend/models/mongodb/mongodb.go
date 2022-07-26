package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client         *mongo.Client
	defaultSetting = Setting{
		BucketName:   "",
		BucketRegion: "",
		EndPoint:     "",
		Footer:       "",
		GlobalCSS:    "",
		GlobalJS:     "",
		Language:     "简体中文",
		Logo:         "",
		PageSize:     20,
		SecretId:     "",
		SecretKey:    "",
		Theme:        "light",
		Wallpaper:    "",
		WebsiteTitle: "AnimeCat",
		Aim:          "system",
	}
)

func init() {
	err := InitMongodbClient()
	if err != nil {
		log.Fatal(err)
	}
	err = InitSetting()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("setting config inited")
}

func InitMongodbClient() error {
	account := os.Getenv("mongodb_account")
	password := os.Getenv("mongodb_password")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{Password: password, Username: account})

	// 连接到MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	log.Println("Connected to MongoDB!")
	return nil
}

func GetSetting() (Setting, error) {
	var res Setting
	var filter = bson.M{
		"aim": "system",
	}
	collection := client.Database("AnimeCat").Collection("setting")
	err := collection.FindOne(context.TODO(), filter).Decode(&res)
	return res, err
}

func InitSetting() error {
	var filter = bson.M{"aim": "system"}
	err := client.Database("AnimeCat").Collection("setting").FindOne(context.TODO(), filter).Err()
	if err == mongo.ErrNoDocuments {
		_, err := client.Database("AnimeCat").Collection("setting").InsertOne(context.TODO(), defaultSetting)
		return err
	} else {
		return err
	}
}
