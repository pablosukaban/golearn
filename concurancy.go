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
	go sum(s[3:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func ChannelsFibs(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func RunChannelsFibs() {
	c := make(chan int, 10)

	go ChannelsFibs(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}

func fibsWithSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func runFibsWithSelect() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibsWithSelect(c, quit)
}

func bombTimer() {
	tick := time.Tick(200 * time.Millisecond)
	boom := time.After(2000 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOOOM")
			return
		default:
			fmt.Println("...................")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
