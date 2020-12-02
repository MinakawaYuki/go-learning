package main

import "fmt"

func main() {
	var arrayTest []int
	fmt.Println(arrayTest)
	array()
}

func array() {
	var a = [3][5]int{{1, 2, 3, 4, 5}, {0, 9, 8, 7, 6}, {3, 4, 5, 6, 7}}
	var i, j int
	for i = 0; i < 3; i++ {
		for j = 0; j < 5; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
