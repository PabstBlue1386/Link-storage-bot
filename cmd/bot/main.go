package main

import (
	"fmt"
	"link-storage-bot/internal/link"
)

func main() {

	l1 := link.NewLink("https://vk.com", "VK", "social media", []string{"meeting", "chat", "entertainment"})
	l2 := link.NewLink("https://youtube.com", "Youtube", " video hosting", []string{"video", "entertainment"})
	l3 := link.NewLink("https://translate.google.com", "translate", "education", []string{"education", "helper"})

	storage := link.NewStorage()
	savedLink, err := storage.Add(l1)
	if err != nil {
		return
	}
	fmt.Printf("A link has been added!\n Title: %s, ID: %v", savedLink.Title, savedLink.ID)
	fmt.Println()

	savedLink, err = storage.Add(l2)
	if err != nil {
		return
	}
	fmt.Printf("A link has been added!\n Title: %s, ID: %v", savedLink.Title, savedLink.ID)
	fmt.Println()

	savedLink, err = storage.Add(l3)
	if err != nil {
		return
	}
	fmt.Printf("A link has been added!\n Title: %s, ID: %v", savedLink.Title, savedLink.ID)
	fmt.Println()

	links := storage.List()
	for _, l := range links {
		fmt.Printf("\nID: %v\nTitle: %s\nURL: %s\nCategory: %s\nTags: %v\nAddedAt: %v\n", l.ID, l.Title, l.URL, l.Category, l.Tags, l.AddedAt.Format("02.01.2006 15:04"))
	}

}
