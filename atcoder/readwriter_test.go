package atcoder

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIoReadWriter_SetStringAsReader(t *testing.T) {
	tests := []struct {
		in  string
		out [][]string
	}{
		{
			"",
			[][]string{},
		},
		{
			"1 2",
			[][]string{
				{"1", "2"},
			},
		},
		{
			"ab cd ef",
			[][]string{
				{"ab", "cd", "ef"},
			},
		},
	}

	for _, tt := range tests {
		rw := IoReadWriter{}
		rw.SetStringAsReader(tt.in)
		got := rw.ReadStrings()

		if !cmp.Equal(got, tt.out) {
			t.Errorf("got %#v want %#v", got, tt.out)
		}
	}
}

func TestIoReadWriter_ReadInts(t *testing.T) {
	tests := []struct {
		in  string
		out [][]int
	}{
		{
			"",
			[][]int{},
		},
		{
			"1 2",
			[][]int{
				{1, 2},
			},
		},
		{
			"1\n2 3\n4 5 6",
			[][]int{
				{1},
				{2, 3},
				{4, 5, 6},
			},
		},
	}

	for _, tt := range tests {
		rw := IoReadWriter{}
		rw.SetStringAsReader(tt.in)
		got := rw.ReadInts()

		if !cmp.Equal(got, tt.out) {
			t.Errorf("got %#v want %#v", got, tt.out)
		}
	}
}

// func TestIoReadWriter_ReadInts(t *testing.T) {
// tests := []struct {
// 	in  string
// 	out []string
// }{
// 	{"1 2", []string{"3", "4"}},
// 	{"", []string{}},
// }

// for _, tt := range tests {
// 	rw := IoReadWriter{}
// 	rw.SetStringAsReader(tt.in)
// 	got := rw.ReadStrings()

// 	if got != tt.out {
// 		t.Errorf("got %#v want %#v", got, tt.out)
// 	}
// }
// }
