package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSlice_ForEach(t *testing.T) {
	tests := []struct {
		name        string
		args        []interface{}
		expectedIdx []int
	}{
		{
			name:        "int",
			args:        []interface{}{35, 63, 24, 63, 24},
			expectedIdx: []int{0, 1, 2, 3, 4},
		},
		{
			name:        "string",
			args:        []interface{}{"a", "b", "c"},
			expectedIdx: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := make([]interface{}, 0, len(tt.args))
			idxArr := make([]int, 0, len(tt.args))
			ForEach(tt.args, func(val interface{}, idx int) {
				output = append(output, val)
				idxArr = append(idxArr, idx)
			})
			require.Equal(t, tt.expectedIdx, idxArr)
			require.Equal(t, tt.args, output)
		})
	}
}

func TestSlice_Filter(t *testing.T) {
	tests := []struct {
		name     string
		args     []interface{}
		fn       func(val interface{}, idx int) bool
		expected []interface{}
	}{
		{
			name: "int",
			args: []interface{}{35, 63, 28, 63, 24},
			fn: func(val interface{}, idx int) bool {
				return val.(int)%2 == 0
			},
			expected: []interface{}{28, 24},
		},
		{
			name: "string",
			args: []interface{}{"a", "b", "c"},
			fn: func(val interface{}, idx int) bool {
				return idx%2 == 1
			},
			expected: []interface{}{"b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := Filter(tt.args, tt.fn)
			require.Equal(t, tt.expected, output)
		})
	}
}

func TestSlice_Map(t *testing.T) {
	tests := []struct {
		name     string
		args     []interface{}
		fn       func(val interface{}, idx int) interface{}
		expected []interface{}
	}{
		{
			name: "int",
			args: []interface{}{35, 63, 28, 63, 24},
			fn: func(val interface{}, idx int) interface{} {
				return val.(int) * 2
			},
			expected: []interface{}{70, 126, 56, 126, 48},
		},
		{
			name: "string",
			args: []interface{}{"a", "b", "c"},
			fn: func(val interface{}, idx int) interface{} {
				return val.(string) + "1"
			},
			expected: []interface{}{"a1", "b1", "c1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := Map(tt.args, tt.fn)
			require.Equal(t, tt.expected, output)
		})
	}
}

func TestSlice_Find(t *testing.T) {
	tests := []struct {
		name     string
		args     []interface{}
		fn       func(val interface{}, idx int) bool
		expected interface{}
	}{
		{
			name: "int",
			args: []interface{}{35, 63, 28, 63, 24},
			fn: func(val interface{}, idx int) bool {
				return val.(int) == 63
			},
			expected: 63,
		},
		{
			name: "string",
			args: []interface{}{"a", "b", "c"},
			fn: func(val interface{}, idx int) bool {
				return val.(string) == "b"
			},
			expected: "b",
		},
		{
			name: "not found",
			args: []interface{}{"a", "b", "c"},
			fn: func(val interface{}, idx int) bool {
				return val.(string) == "d"
			},
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, ok := Find(tt.args, tt.fn)
			require.Equal(t, tt.expected != nil, ok)
			require.Equal(t, tt.expected, output)
		})
	}
}

func TestSlice_Copy(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
	}{
		{
			name: "int",
			args: []interface{}{35, 63, 28, 63, 24},
		},
		{
			name: "string",
			args: []interface{}{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := Copy(tt.args)
			require.Equal(t, tt.args, output)
		})
	}
}
