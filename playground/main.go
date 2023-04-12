package main

import "fmt"

func main() {

	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 14, 16}

	list = append(list, 11)

	for _, number := range list {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Printf("%v is odd\n", number)
		}
	}

}
