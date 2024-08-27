package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
}

// to create a interface, we use the type keyword
func main() {
	title, content := getNoteData()
	todoContent := getUserInput("Todo content: ")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userTodo, err := todo.New(todoContent)

	if err != nil {
		fmt.Println(err)
		return
	}

	userTodo.Display()
	err = userTodo.Save()

	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}

	fmt.Println("Saving the todo succeeded!")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
