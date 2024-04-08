package main

import (
	"bufio"
	"fmt"
	"os"
	"realChakrawarti/note/note"
	"realChakrawarti/note/todo"
	"strings"
	"time"
)

type saver interface {
	Save() error
}

// Embedded Interface
type outputtable interface {
	saver
	Display()
}

func printAnything(value interface{}) {
	fmt.Println(value)
}

func printAny(value any) {
	fmt.Println(value)
}

func main() {
	title, content := getPostData()
	todoText := takeInput("Enter your todo: ")

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = output(note)

	if err != nil {
		return
	}

	err = output(todo)

	if err != nil {
		return
	}

	printAnything("Thank you for using notes-cli")
	currentDataTime := time.Now().Format(time.RFC850)
	printAny(currentDataTime)

}

func output(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Failed to save on disk")
		return err
	}

	fmt.Println("Saved on the disk successfully")
	return nil
}

func getPostData() (string, string) {
	postTitle := takeInput("Enter title: ")

	postContent := takeInput("Enter content: ")

	return postTitle, postContent
}

func takeInput(label string) string {
	fmt.Print(label)
	readText := bufio.NewReader(os.Stdin)

	inputValue, err := readText.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	inputValue = strings.TrimSuffix(inputValue, "\n")
	inputValue = strings.TrimSuffix(inputValue, "\r")

	return inputValue

}
