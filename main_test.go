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
			t.Errorf("Ожидалось (%d, %d), получили (%d, %d) = getRange(%s)",
				test.min, test.max, resMin, resMax, test.value)
		}
	}
}

func TestCreateAddr(t *testing.T) {
	result := createAddr("http", "zaycev.net", "search.html", "one two three", -1)
	if result != "http://zaycev.net/search.html?query_search=one+two+three" {
		t.Errorf("Ожидалось %s, получили %s", result, "http://zaycev.net/search.html?query_search=one+two+three")
	}
	result = createAddr("http", "zaycev.net", "search.html", "один+два+три", 2)
	if result != "http://zaycev.net/search.html?page=2&query_search=%D0%BE%D0%B4%D0%B8%D0%BD%2B%D0%B4%D0%B2%D0%B0%2B%D1%82%D1%80%D0%B8" {
		t.Errorf("Ожидалось %s, получили %s", result,
			"http://zaycev.net/search.html?page=2&query_search=%D0%BE%D0%B4%D0%B8%D0%BD%2B%D0%B4%D0%B2%D0%B0%2B%D1%82%D1%80%D0%B8")
	}
	result = createAddr("https", "zaycev.net", "search.html", "123", 1)
	if result != "https://zaycev.net/search.html?query_search=123" {
		t.Errorf("Ожидалось %s, получили %s", result, "https://zaycev.net/search.html?query_search=123")
	}
}

func TestParse(t *testing.T) {
	//
}
