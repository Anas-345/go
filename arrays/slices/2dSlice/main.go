package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}
	for i := 0; i < rows; i++ {
		innerMatrix := []int{}
		for j := 0; j < cols; j++ {
			innerMatrix = append(innerMatrix, i*j)
		}
		matrix = append(matrix, innerMatrix)
	}
	return matrix
}

func main() {
	matrix := createMatrix(5, 10)
	for _, v := range matrix {
		fmt.Println(v)
	}
}
