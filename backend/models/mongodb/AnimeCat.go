package mongodb

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	jaxleof         = primitive.ObjectID{'A', 'n', 'i', 'm', 'e', 'C', 'a', 't'}
	defaultAnimeCat = AnimeCat{
		ID:         jaxleof,
		Name:       "",
		UpdateTime: time.Now().UnixMilli(),
		Size:       0,
		IsDir:      true,
		DirChild:   []Cat{},
		ObjChild:   []Cat{},
	}
)

func InitAnimeCat() error {
	var filter = bson.M{"_id": jaxleof}
	err := client.Database("AnimeCat").Collection("AnimeCat").FindOne(context.TODO(), filter).Err()
	if err == mongo.ErrNoDocuments {
		_, err := client.Database("AnimeCat").Collection("AnimeCat").InsertOne(context.TODO(), defaultAnimeCat)
		if err == nil {
			log.Println("database AnimeCat inited")
		}
		return err
	} else if err == nil {
		return nil
	} else {
		return err
	}
}

func GetAnimeCatDir(path []string) (AnimeCat, error) {
	var filter = bson.M{"_id": jaxleof}
	var cat AnimeCat
	err := client.Database("AnimeCat").Collection("AnimeCat").FindOne(context.TODO(), filter).Decode(&cat)
	if err != nil {
		return AnimeCat{}, nil
	}
	for i := 0; i < len(path); i++ {
		err := cat.switchDir(path[i])
		if err != nil {
			return AnimeCat{}, err
		}
	}
	return cat, nil
}
func GetAnimeCatObj(path []string) (AnimeCat, error) {
	var filter = bson.M{"_id": jaxleof}
	var cat AnimeCat
	err := client.Database("AnimeCat").Collection("AnimeCat").FindOne(context.TODO(), filter).Decode(&cat)
	if err != nil {
		return AnimeCat{}, nil
	}
	for i := 0; i < len(path)-1; i++ {
		err := cat.switchDir(path[i])
		if err != nil {
			return AnimeCat{}, err
		}
	}
	err = cat.switchObj(path[len(path)-1])
	if err != nil {
		return AnimeCat{}, err
	}
	return cat, err
}

func (cat *AnimeCat) switchDir(path string) error {
	var found = false
	for i := 0; i < len(cat.DirChild); i++ {
		if cat.DirChild[i].Name == path {
			found = true
			var err error
			*cat, err = GetAnimeCat(cat.DirChild[i].ID)
			if err != nil {
				return err
			}
		}
	}
	if !found {
		return errors.New("path can't been founded")
	}
	return nil
}

func (cat *AnimeCat) switchObj(path string) error {
	var found = false
	for i := 0; i < len(cat.ObjChild); i++ {
		if cat.DirChild[i].Name == path {
			found = true
			var err error
			*cat, err = GetAnimeCat(cat.DirChild[i].ID)
			if err != nil {
				return err
			}
		}
	}
	if !found {
		return errors.New("path can't been founded")
	}
	return nil
}

func GetAnimeCat(id primitive.ObjectID) (AnimeCat, error) {
	var res AnimeCat
	var filter = bson.M{
		"_id": id,
	}
	err := client.Database("AnimeCat").Collection("AnimeCat").FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return res, err
	} else {
		return res, nil
	}
}
