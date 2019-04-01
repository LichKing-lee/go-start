// 배열을 map으로 변환하는 프로그램을 작성하라.
package main

func main() {
	MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}

func MaximumSubarraySum(numbers []int) int {
	if isAllLessThanZero(numbers) {
		return 0
	}

	max := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < 1 {
			continue
		}

		for j := len(numbers) - 1; j >= 0; j-- {
			if j < i {
				break
			}

			if numbers[j] < 1 {
				continue
			}

			sum := sum(numbers[i:(j+1)])

			if max < sum {
				max = sum
			}
		}
	}

	return max
}

func sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func isAllLessThanZero(numbers []int) bool {
	for _, n := range numbers {
		if n > 0 {
			return false
		}
	}

	return true
}