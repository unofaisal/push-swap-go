package algo

// args: stacka, stackb, targetNode, cheapestNode
// loop over doing rr untill one of the cheapestNode or targetNode is at the top of their respective node(needs helper func)
// after above, one list has target or cheap node above so we do a finish swap to both coz of uncertainty(need helper func)

// func move(a, b *List) {

// }

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
