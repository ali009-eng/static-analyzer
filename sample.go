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

func checky() {
	result1, err := doSomething()
	fmt.Println(result1, err)

	if result1 > 0 {
		fmt.Println("positive")
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	result2, err := doSomethingElse()
	fmt.Println(result2, err)
}
