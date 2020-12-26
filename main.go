package main

import (
	"fmt"
)

func filter(num *int, result *string) {
	if *num%10 == 0 {
		*result = fmt.Sprintf("%d is a common multiple of  2 & 5", *num)
	} else if *num%2 == 0 {
		*result = fmt.Sprintf("%d is a common multiple of  2", *num)
	} else if *num%5 == 0 {
		*result = fmt.Sprintf("%d is a common multiple of  5", *num)
	} else if len(*result) == 0 {
		panic("error")
	}
}
func input(num *int, result *string) {
	defer func() {
		if e := recover(); e == "error" {
			fmt.Println(fmt.Sprintf("%d is not a common multiple of 2 & 5", *num))
		}
	}()
	filter(num, result)
	fmt.Println(*result)
}
func main() {
	var i *int
	i = new(int)

	var result *string
	result = new(string)

	for *i = 1; *i <= 100; *i++ {
		*result = ""
		input(i, result)
	}
}
