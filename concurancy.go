package main

import (
	"fmt"
	"time"
)

func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	res := 0
	for _, v := range s {
		res += v
	}

	c <- res

}

func Chanel1() {
	s := []int{1, 2, 3, 4, 5, 6}

	c := make(chan int)

	go sum(s[:3], c)
	fmt.Println(c)
	go sum(s[3:], c)
	fmt.Println(c)
	x, y := <-c, <-c
	fmt.Println(x, y)
}
