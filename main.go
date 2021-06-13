package main

import (
	"fmt"

	"github.com/yuto51942/go-package-sample/fibonacci"
)

func main() {
	a := fibonacci.FibonacciCal(1)

	result, err := a.Cal(10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
