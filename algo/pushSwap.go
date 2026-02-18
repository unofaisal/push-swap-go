package algo
// import "fmt"

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
    
    for size(a) > 3 {
        pb(a, b)
    }
    tinySort(a)
    
}

func init_list(a, b *List) {
    set_current_position(a)
    set_current_position(b)
    set_target_node(a, b)
    set_price(a, b)
    set_cheapest(b)
    
}

func set_current_position(l *List) {
    curr := l.Head
    index := 0
    for curr != nil {
        curr.index = index
        index++
        curr = curr.Next
    }
}

func set_target_node(list, target_list *List) {
    node := list.Head
    target_node := target_list.Head
    for node != nil {
        for target_node != nil {
            if node.N
        }
    }
}

func set_price(a, b *List) {
    sizea := size(a)
    sizeb := size(b)
    curra := a.Head
    currb := b.Head
    for curr != nil {
        if curr.index < med {
            curr.above_median = true
        }else {
            curr.above_median = false
        }
        curr = curr.Next
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
    var min *Node
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