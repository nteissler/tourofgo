package main

import (
	"golang.org/x/tour/tree"
	"fmt"
	)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		_walk(t.Left, ch)
	}
	ch <-t.Value
	if t.Right != nil {
		_walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v1 := range ch1 {
		if v1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	/*ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}*/
	fmt.Println("going")
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println("gone")
}


	// why was the world only printed 4 times and not 5 in first slide thread
