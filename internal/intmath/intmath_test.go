package intmath

import "testing"

func TestExp2(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"2^-1=0", -1, 0},
		{"2^0=1", 0, 1},
		{"2^1=2", 1, 2},
		{"2^2=4", 2, 4},
		{"2^3=8", 3, 8},
		{"2^4=16", 4, 16},
		{"2^5=32", 5, 32},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Exp2(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v instead", tc.expected, actual)
			}
		})
	}
}

func TestLog2(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"log2(0)=0", 0, 0},
		{"log2(1)=0", 1, 0},
		{"log2(2)=1", 2, 1},
		{"log2(4)=2", 4, 2},
		{"log2(8)=3", 8, 3},
		{"log2(16)=4", 16, 4},
		{"log2(31)=4", 31, 4},
		{"log2(32)=5", 32, 5},
		{"log2(33)=5", 33, 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Log2(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v instead", tc.expected, actual)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		low      int
		high     int
		expected int
	}{
		{"wrap(0,0,4)=0", 0, 0, 4, 0},
		{"wrap(4,0,4)=4", 4, 0, 4, 4},
		{"wrap(5,0,4)=0", 5, 0, 4, 0},
		{"wrap(-1,0,4)=0", -1, 0, 4, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Wrap(tc.input, tc.low, tc.high)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v instead", tc.expected, actual)
			}
		})
	}
}
