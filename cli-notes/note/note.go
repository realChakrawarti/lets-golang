package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title    string
	Content  string
	CreateAt time.Time
}

func (n Note) Display() {
	fmt.Printf("Your note titled \"%v\" created at %v has following content:\n\n%v\n", n.Title, n.CreateAt, n.Content)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonData, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid value provided")
	}

	return Note{
		Title:    title,
		Content:  content,
		CreateAt: time.Now(),
	}, nil
}
