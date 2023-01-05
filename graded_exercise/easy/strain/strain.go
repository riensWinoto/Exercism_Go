package main

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var result Ints
	for _, val := range i {
		if filter(val) {
			result = append(result, val)
		}
	}
	return result
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var result Ints
	for _, val := range i {
		if !filter(val) {
			result = append(result, val)
		}
	}
	return result
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var result Lists
	for _, val := range l {
		if filter(val) {
			result = append(result, val)
		}
	}
	return result
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var result Strings
	for _, val := range s {
		if filter(val) {
			result = append(result, val)
		}
	}
	return result
}
