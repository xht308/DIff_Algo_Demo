package main

import (
	"fmt"
)

func main() {
	src := []string{"c", "a", "b"}
	dest := []string{"a"}
	// operations := getBasicDiffOperations(text1, text2)
	// for _, operation := range operations.Slice() {
	// 	fmt.Println(operation)
	// }
	// headX, headY, tailX, tailY := findMiddleSnake(src, dest, 0, 0)
	// temp := getMayersDiffOperations(src, dest)
	fmt.Println(getLinearMayersDiffOperations(src, dest))
}