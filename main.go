package main

import (
	"github.com/andrii-minchekov/goslice/calc"
	"github.com/andrii-minchekov/goslice/model"
)

func main() {
	slice := model.NumSlice{1, 2, 3, 10}
	println("slice sum =", slice.Sum())
	println("slice mult =", calc.Multiply(slice))
}
