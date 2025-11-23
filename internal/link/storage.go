package link

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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

func (storage *Storage) List() []Link {
	return storage.Links
}

func (storage *Storage) Save(path string) error {
	if storage == nil {
		return fmt.Errorf("emtpy storage")
	}

	data, err := json.MarshalIndent(storage, "", "  ")
	if err != nil {
		return fmt.Errorf("error converting to byte %s\n", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file%s\n", err)
	}

	return nil
}

func (storage Storage) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file not exist%s\n", err)
		}
		return fmt.Errorf("failed to read file%s\n", err)
	}
	err = json.Unmarshal(data, &storage)
	if err != nil {
		return fmt.Errorf("error unmarhsal data%s\n", err)
	}

	return nil
}
