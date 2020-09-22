package math

func Sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func Add(a int, b int) int {
	return a + b
}
