package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(outer())
}

func outer() int {
	time.Sleep(1 * time.Second)
	return inner1() + inner2()
}

func inner1() int {
	time.Sleep(1 * time.Second)
	return 1
}

func inner2() int {
	time.Sleep(1 * time.Second)
	return 1
}
