package main

import (
	"bufio"
	"fmt"
	"os"
	"realChakrawarti/note/note"
	"realChakrawarti/note/todo"
	"strings"
)

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

	note.Display()
	err = note.Save()

	if err != nil {
		fmt.Println("Failed to save the note on disk")
		return
	}

	fmt.Println("Note saved on the disk successfully")
	
	todo.Display()

	err = todo.Save()

	if err != nil {
		fmt.Println("Failed to save the todo on disk")
		return
	}

	fmt.Println("Todo saved on the disk successfully")

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
