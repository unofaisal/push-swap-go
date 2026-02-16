package main

import (
	"strings"
	"errors"
)
type Node struct {
	n int
	next *Node
	prev *Node
}

func main() {

}


func parseArg(arg []string) (*Node, error) {
	node := &Node
	if len(arg) > 1 {
		return nil, errors.New("ERROR")
	}
	

	lS:= strings.Split(arg[0], " ")

	// if err != nil {
	// 	return nil, errors.New("ERROR")
	// }

	for _, n := range lS {
		node = appendNode(n)
	}
	return node
}


