package link

import "time"

type Link struct {
	ID       int       `json:"id"`
	URL      string    `json:"url"`
	Title    string    `json:"title"`
	Category string    `json:"category"`
	Tags     []string  `json:"tags"`
	AddedAt  time.Time `json:"added_at"`
}

func NewLink(url, title, category string, tags []string) Link {
	return Link{
		URL:      url,
		Title:    title,
		Category: category,
		Tags:     tags,
	}
}
