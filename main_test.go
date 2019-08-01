package main

import "testing"

func TestGetRange(t *testing.T) {
	type Test struct {
		value string
		min   int
		max   int
	}
	var arrTests = []Test{
		{"1", 1, 1},
		{"1-5", 1, 5},
		{"0", 1, 1},
		{"7", 7, 7},
		{"5-", 5, -1},
		{"-3", 1, 3},
		{"5-1", 1, 5},
		{"-1-1", 1, 1},
		{"a-b", 1, 1},
	}
	var resMin, resMax int
	for _, test := range arrTests {
		resMin, resMax = getRange(test.value)
		if resMin != test.min || resMax != test.max {
			t.Errorf("Ожидалось (%d, %d), пришло (%d, %d) = getRange(%s)",
				test.min, test.max, resMin, resMax, test.value)
		}
	}
}

func TestParser(t *testing.T) {
	//
}
