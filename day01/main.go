package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := readInput()
	sort.Ints(input)

	i, j, err := findPair(input, 2020)
	check(err)
	fmt.Println(input[i], "+", input[j], "= 2020")
	fmt.Println(input[i], "*", input[j], "=", input[i]*input[j])

	i, j, k, err := findTriple(input, 2020)
	check(err)
	fmt.Println(input[i], "+", input[j], "+", input[k], "= 2020")
	fmt.Println(input[i], "*", input[j], "*", input[k], "=", input[i]*input[j]*input[k])
}

func readInput() []int {
	var input []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		check(err)
		input = append(input, value)
	}
	check(scanner.Err())
	return input
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func findPair(input []int, target int) (int, int, error) {
	i := 0
	j := len(input) - 1

	for i < j {
		sum := input[i] + input[j]
		if sum == target {
			return i, j, nil
		} else if sum < target {
			i = i + 1
		} else if sum > target {
			j = j - 1
		} else {
			panic("this shouldn't happen")
		}
	}

	return 0, 0, errors.New("pair not found")
}

func findTriple(input []int, target int) (int, int, int, error) {
	for i := 0; i < len(input)-3; i++ {
		remainingInput := input[i+1:]
		remainingTarget := target - input[i]
		j, k, err := findPair(remainingInput, remainingTarget)
		if err == nil {
			return i, i + j + 1, i + k + 1, nil
		}
	}

	return 0, 0, 0, errors.New("triple not found")
}
