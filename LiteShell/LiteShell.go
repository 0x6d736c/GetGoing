package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin) // Set-up scanner for read-in from Stdin

	//Infinite loop
	for {
		fmt.Print("> ") //Display prompt
		input.Scan()    //Receive input

		inputString := input.Text() //Receive the input text from the scan
		executeCommand(inputString)
	}
}

//Executes the argument passed-in via the shell
func executeCommand(inputString string) {
	commands := strings.Split(inputString, " ") //Split the input on whitespace
	command, args := commands[0], commands[1:]  //Set command to the executable, args to the args to the executable

	commandExecution := exec.Command(command, args...) //Execute the command with relevant arguments

	result, err := commandExecution.Output() //Retrieve the relevant outputs from execution

	//Handle error, reload shell
	if err != nil {
		fmt.Print("Invalid command: ")
		for _, str := range commands {
			fmt.Print(str, " ")
		}
		fmt.Println()
	}

	//Output the result of the execution
	fmt.Print(string(result))
}
