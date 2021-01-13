package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if scanner.Err() != nil {
		return lines, scanner.Err()
	}
	return lines, nil
}

func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}
	count := 0
	for i, line := range input {
		char := line[(i*3)%len(line)]
		if string(char) == "#" {
			count++
		}
	}
	fmt.Println(count)
}
