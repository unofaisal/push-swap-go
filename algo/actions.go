package algo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

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
	fmt.Print("pb\n")
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
	fmt.Print("pa\n")
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
	fmt.Print("sa\n")
}

func sb(b *List) {
	swap(b)
	fmt.Print("sb\n")
}

func ss(a, b *List) {
	swap(a)
	swap(b)
	fmt.Print("ss\n")
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
	fmt.Print("ra\n")
}

func rb(b *List) {
	rotate(b)
	fmt.Print("rb\n")
}

func rr(a, b *List) {
	rotate(a)
	rotate(b)
	fmt.Print("rr\n")
}

func reverse_rotate(a *List) {
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
	fmt.Print("rra\n")

}

func rrb(b *List) {
	reverse_rotate(b)
	fmt.Print("rrb\n")
}

func rrr(a, b *List) {
	reverse_rotate(a)
	reverse_rotate(b)
	fmt.Print("rrr\n")
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

	// if err != nil {
	// 	return nil, errors.New("ERROR")
	// }

	for _, n := range lS {
		x, _ := strconv.Atoi(n)
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
