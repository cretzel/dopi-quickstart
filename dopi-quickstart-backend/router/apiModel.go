package router

import (
	"dopi-quickstart/model"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ItemDto struct {
	Id        string    `json:"id,omitempty"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateItemDto struct {
	Text string `json:"text"`
}

func toItemDto(item *model.Item) *ItemDto {
	return &ItemDto{
		Id:        item.Id,
		Text:      item.Text,
		CreatedAt: item.CreatedAt,
	}
}

func createItemDtoToItem(createItem *CreateItemDto) *model.Item {
	return &model.Item{
		Text: createItem.Text,
	}
}

func toItemDtos(items []model.Item) []ItemDto {
	dtos := make([]ItemDto, len(items))
	for i, u := range items {
		dtos[i] = *toItemDto(&u)
	}
	return dtos
}

func Error(w http.ResponseWriter, code int, message string) {
	body, _ := json.Marshal(map[string]string{"error": message})
	w.WriteHeader(code)
	w.Write(body)
}

func StatusError(w http.ResponseWriter, err error) {
	log.Println(err)
	reason := http.StatusText(http.StatusInternalServerError)
	code := http.StatusInternalServerError

	switch err.(type) {
	case *model.StatusError:
		reason = err.(*model.StatusError).Reason
		code = err.(*model.StatusError).Code
	}
	body, _ := json.Marshal(map[string]string{"error": reason})
	w.WriteHeader(code)
	w.Write(body)
}
