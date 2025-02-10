package gobytes

import "testing"

func TestBoolsToBytes(t *testing.T) {
	tests := []struct {
		input    []bool
		expected []byte
	}{
		{
			[]bool{false, false, false, false, false, false, false, true},
			[]byte{1},
		},
		{
			[]bool{false, false, false, false, false, false, true, true},
			[]byte{3},
		},
		{
			[]bool{false, false, false, false, false, false, true, false},
			[]byte{2},
		},
		{
			[]bool{false, false, false, false, false, false, true},
			[]byte{2},
		},
		{
			[]bool{false, false, false, false, false, false, true, false, true},
			[]byte{2, 128},
		},
	}

	for _, tt := range tests {
		result := BoolsToBytes(tt.input)

		if len(result) != len(tt.expected) {
			t.Errorf("wrong result len: expected %d, got %d", len(tt.expected), len(result))
			continue
		}
		for i, b := range tt.expected {
			if result[i] != b {
				t.Errorf("wrong byte at index %d: expected %d, got %d", i, b, result[i])
			}
		}
	}
}
