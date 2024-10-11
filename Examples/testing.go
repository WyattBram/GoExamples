package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(twoSum(arr, 6))

}

func twoSum(arr []int, target int) []int {

	hashMap := map[int]int{}

	for i, value := range arr {
		comp := target - value

		val, ok := hashMap[value]

		if ok {
			return []int{i, val}
		}

		hashMap[comp] = i
	}
	return []int{-1, -1}
}
