package math

import "testing"

func TestSum(t *testing.T) {
	sum := Sum([]int{1, 2, 3, 4})
	expect := 10

	if sum != expect {
		t.Errorf("Expect = %d, Got = %d", expect, sum)
	}
}

func TestAdd(t *testing.T) {
	add := Add(1, 2)
	expect := 3

	if add != expect {
		t.Errorf("Expect = %d, Got = %d", expect, add)
	}
}
