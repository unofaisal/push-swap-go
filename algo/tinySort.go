package algo

// need a hellper func to find higest digit in stack

// make sure big is at bottom

// swap first two if not sorted

func tinySort(l *List) {
	
	max_node := max(l)
	curr := l.Head
	if curr.N == max_node.N {
		ra(l)
	} else if curr.Next.N == max_node.N {
		rra(l)
	}

	if l.Head.N > l.Head.Next.N {
		sa(l)
	}
	printList(l)
}