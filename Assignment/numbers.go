package main

import "fmt"

type numbers []int

func scanner() numbers {

	scannedNumbers := numbers{}

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, numbersOneTen := range numbers {
		if numbersOneTen%2 == 0 {
			fmt.Println(numbersOneTen, "is even!")
		} else {
			fmt.Println(numbersOneTen, "is odd!")
		}
	}
	return scannedNumbers
}
