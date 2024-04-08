package notes

import (
	"Zettelkasten/internal/metadata"
	"time"
)

type Note struct {
	ID       string
	Title    string
	Content  string
	Date     time.Time
	Metadata metadata.Metadata
}

func New(id, title, content string, metadata metadata.Metadata) Note {
	return Note{
		ID:       id,
		Title:    title,
		Content:  content,
		Date:     time.Now(),
		Metadata: metadata,
	}
}
