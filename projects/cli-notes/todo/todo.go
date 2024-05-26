package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content  string `json:"content"`
}

func (t Todo) Display() {
	fmt.Printf("Your todo: %v", t.Content)
}

func (t Todo) Save() error {
	fileName := "todo.json"

	jsonData, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid value provided")
	}

	return Todo{
		Content:  content,
	}, nil
}
