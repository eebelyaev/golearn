package main

import "testing"

func TestTopoSort(t *testing.T) {
	var tests = []struct {
		input      map[string]map[string]struct{}
		isCycles   bool
		orderCount int
	}{
		{
			input: map[string]map[string]struct{}{
				"a": {"b": {}},
				"b": {"c": {}},
			},
			isCycles:   false,
			orderCount: 3,
		},
		{
			input: map[string]map[string]struct{}{
				"a": {"b": {}},
				"b": {"a": {}},
			},
			isCycles:   true,
			orderCount: 2,
		},
		{
			input: map[string]map[string]struct{}{
				"a": {"b": {}},
				"b": {"c": {}},
				"c": {"a": {}},
			},
			isCycles:   true,
			orderCount: 3,
		},
		{
			input: map[string]map[string]struct{}{
				"a": {"b": {}},
				"b": {"c": {}},
				"c": {"d": {}},
				"d": {"a": {}},
			},
			isCycles:   true,
			orderCount: 4,
		},
		{
			input: map[string]map[string]struct{}{
				"a": {"b": {}, "d": {}},
				"b": {"c": {}},
				"c": {"d": {}},
			},
			isCycles:   false,
			orderCount: 4,
		},
	}
	for i, test := range tests {
		list, err := topoSort(test.input)
		if (err == nil) == test.isCycles {
			t.Errorf("#%d isCycles: expected = %t, got = %t",
				i, test.isCycles, err != nil)
		} else if err == nil && len(list) != test.orderCount {
			t.Errorf("#%d orderCount: expected = %d,  got = %d",
				i, test.orderCount, len(list))
		}
	}
}
