
package main

import "fmt"

func main() { 
	// if I want to return more than one value
	// var str1 , str2 string
	// str1, str2 = createSayHi("Carla")

	// const str3 // const is available
	greetings := createSayHi("Carla"); // get the value that return this function and save it in a variable called greetings ":=" create and save variable

	fmt.Println(greetings)
}

func createSayHi (name string) string { //the second string is the type of value returned // I can return more than one value (string, string)
	return fmt.Sprintf("hi! %s",name);
}