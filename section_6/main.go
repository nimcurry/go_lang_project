package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// conention in go. if you build an interface that requires on only one method then it should be named as name+er
type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

//generics

func add[T int | float64 | string](a, b T) T {
	return a + b

}

func main() {

	result := add(1, 2)

	fmt.Println(result)

	printSomething(1)
	printSomething("hello lavdi")
	printSomething(2.30)

	title, content := getNoteData()
	todoText := getTodo()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	newText, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething(newText)

	err = outputData(newText)
	if err != nil {
		return
	}

	outputData(userNote)
	//no error catch needed as program will complete after outputdata(usernote) function

	/* 	err = newText.Save()

	   	if err != nil {
	   		fmt.Println("saving the text failed...")
	   		return
	   	}

	   	fmt.Println("saving the text succeeded") */

}

// special interface which could accept "Any" value type
func printSomething(value interface{}) {
	typedval, ok := value.(int)

	if ok {
		fmt.Println("typedval is int", typedval)
		return
	}
	/* switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println("String: ", value)

	} */
	fmt.Println(value)
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("saving the note failed...")
		return err
	}

	fmt.Println("saving the note succeeded")
	return nil
}

func getTodo() string {
	text := getUserInput("todo text: ")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	//var value string
	//fmt.Scan(&value)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
