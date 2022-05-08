package main

import (
	"fmt"
	"os"
	"strconv"
)

func greet(name, year string) {
	fmt.Println("Hello! My name is " + name + ".")
	fmt.Println("I was created in " + year + ".")
}

func showName() {
	var name string
	fmt.Println("Please, remind me your name.")
	fmt.Scan(&name)
	fmt.Println("What a great name you have, " + name + "!")
}

func guessAge() {
	var rem3, rem5, rem7, age int

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")
	fmt.Scan(&rem3, &rem5, &rem7)

	age = (rem3*70 + rem5*21 + rem7*15) % 105
	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}

func count() {
	var n int

	fmt.Println("Now I will prove to you that I can count to any number you want.")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		fmt.Printf("%d!\n", i)
	}
}

func startQuiz() {
	fmt.Println("Let's test your programming knowledge.")
	// write your code here
	question1 := "Why do we use methods?"
	answer1 := "1. To repeat a statement multiple times."
	answer2 := "2. To decompose a program into several small subroutines."
	answer3 := "3. To determine the execution time of a program."
	answer4 := "4. To interrupt the execution of a program."
	question1CorrectAnswer := 2
	fmt.Printf("%s\n%s\n%s\n%s\n%s\n", question1, answer1, answer2, answer3, answer4)
	var userAnswer int
	fmt.Scan(&userAnswer)
	for question1CorrectAnswer == 2 {
		if userAnswer == question1CorrectAnswer {
			fmt.Println("Completed, have a nice day!")
			sayGoodbye()
		} else {
			fmt.Println("Please, try again.")
			fmt.Scan(&userAnswer)
		}
	}

}

func sayGoodbye() {
	fmt.Println("Congratulations, have a nice day!")
	os.Exit(0)
}

func main() {
	greet("Aid", "2020") // change it as you need
	showName()
	guessAge()
	count()
	startQuiz()
	//sayGoodbye()
}
