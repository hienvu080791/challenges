package main

import (
	"demo_challenges/source/router"
)

func main() {
	//testing the algorithms package
	//fmt.Println(algorithms.GrayCode(2))
	//fmt.Println(algorithms.SumOfDistancesInTree(2, [][]int{{0, 1}}))
	//fmt.Println(algorithms.FindLength([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}))

	// Initialize the router
	r := router.Init()
	_ = r.Run(":" + "8888")
}
