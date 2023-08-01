package main

import (
	"Soccer-Penalty-Kick-ML-Threading/testing"
	"fmt"
)

func main() {
	/**
	 * Gets user input for which test to run
	 * 1. Run XOR test for simple neural network model
	 * 2. Run Environment Space test for random steps in soccer game env
	 * 0. Exit program
	 */
	input := ""
	for input != "0" {
		fmt.Printf("\nChoose which test to run:\n1. Simple XOR NN Model\n2. Test Environment Space for Soccer Game\n0. Exit program\n\nChoice: ")
		fmt.Scanln(&input)
		if input == "1" {
			testing.RunXorTest()
		} else if input == "2" {
			testing.RunEnvSpaceTest()
		} else if input == "0" {
			fmt.Printf("\nExiting program...\n")
			break
		} else {
			fmt.Printf("\nInvalid selection...\n\n")
		}
	}
}
