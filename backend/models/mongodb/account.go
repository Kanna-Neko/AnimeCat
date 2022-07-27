package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func GetPassword() (string, error) {
	var res Setting
	var filter = bson.M{
		"aim": "system",
	}
	collection := client.Database("AnimeCat").Collection("setting")
	err := collection.FindOne(context.TODO(), filter).Decode(&res)
	return res.Password, err
}
