package model

import "time"

type Item struct {
	Id        string `bson:"_id"`
	User      string
	Text      string
	CreatedAt time.Time
}

type ItemService interface {
	CreateItem(item *Item) (*Item, error)
	GetItems() ([]Item, error)
	GetItemById(id string) (*Item, error)
}

type StatusError struct {
	Code   int
	Reason string
}

func (m StatusError) Error() string {
	return m.Reason
}
