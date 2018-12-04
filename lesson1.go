package main

import (
	"fmt"
)

func GreetSomeone() (string, int) {
	fmt.Printf("Hello, what is your name?\n>")

	var uName string
	fmt.Scanf("%s\n", &uName)
	var age int
	fmt.Printf("How old are you?\n>")
	fmt.Scanf("%d\n", &age)

	return uName, age
}

func SayGoodbye(name string) {

	fmt.Printf("Goodbye to %s\n", name)

}

func main() {

	name, age := GreetSomeone()

	fmt.Printf("Hello %s, you are %d years old.\n", name, age)

	age = age + 5
	fmt.Printf("In 5 years you will be %d years old\n", age)

	SayGoodbye(name)
}
