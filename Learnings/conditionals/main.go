package main

import "fmt"

func if_example(num int) {
	if num == 7 {
		fmt.Println("its 7, Thala for a reason 🚁")
	}
}

func if_else_example(age int) {
	if age < 18 {
		fmt.Println("Enjoy Life while you can 😉 ") 
	} else {
		fmt.Println("Damn you're an adult now 😆")
	}
}

func if_elif_example(grade string) {
	if grade == "A" {
		fmt.Println("you are topper it seems 🤔")
	} else if grade == "B" {
		fmt.Println("top second it seems 😢")
	} else {
		fmt.Println("ever tried, ever failed? try better fail better 😅")
	}
}


func main() {
	if_example(7)
	fmt.Println("_____")

	if_else_example(17)
	fmt.Println("_____")

	if_else_example(20)
	fmt.Println("_____")

	if_elif_example("A")
	fmt.Println("_____")

	if_elif_example("B")
	fmt.Println("_____")

	if_elif_example("C")
	fmt.Println("_____")

}