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

func CloseWalk(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go CloseWalk(t1, ch1)
	go CloseWalk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		fmt.Println("vals", v1, v2, ok1, ok2)

		if ok1 != ok2 {
			fmt.Println("длина не та")
			return false
		}

		if v1 != v2 {
			fmt.Println("разные значения")
			return false
		}

		if !ok1 && !ok2 {
			fmt.Println("хз че это")
			return true
		}

	}

	//for i := 0; i < 10; i++ {
	//	x1 := <-ch1
	//	x2 := <-ch2
	//
	//	if x1 != x2 {
	//		return false
	//	}
	//}
	//
	//return true

}

//func main() {
//	t1 := tree.New(1)
//	t2 := tree.New(1)
//	t3 := tree.New(2)
//
//	fmt.Println(Same(t1, t2))
//	fmt.Println(Same(t1, t3))
//}
