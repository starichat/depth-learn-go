package main

import "fmt"

func main() {
	m := make(chan int)
	_, ok := <-m
	fmt.Println(ok)
}
