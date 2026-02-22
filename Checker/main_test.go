package main

import (
	"testing"
)

func TestIsSorted(t *testing.T) {
	l := &List{}
	appendNode(1, l)
	appendNode(2, l)
	appendNode(3, l)

	if !isSorted(l) {
		t.Error("Expected list to be sorted")
	}

	l2 := &List{}
	appendNode(3, l2)
	appendNode(2, l2)
	appendNode(1, l2)

	if isSorted(l2) {
		t.Error("Expected list to not be sorted")
	}
}

func TestReadInstructions(t *testing.T) {
	a := &List{}
	appendNode(1, a)
	appendNode(2, a)
	appendNode(3, a)

	b := &List{}

	
	err := readInstructions(a, b)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !isSorted(a) || b.Head != nil {
		t.Error("Expected list A to be sorted and list B to be empty")
	}
}

func TestParseArg(t *testing.T) {
	args := []string{"3 2 1"}
	l, err := parseArg(args)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if l.Head == nil || l.Head.N != 3 || l.Head.Next == nil || l.Head.Next.N != 2 || l.Head.Next.Next == nil || l.Head.Next.Next.N != 1 {
		t.Error("Expected list to contain 3 -> 2 -> 1")
	}

	argsInvalid := []string{"3 two 1"}
	_, err = parseArg(argsInvalid)
	if err == nil {
		t.Error("Expected error for invalid input, got nil")
	}
}


