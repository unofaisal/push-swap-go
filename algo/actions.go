package algo

import (
	"fmt"
	"errors"
	"strconv"
	"strings"
)

func pb(a, b *List) {
	curra := a.Head
	currb := b.Head

	a0Data := curra.N
	b0NNode := &Node{N:a0Data, Next:currb}
	
	b.Head = b0NNode
	if currb == nil {
		a.Head = curra.Next
		b.Tail = b0NNode
		b.Head.Next = nil
		b.Head.Prev = nil
	} else {
		a.Head = curra.Next
		currb.Prev = b.Head
		b.Head.Next = currb
		b.Head.Prev = nil
	}
}

func pa(a, b *List) {
	curra := a.Head
	currb := b.Head

	b0Data := currb.N
	a0NNode := &Node{N:b0Data, Next:curra}
	
	a.Head = a0NNode
	if curra == nil {
		b.Head = currb.Next
		a.Tail = a0NNode
		a.Head.Next = nil
		a.Head.Prev = nil
	} else {
		b.Head = currb.Next
		curra.Prev = a.Head
		a.Head.Next = curra
		a.Head.Prev = nil
	}
}

func sa(a *List) {
	curr := a.Head
	next := curr.Next
	if curr == nil || curr.Next == nil {
		return
	}
	curr.Next = next.Next
	next.Next = curr
	a.Head = next
	
}

func sb(b *List) {
	curr := b.Head
	next := curr.Next
	if curr == nil || curr.Next == nil {
		return
	}
	curr.Next = next.Next
	next.Next = curr
	b.Head = next
	
}

func ss(a, b *List) {
	sa(a)
	sb(b)
}

func ra(a *List) {
	curr := a.Head
	currT := a.Tail
	a0Data := curr.N
	a.Head = curr.Next
	
	aTNode := &Node{N: a0Data, Next: nil, Prev: currT}
	a.Tail.Next = aTNode
	a.Tail = aTNode
	a.Tail.Next = nil
	
}

func rb(b *List) {
	curr := b.Head
	currT := b.Tail
	b0Data := curr.N
	b.Head = curr.Next
	
	bTNode := &Node{N: b0Data, Next: nil, Prev: currT}
	b.Tail.Next = bTNode
	b.Tail = bTNode
	b.Tail.Next = nil
}

func rr(a, b *List) {
	ra(a)
	rb(b)
}

func rra(a *List) {
	currH := a.Head
	currT := a.Tail
	// aTData := currT.N
	// aTNNode := &Node{N:aTData, Next:currH}
	// a.Head = aTNNode
	// fmt.Println("hellloooo: ", currT.Prev)
	// fmt.Println("hellloooo: ", currT)
	// a.Tail = currT.Prev
	// a.Tail.Next = nil

	a.Tail = currT.Prev
	a.Tail.Next = nil
	a.Head.Prev = currT
	a.Head = currT
	a.Head.Next = currH
	a.Head.Prev = nil
	
}

func rrb(b *List) {
	currH := b.Head
	currT := b.Tail
	// bTData := currT.N
	// bTNNode := &Node{N:bTData, Next:currH}
	
	// // fmt.Println("hellloooo: ", currT.Prev)
	// b.Tail = currT.Prev
	// b.Tail.Next = nil
	// b.Head = bTNNode

	b.Tail = currT.Prev
	b.Tail.Next = nil
	b.Head.Prev = currT
	b.Head = currT
	b.Head.Next = currH
	b.Head.Prev = nil
}

func rrr(a, b *List) {
	rra(a)
	rrb(b)
}

func printList(l *List) {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%v [%v] %v ->", curr.N, curr.index, curr.above_median)
		curr = curr.Next
	}
}


func parseArg(arg []string) (*List, error) {
	l := &List{}
	if len(arg) > 1 {
		return nil, errors.New("ERROR")
	}
	

	lS:= strings.Split(arg[0], " ")

	// if err != nil {
	// 	return nil, errors.New("ERROR")
	// }

	for _, n := range lS {
		x, _ := strconv.Atoi(n)
		appendNode(x, l)
	}
	return l, nil
}


func appendNode(n int, l *List) *List  {
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


