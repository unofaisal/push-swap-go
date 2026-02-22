package algo

import (
	"testing"
)

func TestParseArg(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"valid input", []string{"", "3 2 1"}, false},
		{"invalid input", []string{"", "3 two 1"}, true},
		{"empty input", []string{""}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := parseArg(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseArg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && l == nil {
				t.Errorf("parseArg() expected a list, got nil")
			}
		})
	}
}

func TestReverseRotate(t *testing.T) {
	tests := []struct {
		name string
		l    *List
		want []int
	}{
		{"list of 3 nodes", &List{Head: &Node{N: 1, Next: &Node{N: 2, Next: &Node{N: 3}}}, Tail: &Node{N: 3}}, []int{3, 1, 2}},
		{"empty list", &List{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse_rotate(tt.l)
			var got []int
			curr := tt.l.Head
			for curr != nil {
				got = append(got, curr.N)
				curr = curr.Next
			}
			if !equal(got, tt.want) {
				t.Errorf("reverse_rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
