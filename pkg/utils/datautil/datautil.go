package datautil

import "sort"

// SortInts sorts a slice of integers in ascending order
func SortInts(data []int) []int {
	sort.Ints(data)
	
	return data
}

// FilterStrings filters a slice of strings based on a condition
func FilterStrings(data []string, condition func(string) bool) []string {
	var result []string
	
	for _, v := range data {
		if condition(v) {
			result = append(result, v)
		}
	}

	return result
}

// AggregateInts sums up a slice of integers
func AggregateInts(data []int) int {
	var sum int

	for _, v := range data {
		sum += v
	}

	return sum
}
