package model

type NumSlice []int

func (nums NumSlice) Sum() int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
