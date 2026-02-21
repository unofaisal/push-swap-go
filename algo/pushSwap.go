package algo

// push all nodes to stack B until len of stack = 3, need helper func to check stack len

// if stack len = 3 perform tinySort

// after tiny sort, find target nodes and cost to get cheaper node to push back to a

// loop over stack b and use move nodes to move nodes

// the above loop sorted the list but the smallest digits may not be up so check if sorted

// if not sorted find the smallest node
//   if smallest node above medium use ra
//  if below medium use rra

func Push_swap(args []string) {
	a, _ := parseArg(args)
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

	printList(a)
}

func move_nodes(a, b *List, a_to_b bool) {
	init_list(a, b, a_to_b)

	var source, target *List
	var push func(*List, *List)
	var sourceLabel, targetLabel string

	if a_to_b {
		source, target = a, b
		push = pb
		sourceLabel, targetLabel = "a", "b"
	} else {
		source, target = b, a
		push = pa
		sourceLabel, targetLabel = "b", "a"
	}

	cheapest := get_cheapest(source)

	if cheapest.above_median && cheapest.target_node.above_median {
		rotate_both(a, b, cheapest, a_to_b)
	}
	if !(cheapest.above_median) && !(cheapest.target_node.above_median) {
		reverse_rotate_both(a, b, cheapest, a_to_b)
	}

	cheapest = get_cheapest(source)
	finish_rotation(source, cheapest, sourceLabel)
	finish_rotation(target, cheapest.target_node, targetLabel)
	push(a, b)
}

func finish_rotation(l *List, cheapest *Node, a_or_b string) {
	// curr := l.Head
	for l.Head != cheapest {
		set_current_position(l)
		if a_or_b == "a" {
			if cheapest.above_median {
				ra(l)
			} else {
				rra(l)
			}
		} else if a_or_b == "b" {
			if cheapest.above_median {
				rb(l)
			} else {
				rrb(l)
			}
		}
	}
}

func rotate_both(a, b *List, cheapest *Node, a_to_b bool) {
	var source, target *List
	if a_to_b {
		source, target = a, b
	} else {
		source, target = b, a
	}
	for source.Head != cheapest && target.Head != cheapest.target_node {
		rr(a, b)
	}
	set_current_position(a)
	set_current_position(b)
}

func reverse_rotate_both(a, b *List, cheapest *Node, a_to_b bool) {
	var source, target *List
	if a_to_b {
		source, target = a, b
	} else {
		source, target = b, a
	}
	for source.Head != cheapest && target.Head != cheapest.target_node {
		rrr(a, b)
	}
	set_current_position(a)
	set_current_position(b)
}
func init_list(a, b *List, a_to_b bool) {
	set_current_position(a)
	set_current_position(b)
	set_target_node(a, b, a_to_b)
	set_price(a, b, a_to_b)
	if a_to_b {
		set_cheapest(a)
	} else {
		set_cheapest(b)
	}
}

func set_current_position(l *List) {
	curr := l.Head
	// prev := &Node{}
	index := 0
	for curr != nil {
		curr.index = index
		index++
		curr = curr.Next
	}
	med := size(l) / 2
	curr = l.Head
	for curr != nil {
		if curr.index < med {
			curr.above_median = true
		} else {
			curr.above_median = false
		}
		curr = curr.Next
	}
}

func set_target_node(a, b *List, a_to_b bool) {
	if a_to_b {
		// find target in b for each a node (descending order in b)
		nodeA := a.Head
		for nodeA != nil {
			var bestTarget *Node
			minDiff := 2147483647
			nodeB := b.Head
			for nodeB != nil {
				diff := nodeA.N - nodeB.N
				if diff > 0 && diff < minDiff {
					minDiff = diff
					bestTarget = nodeB
				}
				nodeB = nodeB.Next
			}
			if bestTarget == nil {
				bestTarget = max(b)
			}
			nodeA.target_node = bestTarget
			nodeA = nodeA.Next
		}
	} else {
		// find target in a for each b node (ascending order in a)
		nodeB := b.Head
		for nodeB != nil {
			var bestTarget *Node
			minDiff := 2147483647
			nodeA := a.Head
			for nodeA != nil {
				diff := nodeA.N - nodeB.N
				if diff > 0 && diff < minDiff {
					minDiff = diff
					bestTarget = nodeA
				}
				nodeA = nodeA.Next
			}
			if bestTarget == nil {
				bestTarget = min(a)
			}
			nodeB.target_node = bestTarget
			nodeB = nodeB.Next
		}
	}
}
func set_price(a, b *List, a_to_b bool) {
	var source, target *List
	if a_to_b {
		source, target = a, b
	} else {
		source, target = b, a
	}

	sizeSource := size(source)
	sizeTarget := size(target)
	curr := source.Head

	for curr != nil {
		var costSource int
		if curr.above_median {
			costSource = curr.index
		} else {
			costSource = sizeSource - curr.index
		}

		var costTarget int
		if curr.target_node.above_median {
			costTarget = curr.target_node.index
		} else {
			costTarget = sizeTarget - curr.target_node.index
		}

		if curr.above_median == curr.target_node.above_median {
			if costSource > costTarget {
				curr.push_cost = costSource
			} else {
				curr.push_cost = costTarget
			}
		} else {
			curr.push_cost = costSource + costTarget
		}

		curr = curr.Next
	}
}

func set_cheapest(l *List) {
	curr := l.Head
	cheap := l.Head
	for curr != nil {
		if curr.push_cost < cheap.push_cost {
			cheap = curr
		}
		curr.cheapest = false
		curr = curr.Next
	}
	cheap.cheapest = true
}

func get_cheapest(l *List) *Node {
	curr := l.Head
	for curr != nil {
		if curr.cheapest {
			return curr
		}
		curr = curr.Next
	}
	return nil
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
