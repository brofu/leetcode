package main

import (
	"flag"
	"fmt"

	"github.com/brofu/leetcode/backtrack"
	"github.com/brofu/leetcode/version"
)

func main() {

	versionFlag := flag.Bool("version", false, "Version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Build Date:", version.BuildDate)
		fmt.Println("Git Commit:", version.GitCommit)
		fmt.Println("Version:", version.Version)
		fmt.Println("Go Version:", version.GoVersion)
		fmt.Println("OS / Arch:", version.OsArch)
		return
	}
	fmt.Println("Hello.")

	result := backtrack.Permute([]int{5, 4, 6, 2})
	fmt.Println("Hello.", result)
}

/*
func insertAndMergeAudioRegion(existings [][]int, newRegion []int) [][]int {

	var (
		result [][]int
	)

	// edge case 1
	if len(existings) == 0 {
		return [][]int{newRegion}
	}

	startMap := make(map[int]int)
	for idx, region := range existings {
		startMap[region[0]] = idx
	}

	start, end := existings[0][0], existings[len(existings)-1][0]

	for start < end {
		mid = (start + end) / 2
		if newRegion[0] <= mid {
			end = mid
		} else {
			start = mid + 1
		}
	}
	for mid > 0 {
		hitIndex, exists := startMap[mid]
		if exists {
			break
		}
	}

	return result
}
*/

func insertAndMergeAudioRegion(existings [][]int, newRegion []int) [][]int {

	var (
		startIndex, endIndex = 0, 0
	)

	if len(existings) == 0 {
		return [][]int{newRegion}
	}
	// get the start index of the new region
	for idx, region := range existings {
		if region[0] <= newRegion[0] {
			startIndex = idx
		}
	}

	if newRegion[0] > existings[startIndex][0] {
		newRegion[0] = existings[startIndex][0]
	}

	// get the end index of the new region
	for idx, region := range existings {
		if region[1] < newRegion[1] {
			endIndex = idx
		}
	}
	if endIndex < len(existings)-1 {
		if newRegion[1] >= existings[endIndex+1][0] {
			endIndex += 1
			newRegion[1] = existings[endIndex][1]
		}
	}

	result := make([][]int, 0)

	for idx := 0; idx < len(existings); {
		existRegion := existings[idx]
		if idx < startIndex {
			result = append(result, existRegion)
			idx++
			continue
		}
		if idx == startIndex {
			result = append(result, newRegion)
			for i := idx; i <= endIndex; i++ {
				idx++
			}
			continue
		}
		result = append(result, existings[idx])
		idx++
	}

	return result
}
