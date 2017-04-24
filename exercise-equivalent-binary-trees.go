package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	chTree1 := make(chan int)
	chTree2 := make(chan int)
	go Walk(t1, chTree1)
	go Walk(t2, chTree2)

	for i := 0; i < 10; i++ {
		if <-chTree1 != <-chTree2 {
			return false
		}
	}
	return true
}

func main() {
	walkTestCh := make(chan int)
	go Walk(tree.New(1), walkTestCh)
	for i := 0; i < 10; i++ {
		fmt.Print(<-walkTestCh, ",")
	}

	fmt.Println()

	fmt.Println("expected output(true) :", Same(tree.New(1), tree.New(1)))
	fmt.Println("expected output(false):", Same(tree.New(1), tree.New(2)))
	fmt.Println("expected output(true) :", Same(tree.New(55), tree.New(55)))

}
