package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func WalkInternal(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		WalkInternal (t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		WalkInternal (t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	WalkInternal(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for v1 := range c1 {
		v2, success := <- c2
		if v1 != v2 || success == false {
			return false
		}
	}
	
	// check if c2 has extra nodes
	_, success := <- c2
	if success == true {
		return false
	}
	
	return true
}

func main() {
	var b1 = Same(tree.New(1), tree.New(1))
	var b2 = Same(tree.New(1), tree.New(2))
	fmt.Println(b1, b2)
}

