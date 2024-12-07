package db

import (
	"database/sql"
	"log"
)

func Seed(db *sql.DB) {
	snippets := []struct {
		Title   string
		Content string
		Expires string
	}{
		{
			Title:   "An old silent pond",
			Content: "An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō",
			Expires: "2024-12-13 16:04:07",
		},
		{
			Title:   "Over the wintry forest",
			Content: "Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki",
			Expires: "2025-01-05 16:04:07",
		},
		{
			Title:   "First autumn morning",
			Content: "First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo",
			Expires: "2025-12-06 16:04:07",
		},
	}

	for _, snippet := range snippets {
		_, err := db.Exec(`INSERT INTO snippets (title, content, expires) VALUES (?, ?, ?)`, snippet.Title, snippet.Content, snippet.Expires)

		if err != nil {
			log.Fatalf("Error seeding users: %v", err)
		}
	}
	log.Println("Seeding completed successfully!")
}
