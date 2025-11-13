package link

import (
	"errors"
	"time"
)

type Storage struct {
	Links  []Link `json:"links"`
	NextID int    `json:"nextID"`
}

func (storage *Storage) Add(link Link) (Link, error) {

	if link.Title == "" {
		return Link{}, errors.New("Empty title")
	}
	if link.URL == "" {
		return Link{}, errors.New("Empty URL")
	}
	if link.Category == "" {
		return Link{}, errors.New("Empty category")
	}
	if len(link.Tags) < 1 {
		return Link{}, errors.New("Minimum one tag")
	}

	link.ID = storage.NextID
	link.AddedAt = time.Now()

	storage.Links = append(storage.Links, link)
	storage.NextID++
	return link, nil
}
