package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk 遍历树 t，并树中所有的值发送到信道 ch。
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same 判断 t1 和 t2 是否包含相同的值。
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	Walk(t1, c1)
	Walk(t1, c2)
	return c1 == c2
}

func main() {
	// fmt.Print(Same(tree.New(1), tree.New(1)))
	// fmt.Print(Same(tree.New(1), tree.New(2)))
	fmt.Println(tree.New(1).Left.Left.Left)
}
