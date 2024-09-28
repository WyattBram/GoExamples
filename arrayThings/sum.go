package main

func Sum(arr []int) int {

	sum := 0

	for _, number := range arr {
		sum += number
	}

	return sum

}

func SumAll(numbersToSum ...[]int) []int {

	length := len(numbersToSum)

	returnArr := make([]int, length)

	for i, numbers := range numbersToSum {

		returnArr[i] = Sum(numbers)
	}

	var arr []int

	for _, numbers := range numbersToSum {
		arr = append(arr, Sum(numbers))
	}
	return arr
}

func SumAllTails(numberOfSlices ...[]int) []int {

	lengthOfArray := len(numberOfSlices)

	res := make([]int, lengthOfArray)

	for i, num := range numberOfSlices {
		if len(num) == 0 {
			res[i] = 0
		} else {
			res[i] = Sum(num[1:])
		}
	}
	return res

}
