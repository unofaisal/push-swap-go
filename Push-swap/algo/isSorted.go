package algo

func isSorted(l *List) bool {
	curr := l.Head
	for curr != nil && curr.Next != nil {
		if curr.N > curr.Next.N {
			return false
		}
		curr = curr.Next
	}
	return true
}
