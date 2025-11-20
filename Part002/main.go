package main

import "fmt"

func Insert(slice []int, element int, index int) []int {
	if index < 0 || index > len(slice) {
		return slice
	}
	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])
	slice[index] = element
	return slice
}

func Chunk(slice []int, size int) [][]int {
	matrix := make([][]int, size)
	k := 0
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if k < len(slice) {
				matrix[i][j] = slice[k]
				k++
			}
		}
	}
	return matrix
}

func main() {
	//1
	slice := []int{1, 2, 3, 4, 5}
	element := 99
	index := 2
	newSlice := Insert(slice, element, index)
	fmt.Println("Original Slice:", slice)
	fmt.Println("New Slice:     ", newSlice)

	//2
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newSlice1 := Chunk(slice1, 4)
	fmt.Println("Original Slice:", slice1)
	fmt.Println("New Slice:     ", newSlice1)
}
