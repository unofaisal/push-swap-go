package algo

import "fmt"

// push all nodes to stack B until len of stack = 3, need helper func to check stack len

// if stack len = 3 perform tinySort

// after tiny sort, find target nodes and cost to get cheaper node to push back to a

// loop over stack b and use move nodes to move nodes

// the above loop sorted the list but the smallest digits may not be up so check if sorted

// if not sorted find the smallest node
//   if smallest node above medium use ra
//  if below medium use rra

func Push_swap(args []string) {
	if len(args) == 0 {
		return
	}
	a, err := parseArg(args)
	if err != nil {
		fmt.Println("Error")
		return
	}

	if checkDuplicates(a) {
		fmt.Println("Error")
		return
	}
	if isSorted(a) {
		return
	}

	b := &List{}
	pb(a, b)
	pb(a, b)

	for size(a) > 3 {
		move_nodes(a, b, true)
	}

	tinySort(a)

	for b.Head != nil {
		move_nodes(a, b, false)
	}

	set_current_position(a)
	minNode := min(a)
	for a.Head != minNode {
		set_current_position(a)
		if minNode.above_median {
			ra(a)
		} else {
			rra(a)
		}
	}
}
func size(l *List) int {
	curr := l.Head
	count := 0
	for curr != nil {
		count++
		curr = curr.Next
	}
	return count
}

func max(l *List) *Node {
	curr := l.Head
	max := curr
	for curr != nil {
		if max.N < curr.N {
			max = curr
		}
		curr = curr.Next
	}
	return max
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
