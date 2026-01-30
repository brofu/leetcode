package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertAndMergeAudioRegion(t *testing.T) {
	// Define test cases using a struct
	testCases := []struct {
		name            string
		enabled         bool
		existingRegions [][]int
		newRegion       []int
		want            [][]int
	}{
		// Example 1: New region overlaps with the first existing region
		{
			name:            "Example 1: Overlap at the beginning",
			enabled:         true,
			existingRegions: [][]int{{0, 2}, {4, 6}},
			newRegion:       []int{1, 3},
			want:            [][]int{{0, 3}, {4, 6}},
		},
		// Example 2: New region fits between two existing regions
		{
			name:            "Example 2: No overlap, insert in middle",
			enabled:         true,
			existingRegions: [][]int{{1, 2}, {7, 8}},
			newRegion:       []int{4, 5},
			want:            [][]int{{1, 2}, {4, 5}, {7, 8}},
		},
		// Example 3: New region completely covers all existing regions
		{
			name:            "Example 3: New region covers all",
			enabled:         true,
			existingRegions: [][]int{{1, 5}, {7, 8}},
			newRegion:       []int{0, 10},
			want:            [][]int{{0, 10}},
		},
		// Example 4: Inserting into an empty list
		{
			name:            "Example 4: Empty initial list",
			enabled:         true,
			existingRegions: [][]int{},
			newRegion:       []int{2, 5},
			want:            [][]int{{2, 5}},
		},
		{
			name:            "Example 5: Insert at the end",
			enabled:         true,
			existingRegions: [][]int{{1, 2}, {3, 4}},
			newRegion:       []int{5, 6},
			want:            [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		// Example 6: New region is before all existing regions
		{
			name:            "Example 6: Insert at the beginning",
			enabled:         true,
			existingRegions: [][]int{{3, 4}, {5, 6}},
			newRegion:       []int{1, 2},
			want:            [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		// Example 7: New region merges multiple existing regions
		{
			name:            "Example 7: Merge multiple regions",
			enabled:         true,
			existingRegions: [][]int{{1, 3}, {4, 6}, {8, 9}},
			newRegion:       []int{2, 8},
			want:            [][]int{{1, 9}},
		},
		// Example 8: New region touches and merges with the first region
		{
			name:            "Example 8: Touch and merge first region",
			enabled:         true,
			existingRegions: [][]int{{3, 4}, {5, 6}},
			newRegion:       []int{1, 3},
			want:            [][]int{{1, 4}, {5, 6}},
		},
		// Example 9: New region touches and merges with a middle region
		{
			name:            "Example 9: Touch and merge middle region",
			enabled:         true,
			existingRegions: [][]int{{1, 3}, {5, 6}},
			newRegion:       []int{5, 8},
			want:            [][]int{{1, 3}, {5, 8}},
		},

		// Example 10: New region overlaps with the second region
		{
			name:            "Example 10: Overlap with second region",
			enabled:         true,
			existingRegions: [][]int{{1, 3}, {5, 6}},
			newRegion:       []int{4, 8},
			want:            [][]int{{1, 3}, {4, 8}},
		},
	}

	// Loop through the test cases
	for _, tc := range testCases {
		if !tc.enabled {
			continue
		}
		t.Run(tc.name, func(t *testing.T) {
			got := insertAndMergeAudioRegion(tc.existingRegions, tc.newRegion)
			if !reflect.DeepEqual(got, tc.want) {
				fmt.Println("existings:", tc.existingRegions, "new:", tc.newRegion)
				fmt.Println("want:", tc.want)
				fmt.Println("got:", got)
			}
		})
	}
}
