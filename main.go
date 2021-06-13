package main

import (
	"fmt"

	"github.com/yuto51942/go-package-sample/fibonacci"
)

func main() {
	a := fibonacci.FibonacciCal(1)

	result := a.Cal(10)

	fmt.Println(result)
}
