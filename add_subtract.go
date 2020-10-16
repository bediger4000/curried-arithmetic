package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type fn func(...int) fn

func main() {
	curried := add_subtract(numberFromString(os.Args[1]))
	for _, str := range os.Args[2:] {
		curried = curried(numberFromString(str))
	}
	curried() // show calculated answer
}

func numberFromString(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func subtract_add(n ...int) fn {
	return func(a ...int) fn {
		if len(a) == 0 {
			fmt.Printf("(+) %d\n", n[0])
			return nil
		}
		return add_subtract(n[0] - a[0])
	}
}

func add_subtract(n ...int) fn {
	return func(a ...int) fn {
		if len(a) == 0 {
			fmt.Printf("(-) %d\n", n[0])
			return nil
		}
		return subtract_add(n[0] + a[0])
	}
}
