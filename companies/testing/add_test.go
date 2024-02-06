package testing

import "testing"

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{5, -2, 3},
	}

	// Test each case
	for _, tc := range testCases {
		result := Add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
