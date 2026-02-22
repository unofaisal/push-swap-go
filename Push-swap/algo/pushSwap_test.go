package algo

import (
	"testing"
)

func TestPush_swap(t *testing.T) {
	args := []string{"3", "2", "1"}
	Push_swap(args)
}

func TestSize(t *testing.T) {
	l := &List{Head: &Node{N: 1, Next: &Node{N: 2, Next: &Node{N: 3}}}}
	if got := size(l); got != 3 {
		t.Errorf("size() = %v, want %v", got, 3)
	}
}

func TestMax(t *testing.T) {
	l := &List{Head: &Node{N: 1, Next: &Node{N: 2, Next: &Node{N: 3}}}}
	if got := max(l); got.N != 3 {
		t.Errorf("max() = %v, want %v", got.N, 3)
	}
}

func TestMin(t *testing.T) {
	l := &List{Head: &Node{N: 1, Next: &Node{N: 2, Next: &Node{N: 3}}}}
	if got := min(l); got.N != 1 {
		t.Errorf("min() = %v, want %v", got.N, 1)
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name string
		l    *List
		want bool
	}{
		{"sorted list", &List{Head: &Node{N: 1, Next: &Node{N: 2, Next: &Node{N: 3}}}}, true},
		{"unsorted list", &List{Head: &Node{N: 3, Next: &Node{N: 2, Next: &Node{N: 1}}}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSorted(tt.l); got != tt.want {
				t.Errorf("IsSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
