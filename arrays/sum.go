package arrays

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Arrays have fixed length. Slices have a dynamic length.
// Arrays also encode the size of the array in the type. Slices do not.
// The only difference in declaring one or the other is initializing the array with a size.

// go test --cover to run test coverage

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, number := range numbers {
		sums = append(sums, Sum(number))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, numbers := range numbers {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	fmt.Printf("sums: %v\n", sums)

	return sums
}
