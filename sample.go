package main

import "fmt"

func check() {
	result1, err := doSomething()
	fmt.Println(result1, err)

	result2, err := doSomethingElse()
	fmt.Println(result2, err)
}

func doSomething() (int, error) {
	return 1, nil
}

func doSomethingElse() (int, error) {
	return 2, nil
}
