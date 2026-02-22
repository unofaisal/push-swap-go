package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	a, err := parseArg(os.Args[1:])
	if err != nil {
		printError()
		return
	}
	b := &List{}

	err = readInstructions(a, b)
	if err != nil {
		printError()
		return
	}

	if isSorted(a) && b.Head == nil {
		printOK()
	} else {
		printKO()
	}

}

func printError() {
	println("Error")
}

func printOK() {
	println("OK")
}

func printKO() {
	println("KO")
}

func isSorted(l *List) bool {
	curr := l.Head
	for curr != nil && curr.Next != nil {
		if curr.N > curr.Next.N {
			return false
		}
		curr = curr.Next
	}
	return true
}

func readInstructions(a, b *List) error {
	var instruction string
	for {
		_, err := fmt.Scanln(&instruction)
		if err != nil {
			break
		}
		switch instruction {
		case "sa":
			sa(a)
		case "sb":
			sb(b)
		case "ss":
			ss(a, b)
		case "pa":
			pa(a, b)
		case "pb":
			pb(a, b)
		case "ra":
			ra(a)
		case "rb":
			rb(b)
		case "rr":
			rr(a, b)
		case "rra":
			rra(a)
		case "rrb":
			rrb(b)
		case "rrr":
			rrr(a, b)
		default:
			return errors.New("ERROR")
		}
	}
	return nil
}

func pb(a, b *List) {

	if a.Head == nil {
		return
	}
	node := a.Head
	a.Head = node.Next
	if a.Head != nil {
		a.Head.Prev = nil
	} else {
		a.Tail = nil
	}

	node.Next = b.Head
	node.Prev = nil
	if b.Head != nil {
		b.Head.Prev = node
	} else {
		b.Tail = node
	}
	b.Head = node
}

func pa(a, b *List) {

	if b.Head == nil {
		return
	}
	node := b.Head
	b.Head = node.Next
	if b.Head != nil {
		b.Head.Prev = nil
	} else {
		b.Tail = nil
	}

	node.Next = a.Head
	node.Prev = nil
	if a.Head != nil {
		a.Head.Prev = node
	} else {
		a.Tail = node
	}
	a.Head = node
}

func swap(a *List) {
	if a.Head == nil || a.Head.Next == nil {
		return
	}
	curr := a.Head
	next := curr.Next

	curr.Next = next.Next
	if next.Next != nil {
		next.Next.Prev = curr
	} else {
		a.Tail = curr
	}
	next.Prev = nil
	next.Next = curr
	curr.Prev = next
	a.Head = next
}

func sa(a *List) {
	swap(a)
}

func sb(b *List) {
	swap(b)
}

func ss(a, b *List) {
	swap(a)
	swap(b)
}

func rotate(a *List) {
	curr := a.Head
	currT := a.Tail

	a.Head = curr.Next
	a.Head.Prev = nil
	curr.Next = nil
	curr.Prev = currT
	a.Tail.Next = curr
	a.Tail = curr
	a.Tail.Next = nil
}
func ra(a *List) {
	rotate(a)
}

func rb(b *List) {
	rotate(b)
}

func rr(a, b *List) {
	rotate(a)
	rotate(b)
}

func reverse_rotate(a *List) {
	if a.Head == nil || a.Head.Next == nil {
		return
	}
	currH := a.Head
	currT := a.Tail
	a.Tail = currT.Prev
	a.Tail.Next = nil
	a.Head.Prev = currT
	a.Head = currT
	a.Head.Next = currH
	a.Head.Prev = nil
}

func rra(a *List) {
	reverse_rotate(a)

}

func rrb(b *List) {
	reverse_rotate(b)
}

func rrr(a, b *List) {
	reverse_rotate(a)
	reverse_rotate(b)
}

func printList(l *List) {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%v ->", curr.N)
		curr = curr.Next
	}
}

func parseArg(arg []string) (*List, error) {
	l := &List{}
	if len(arg) > 1 {
		return nil, errors.New("ERROR")
	}

	lS := strings.Split(arg[0], " ")
	for _, n := range lS {
		x, err := strconv.Atoi(n)
		if err != nil {
			return nil, errors.New("ERROR")
		}
		appendNode(x, l)
	}
	return l, nil
}

func appendNode(n int, l *List) *List {
	new := &Node{N: n}
	curr := l.Head
	if curr == nil {
		l.Head = new
		l.Tail = new
		new.Next = nil
		new.Prev = nil
		return l
	}
	prev := l.Tail
	l.Tail.Next = new

	l.Tail = new
	l.Tail.Prev = prev
	return l
}

type Node struct {
	N            int
	push_cost    int
	cheapest     bool
	above_median bool
	target_node  *Node
	index        int
	Next         *Node
	Prev         *Node
}

type List struct {
	Head *Node
	Tail *Node
}
