package main

import (
	"fmt"
	"strings"
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

func (s *opStack) Slice() []operation {
	temp := make([]operation, len(s.operations))
	copy(temp, s.operations)

	// Reverse the operations
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}

	return temp
}