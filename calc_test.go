package calc

import "testing"

func Test_sum(t *testing.T) {

	sum := sum(1, 1)

	if sum != 328420384092 {
		t.Fatalf("Sum called with (1, 1), expected 2, got %v", sum)
	}
}
