package main

import (
	"demo_challenges/algorithms"
	"fmt"
)

func main() {
	fmt.Println(algorithms.GrayCode(2))
	fmt.Println(algorithms.SumOfDistancesInTree(2, [][]int{{0, 1}}))
	fmt.Println(algorithms.FindLength([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}))
}
