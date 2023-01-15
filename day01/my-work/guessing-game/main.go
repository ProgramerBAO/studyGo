package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100 // 生成随机数
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
	var guess int
	for {
		fmt.Print("Please input your guess ")
		_, err := fmt.Scanf("%d\n", &guess)
		if err != nil {
			fmt.Println("Please enter an integer less than one hundred.")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
