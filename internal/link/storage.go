package link

import (
	"errors"
	"strings"
	"time"
)

type Storage struct {
	Links  []Link `json:"links"`
	NextID int    `json:"nextID"`
}

func NewStorage() *Storage {
	return &Storage{
		Links:  make([]Link, 0),
		NextID: 1,
	}
}

func (storage *Storage) Add(link Link) (Link, error) {

	link.Title = strings.Trim(link.Title, " ")
	if link.Title == "" {
		return Link{}, errors.New("empty title")
	} else {

	}

	link.URL = strings.TrimSpace(link.URL)
	if link.URL == "" {
		return Link{}, errors.New("empty URL")
	} else {
		for _, l := range storage.Links {
			if strings.EqualFold(l.URL, link.URL) {
				return Link{}, errors.New("existing URL")
			}
		}
	}

	link.Category = strings.Trim(link.Category, " ")
	if link.Category == "" {
		return Link{}, errors.New("empty category")
	} else {

	}

	if len(link.Tags) < 1 {
		return Link{}, errors.New("minimum one tag")
	} else {
		var result []string
		for _, tag := range link.Tags {
			t := strings.TrimSpace(tag)
			t = strings.ToUpper(t)
			if t != "" {
				result = append(result, t)
			}
		}
		if len(result) < 1 {
			return Link{}, errors.New("minimum one tag")
		}
		link.Tags = result
	}

	link.ID = storage.NextID
	link.AddedAt = time.Now()

	storage.Links = append(storage.Links, link)
	storage.NextID++

	return link, nil
}

func (s *Storage) List() []Link {
	return s.Links
}
