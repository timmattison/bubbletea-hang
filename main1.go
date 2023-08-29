package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	myModel := New()
	//myModel := MainModel{}

	Program = tea.NewProgram(&myModel)

	if _, err := Program.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
