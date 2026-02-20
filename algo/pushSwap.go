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
    a, _ := parseArg(args)
    b := &List{}
    set_current_position(a)
    pb(a, b)
    printList(b)
    for size(a) > 3 {
        pb(a, b)
    }
    tinySort(a)
    fmt.Println("\n COUNTE")
        printList(a)
        fmt.Println()
        printList(b)
        fmt.Println("\n ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
    // currb := b.Head
    COUNTER := 0
    for b.Head != nil {
        init_list(a, b)
        move_nodes(a, b)
        fmt.Println("\n COUNTER: ", COUNTER)
        printList(a)
        fmt.Println()
        printList(b)
        fmt.Println("\n ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
        COUNTER++
    }
    set_current_position(a)
    if min(a).above_median {
        fmt.Print(a.Head == min(a))
        for a.Head != min(a) {
            ra(a)
        }
    } else {
        fmt.Print(a.Head != min(a))
        for a.Head != min(a) {
            rra(a)
        }
    }
    fmt.Println()
    printList(a)
}


func move_nodes(target_list, list *List) {
    // currList := list.Head
    // currTList := target_list.Head
    cheapest := get_cheapest(list)

    if cheapest.above_median && cheapest.target_node.above_median {
        rotate_both(target_list, list, cheapest)
    }
    if !(cheapest.above_median) && !(cheapest.target_node.above_median) {
        reverse_rotate_both(target_list, list, cheapest)
    }
    cheapest = get_cheapest(list)
    finish_rotation(target_list, cheapest.target_node, "a")
    finish_rotation(list, cheapest, "b")
    pa(target_list, list)


}

func finish_rotation(l *List, cheapest *Node, a_or_b string) {
    // curr := l.Head
    for l.Head != cheapest {
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

func rotate_both(target_list, list *List, cheapest *Node) {
    
    for list.Head != cheapest && target_list.Head != cheapest.target_node  {
          rr(target_list, list)
    }
    set_current_position(target_list)
    set_current_position(list)
    }

func reverse_rotate_both(target_list, list *List, cheapest *Node) {
    curr := list.Head
    currT := target_list.Head
    for curr != cheapest && currT != cheapest.target_node  {
        rrr(target_list, list)
    }
    set_current_position(target_list)
    set_current_position(list)
}
func init_list(target_list, list *List) {
    set_current_position(target_list)
    set_current_position(list)
    set_target_node(target_list, list)
    set_price(target_list, list)
    set_cheapest(list)    
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
        if curr.index <= med {
            curr.above_median = true
        }else {
            curr.above_median = false
        }
        curr = curr.Next
    }
}

func set_target_node(target_list, list *List) {
    node := list.Head
    target_node := target_list.Head
    target := max(target_list)
    isTarget := false
    // target := target_list.Head
    for node != nil {
        for target_node != nil {
            if target_node.N > node.N && target_node.N < target.N { 
                isTarget = true               
                target = target_node
            }
            target_node = target_node.Next
        }
        if !isTarget {
            target = min(target_list)
        } else{
            node.target_node = target
            
        }
        node = node.Next
    }

}

func set_price(target_list, list *List) {
    sizea := size(target_list)
    sizeb := size(list)
    // curra := target_list.Head
    currb := list.Head
    for currb != nil {
        currb.push_cost = currb.index
        if !(currb.above_median) {
            currb.push_cost = sizeb - currb.index
        }
        if currb.target_node.above_median {
            currb.push_cost += currb.target_node.index
        } else {
            currb.push_cost += sizea - currb.target_node.index
        }
        currb = currb.Next        
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
    }
    return curr
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