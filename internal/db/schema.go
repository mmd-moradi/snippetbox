package db

import (
	"database/sql"
	"log"
)

func CreateTable(db *sql.DB) {
	queries := []string{
		`
			CREATE TABLE snippets (
				id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
				title VARCHAR(100) NOT NULL,
				content TEXT NOT NULL,
				created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				expires DATETIME NOT NULL
			);
		`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Error creating table: %v\nQuery: %s", err, query)
		}
	}

}
