package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

////////////////////////////////////////
/** The utilities for diff algorithms */
////////////////////////////////////////

// Quick func: remvoe spaces on the right
func trimRightSpace(str string) string {
	return strings.TrimRight(str, " \t")
}

// The operation conducted to the string

type opcode bool

const (
	INSERT opcode = true
	DELETE opcode = false
)

type operation struct {
	opcode opcode
	index1  int
	index2  int
}

func (op operation) String() string {
	if op.opcode == INSERT {
		return fmt.Sprintf("INSERT %d %d", op.index1, op.index2)
	} else {
		return fmt.Sprintf("DELETE %d %d", op.index1, op.index2)
	}
}

func (op operation) isInsert() bool {
	return op.opcode == INSERT
}

// A simple implementation of Stack of operations

type opStack struct {
	operations []operation
}

func (s *opStack) push(op operation) {
	s.operations = append(s.operations, op)
}

func (s *opStack) pop() operation {
	op := s.operations[len(s.operations)-1]
	s.operations = s.operations[:len(s.operations)-1]
	return op
}

func (s *opStack) isEmpty() bool {
	return len(s.operations) == 0
}

// Slice & Stack Conversion

func (s *opStack) Slice() []operation {
	temp := make([]operation, len(s.operations))
	copy(temp, s.operations)

	// Reverse the operations
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}

	return temp
}

func getStack(operations []operation) opStack {
	stack := opStack{}
	for i := len(operations) - 1; i >= 0; i-- {
		stack.push(operations[i])
	}
	return stack
}

// Printing Operations

func printOperations(operations []operation) {
	for _, op := range operations {
		fmt.Println(op)
	}
}

const (
	contentLength = 30
	// The pattern for printing the operations
	pattern string = "%-4d%-4d%2s    %-30s      %-30s\n"
)

func cutString(str string, length int) string {
	str = strings.ReplaceAll(trimRightSpace(str), "\t", " ")
	if len(str) > length {
		return str[:length]
	} else {
		return str
	}
}

func printOperationsVerbose(operations []operation, src, dest []string) {
	// Trace the position in the strings
	srcIndex, destIndex := 0, 0
	
	// Traverse the operations
	for _, op := range operations {
		// Determine the sign of the operation
		var sign string
		if op.isInsert() {
			sign = "+"
		} else {
			sign = "-"
		}
		// Print the lines before the operation
		for srcIndex < op.index1 {
			temp := cutString(src[srcIndex], contentLength)
			fmt.Printf(pattern, srcIndex, destIndex, "", temp, temp)
			srcIndex++
			destIndex++
		}
		// Print the operation
		if op.isInsert() {
			fmt.Printf(pattern, srcIndex, destIndex, sign, "", cutString(dest[destIndex], contentLength))
			destIndex++
		} else {
			fmt.Printf(pattern, srcIndex, destIndex, sign, cutString(src[srcIndex], contentLength), "")
			srcIndex++
		}
	}

	// Print the remaining lines
	for srcIndex < len(src) {
		temp := cutString(src[srcIndex], contentLength)
		fmt.Printf(pattern, srcIndex, destIndex, "", temp, temp)
		srcIndex++
		destIndex++
	}
}

var (
	// The color for the insert operation
	insert = color.New(color.FgGreen).Add(color.Bold)
	// The color for the delete operation
	delete = color.New(color.FgRed).Add(color.Bold)
)

func printOperationsFancy(operations []operation, src, dest []string) {
	// Trace the position in the strings
	srcIndex, destIndex := 0, 0
	
	// Traverse the operations
	for _, op := range operations {
		// Determine the sign of the operation
		var sign string
		if op.isInsert() {
			sign = "+"
		} else {
			sign = "-"
		}
		// Print the lines before the operation
		for srcIndex < op.index1 {
			temp := cutString(src[srcIndex], contentLength)
			fmt.Printf(pattern, srcIndex, destIndex, "", temp, temp)
			srcIndex++
			destIndex++
		}
		// Print the operation
		if op.isInsert() {
			insert.Printf(pattern, srcIndex, destIndex, sign, "", cutString(dest[destIndex], contentLength))
			destIndex++
		} else {
			delete.Printf(pattern, srcIndex, destIndex, sign, cutString(src[srcIndex], contentLength), "")
			srcIndex++
		}
	}

	// Print the remaining lines
	for srcIndex < len(src) {
		temp := cutString(src[srcIndex], contentLength)
		fmt.Printf(pattern, srcIndex, destIndex, "", temp, temp)
		srcIndex++
		destIndex++
	}
}

// File operations

func readFileLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file line by line
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}