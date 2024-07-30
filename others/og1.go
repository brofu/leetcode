package others

import "sort"

type Record struct {
	Brand string
	Days  int
}

func resolve(input []Record) bool {

	// sort
	sort.Slice(input, func(a, b int) bool {
		if input[a].Days < input[b].Days {
			return true
		}
		return false
	})

	// check if the records is valid
	previousBrand := ""
	for _, r := range input {
		if r.Brand == "JJ" {
			return true
		}

		if previousBrand == r.Brand {
			return true
		}

		previousBrand = r.Brand
	}

	return false
}

func resolveV2(input []Record) (string, int) {
	validDays := 14
	// sort
	sort.Slice(input, func(a, b int) bool {
		if input[a].Days < input[b].Days {
			return true
		}
		return false
	})

	minDays := map[string]int{
		"PZ": 17,
		"MD": 24,
		"CV": 13,
	}

	// check if the records is valid
	var previous Record
	for _, r := range input {
		if r.Brand == "JJ" {
			return "vaccinated", r.Days + validDays
		}

		if previous.Brand == r.Brand {
			if r.Days-previous.Days >= minDays[previous.Brand] {
				return "vaccinated", r.Days + validDays
			}
		}

		previous = r
	}

	return "not_vaccinated", -1

}
