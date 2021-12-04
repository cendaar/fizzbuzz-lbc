package main

import (
	"fmt"
	"strconv"
)

func main() {
	fizzbuzz(4, 5, 10000, "hello", "world")
}

func fizzbuzz(int1 int, int2 int, limit int, str1 string, str2 string) {
	var output string

	for i:=1; i<=limit; i++ {
		switch {
		case i % (int1*int2) == 0:
			output += str1+str2
		case i % int1 == 0:
			output += str1
		case i % int2 == 0:
			output += str2
		default:
			output += strconv.Itoa(i)
		}

		if i != limit {
			output += ","
		}
	}

	fmt.Println(output)
}