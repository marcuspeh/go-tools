package tools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerics_Ptr(t *testing.T) {
	tests := []struct {
		name string
		args any
		want any
	}{
		{
			name: "test int",
			args: 1,
			want: 1,
		},
		{
			name: "test string",
			args: "test",
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argsPointer := Ptr(tt.args)
			require.Equal(t, tt.want, *argsPointer)
		})
	}
}
