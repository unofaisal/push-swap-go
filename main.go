package main

import (	
	"os"
	"push/algo"
	"fmt"
	"strings"
	"strconv"
	"errors"

)
type Node struct {
	N int
    push_cost int
    cheapest bool
    above_median bool
    target_node *Node
    index int
	Next *Node
	Prev *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func main() {
	args := os.Args[1:]

	algo.Push_swap(args)
	// a, _ := parseArg(args)

	// b := &List{}

	// fmt.Println("\naaaaaaaaaa1")
	// printList(a)
	// fmt.Println()
	// pb(a, b)
	// pb(a, b)
	// pb(a, b)
	// printList(a)
	// fmt.Println("\nbbbbbbbb1")
	// printList(b)
	// pb(a, b)
	// pb(a, b)
	// fmt.Println("\naaaaaaaaaa2")
	// printList(a)
	// fmt.Println("\nbbbbbbbb1")
	// printList(b)
	// pa(a, b)
	// pa(a, b)
	// pa(a, b)
	// fmt.Println("\naaaaaaaaaa3")
	// printList(a)
	// fmt.Println("\nbbbbb2")
	// printList(b)
	
}

func min(l *List) *Node {
    curr := l.Head
    min := l.Head
    for curr != nil {
        if min.N > curr.N {
            min = curr
        }
        curr = curr.Next
    }
    return min
}









func pb(a, b *List) {
	curra := a.Head
	currb := b.Head


	a.Head = curra.Next
	a.Head.Prev = nil

	

	if currb == nil {
		b.Head = curra
		b.Tail = curra
		b.Head.Prev = nil
		b.Head.Next = nil
	} else {
		curra.Next = currb
		currb.Prev = curra
		b.Head = curra
		b.Head.Prev = nil
	}

	// a0Data := curra.N
	// b0NNode := &Node{N:a0Data, Next:currb}
	
	// b.Head = b0NNode
	// if currb == nil {
	// 	a.Head = curra.Next
	// 	b.Tail = b0NNode
	// 	b.Head.Next = nil
	// 	b.Head.Prev = nil
	// } else {
	// 	a.Head = curra.Next
	// 	currb.Prev = b.Head
	// 	b.Head.Next = currb
	// 	b.Head.Prev = nil
	// }
}

func pa(a, b *List) {
	curra := a.Head
	currb := b.Head

	b.Head = currb.Next
	
	
	
	if curra == nil {
		a.Head = currb
		a.Tail = currb
		a.Head.Prev = nil
	} else {
		a.Head = currb
		curra.Prev = currb
		a.Head.Next = curra
		a.Head.Prev = nil
	}

	// b0Data := currb.N
	// a0NNode := &Node{N:b0Data, Next:curra}
	
	// a.Head = a0NNode
	// if curra == nil {
	// 	b.Head = currb.Next
	// 	a.Tail = a0NNode
	// 	a.Head.Next = nil
	// 	a.Head.Prev = nil
	// } else {
	// 	b.Head = currb.Next
	// 	curra.Prev = a.Head
	// 	a.Head.Next = curra
	// 	a.Head.Prev = nil
	// }
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

	a.Head = curr.Next
	a.Head.Prev = nil
	curr.Next = nil
	curr.Prev = currT
	a.Tail.Next = curr
	a.Tail = curr
	a.Tail.Next = nil

	// a0Data := curr.N
	// a.Head = curr.Next
	
	// aTNode := &Node{N: a0Data, Next: nil, Prev: currT}
	// a.Tail.Next = aTNode
	// a.Tail = aTNode
	// a.Tail.Next = nil
	
}

func rb(b *List) {

	curr := b.Head
	currT := b.Tail

	b.Head = curr.Next
	b.Head.Prev = nil
	curr.Next = nil
	curr.Prev = currT
	b.Tail.Next = curr
	b.Tail = curr
	b.Tail.Next = nil

	// curr := b.Head
	// currT := b.Tail
	// b0Data := curr.N
	// b.Head = curr.Next
	
	// bTNode := &Node{N: b0Data, Next: nil, Prev: currT}
	// b.Tail.Next = bTNode
	// b.Tail = bTNode
	// b.Tail.Next = nil
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




// func pb(a, b *List) {
// 	curra := a.Head
// 	currb := b.Head

// 	a0Data := curra.N
// 	b0NNode := &Node{N:a0Data, Next:currb}
	
// 	b.Head = b0NNode
// 	if currb == nil {
// 		a.Head = curra.Next
// 		b.Tail = b0NNode
// 		b.Head.Next = nil
// 		b.Head.Prev = nil
// 	} else {
// 		a.Head = curra.Next
// 		currb.Prev = b.Head
// 		b.Head.Next = currb
// 		b.Head.Prev = nil
// 	}
// }

// func pa(a, b *List) {
// 	curra := a.Head
// 	currb := b.Head

// 	b0Data := currb.N
// 	a0NNode := &Node{N:b0Data, Next:curra}
	
// 	a.Head = a0NNode
// 	if curra == nil {
// 		b.Head = currb.Next
// 		a.Tail = a0NNode
// 		a.Head.Next = nil
// 		a.Head.Prev = nil
// 	} else {
// 		b.Head = currb.Next
// 		curra.Prev = a.Head
// 		a.Head.Next = curra
// 		a.Head.Prev = nil
// 	}
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
// 		curr = curr.Next
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


