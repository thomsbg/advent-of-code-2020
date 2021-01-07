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
	fmt.Println("input:", input)
	i, j, err := findPair(&input, 2020)
	check(err)
	product := input[i] * input[j]
	fmt.Println(input[i], "+", input[j], "= 2020")
	fmt.Println(input[i], "*", input[j], "=", product)
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

func findPair(ptr *[]int, target int) (int, int, error) {
	input := *ptr
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
			panic("This shouldn't happen")
		}
	}

	return 0, 0, errors.New("not found")
}
