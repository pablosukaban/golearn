package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

//type Tree struct {
//	Left  *Tree
//	Value int
//	Right *Tree
//}

func Walk(t *tree.Tree, ch chan int) {

	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)

}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		x1 := <-ch1
		x2 := <-ch2

		if x1 != x2 {
			return false
		}
	}

	return true

}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(2)
	fmt.Println(Same(t1, t2))
	fmt.Println(Same(t1, t3))
}
