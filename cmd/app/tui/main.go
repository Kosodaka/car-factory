package main

import (
	"car-factory/cmd/app/tui/app"
	"fmt"
	"github.com/c-bata/go-prompt"
)

func main() {

	fmt.Println("Car Factory app")
	fmt.Println("Enter 'help' to get list of commands")
	fmt.Println("Please use ctrl+c to exit the program")
	defer fmt.Println("Goodbye!")
	p := prompt.New(
		app.Executor,
		app.Completer, //Плохо работате в консоле IDE лучше смотреть в терминале системы
		prompt.OptionTitle("interactive tui"),
		prompt.OptionPrefix(">>"),
		prompt.OptionInputTextColor(prompt.Purple),
	)
	p.Run()

}
