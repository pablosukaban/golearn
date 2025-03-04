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

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	if t == nil {
		return
	}

	Walk(t.Left, ch)
	//fmt.Println("t val: ", t.Value)
	ch <- t.Value
	Walk(t.Right, ch)

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	t := tree.New(1)
	c := make(chan int)
	fmt.Println(t)

	go Walk(t, c)

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}
