package algo

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