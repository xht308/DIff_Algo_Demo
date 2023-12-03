package main

import (
	"fmt"
)

func main() {
	src := []string{"a", "b", "c", "d", "e"}
	dest := []string{"a", "c", "e"}
	// operations := getBasicDiffOperations(text1, text2)
	// for _, operation := range operations.Slice() {
	// 	fmt.Println(operation)
	// }
	temp := getMayersDiffOperations(src, dest)
	fmt.Println(temp.Slice())
}