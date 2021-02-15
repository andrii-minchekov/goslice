package calc

import "github.com/andrii-minchekov/goslice/model"

const Calc = 100

func Multiply(nums model.NumSlice) int {
	mult := 1
	for _, num := range nums {
		mult = num * mult
	}
	return mult
}
