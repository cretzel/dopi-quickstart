package service

import (
	"context"
	"dopi-quickstart/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ItemService struct {
	client     *Client
	collection *mongo.Collection
}

func NewItemService(client *Client) *ItemService {
	collection := client.GetCollection("dopi-quickstart", "items")
	quickstartService := ItemService{client: client, collection: collection}
	return &quickstartService
}

func (u *ItemService) CreateItem(item *model.Item) (*model.Item, error) {
	_, err := u.collection.InsertOne(context.TODO(), bson.D{{"text", item.Text}, {"createdAt", time.Now()}})
	if err != nil {
		return item, err
	}
	return u.GetItemById(item.Id)
}

func (u *ItemService) GetItemById(id string) (*model.Item, error) {
	item := model.Item{}
	err := u.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&item)
	return &item, err
}

func (u *ItemService) GetItems() ([]model.Item, error) {
	var items []model.Item
	cursor, err := u.collection.Find(context.TODO(), bson.M{}, options.Find().SetSort(bson.D{{"createdAt", -1}}))
	if err != nil {
		return items, err
	}
	if err = cursor.All(context.TODO(), &items); err != nil {
		return items, err
	}

	return items, nil
}
