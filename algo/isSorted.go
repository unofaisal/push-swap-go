package algo

func IsSorted(l *List) bool {
	curr := l.Head
	for curr != nil {
		if curr.N > curr.Next.N {
			return false
		}
	}
	return true
}