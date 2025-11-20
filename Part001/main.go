package main

import "fmt"

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func GetEven(nums []int) []int {
	var even []int
	for _, v := range nums {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	return even
}

func Reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func RemoveElement(nums []int, index int) []int {
	tmp := append(nums[:index], nums[index+1:]...)
	return tmp
}

func main() {
	//1
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Sum %d", Sum(a))

	//2
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range GetEven(b) {
		fmt.Println(v)
	}

	//3
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original %v \n", c)
	Reverse(c)
	fmt.Printf("Reverse %v \n", c)

	//4
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Before slice d: %v \n", d)
	e := RemoveElement(d, 3)
	fmt.Printf(" After slice e: %v \n", e)
}
