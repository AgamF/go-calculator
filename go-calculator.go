package main

import (
	"fmt"
	"strconv"
	"os"
)

var x int;
var y int;

func sum(x int, y int) {
	var result int = x + y
	fmt.Println("Result is " + strconv.Itoa(result))
}

func sub(x int, y int) {
	var result int = x - y
	fmt.Println("Result is " + strconv.Itoa(result))
}

func div(x int, y int) {
	var result int = x / y
	fmt.Println("Result is " + strconv.Itoa(result))
}

func mul(x int, y int) {
	var result int = x * y
	fmt.Println("Result is " + strconv.Itoa(result))
}

func askForInput() {
	fmt.Print("Enter first number: ")
	fmt.Scan(&x)
	fmt.Print("Enter second number: ")
	fmt.Scan(&y)
}

func main() {
	var (
		msg string
		addition string = "+"
		subtraction string = "-"
		division string = "/"
		multiplication string = "*"
		quit string = "quit"
	)
		
	for true {
		fmt.Print("Whats the operation you wish to perform? (+|-|/|*|quit) ")
		fmt.Scan(&msg)

		switch msg {
			case addition:
				askForInput()
				sum(x, y)
			case subtraction:
				askForInput()
				sub(x, y)
			case division:
				askForInput()
				div(x, y)
			case multiplication:
				askForInput()
				mul(x, y)
			case quit:
				os.Exit(0)
			default:
				fmt.Println("Thats not a number!")
		}
	}
}