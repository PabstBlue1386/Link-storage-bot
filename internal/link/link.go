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
