package main

import (	
	"os"
	"push/algo"

)
// type Node struct {
// 	N int
// 	Next *Node
// 	Prev *Node
// }

// type List struct {
// 	Head *Node
// 	Tail *Node
// }

func main() {
	args := os.Args[1:]

	algo.Push_swap(args)
	// a, _ := parseArg(args)
	// // b := &List{}
	// // fmt.Print(IsSorted(a))
	// rra(a)
	// printList(a)
	// fmt.Println()
	// printList(b)
	
}

// func pb(a, b *List) {
// 	curra := a.Head
// 	currb := b.Head

// 	a0Data := curra.N
// 	b0NNode := &Node{N:a0Data, Next:currb}
// 	b.Head = b0NNode
// 	if currb == nil {
// 		b.Tail = b0NNode
// 	}

// 	a.Head = curra.Next

// }

// func pa(a, b *List) {
// 	curra := a.Head
// 	currb := b.Head

// 	b0Data := currb.N
// 	a0NNode := &Node{N:b0Data, Next:currb}
// 	a.Head = a0NNode
// 	if curra == nil {
// 		a.Tail = a0NNode
// 	}
// 	b.Head = currb.Next
// }

// func sa(a *List) {
// 	curr := a.Head
// 	next := curr.Next
// 	if curr == nil || curr.Next == nil {
// 		return
// 	}
// 	curr.Next = next.Next
// 	next.Next = curr
// 	a.Head = next
	
// }

// func sb(b *List) {
// 	curr := b.Head
// 	next := curr.Next
// 	if curr == nil || curr.Next == nil {
// 		return
// 	}
// 	curr.Next = next.Next
// 	next.Next = curr
// 	b.Head = next
	
// }

// func ss(a, b *List) {
// 	sa(a)
// 	sb(b)
// }

// func ra(a *List) {
// 	curr := a.Head
// 	a.Tail.Next = curr
// 	a.Tail = curr
// 	a.Head = curr.Next
// 	curr.Next = nil
// }

// func rb(b *List) {
// 	curr := b.Head
// 	b.Tail.Next = curr
// 	b.Tail = curr
// 	b.Head = curr.Next
// 	b.Tail.Next = nil
// }

// func rr(a, b *List) {
// 	ra(a)
// 	rb(b)
// }

// func rra(a *List) {
// 	currH := a.Head
// 	currT := a.Tail
// 	a.Tail = currT.Prev
// 	a.Tail.Next = nil
// 	currT.Next = currH
// 	a.Head = currT
// }

// func rrb(b *List) {
// 	currH := b.Head
// 	currT := b.Tail
// 	b.Tail = currT.Prev
// 	b.Tail.Next = nil
// 	currT.Next = currH
// 	b.Head = currT
// }

// func rrr(a, b *List) {
// 	rra(a)
// 	rrb(b)
// }



// func IsSorted(l *List) bool {
// 	curr := l.Head
// 	for curr != nil {
// 		if curr.N > curr.Next.N {
// 			return false
// 		}
// 	}
// 	return true
// }

// func printList(l *List) {
// 	curr := l.Head
// 	for curr != nil {
// 		fmt.Printf("%v ->", curr.N)
// 		curr = curr.Next
// 	}
// }


// func parseArg(arg []string) (*List, error) {
// 	l := &List{}
// 	if len(arg) > 1 {
// 		return nil, errors.New("ERROR")
// 	}
	

// 	lS:= strings.Split(arg[0], " ")

// 	// if err != nil {
// 	// 	return nil, errors.New("ERROR")
// 	// }

// 	for _, n := range lS {
// 		x, _ := strconv.Atoi(n)
// 		appendNode(x, l)
// 	}
// 	return l, nil
// }


// func appendNode(n int, l *List) *List  {
// 	new := &Node{N: n}
// 	curr := l.Head
// 	if curr == nil {
// 		l.Head = new
// 		l.Tail = new
// 		new.Next = nil
// 		new.Prev = nil
// 		return l
// 	}
// 	prev := l.Tail
// 	l.Tail.Next = new

// 	l.Tail = new
// 	l.Tail.Prev = prev
// 	return l
// }


