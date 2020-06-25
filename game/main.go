
package main

import "fmt"
import "math/rand"

func main() { 
	numberToDiscover := rand.Intn(10) // return 0 to 101, 0 -> 100

	playGame(numberToDiscover) 
}

func playGame( numberToDiscover int) {
	attemps := 0

	fmt.Println("insert a number please:")

	var insertedNumber int

	for {
		fmt.Scanln(&insertedNumber) // similar to scan but stops scanning at a newline "enter" / from terminal

		if insertedNumber < numberToDiscover {
			fmt.Println("the number is smaller")
			attemps++
		} else if insertedNumber > numberToDiscover{
			fmt.Println("the number is bigger")
			attemps++
		} else if insertedNumber == numberToDiscover {
			fmt.Println("you win!")
			return // do not return anything
		}
	}
}