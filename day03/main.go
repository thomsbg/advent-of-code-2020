package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const tree = `#`

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

func check1(i int, line string) int {
	return (i * 3)
}

func check2(i int, line string) int {
	return i
}

func check3(i int, line string) int {
	return (i * 5)
}

func check4(i int, line string) int {
	return (i * 7)
}

func check5(i int, line string) int {
	if i%2 == 0 {
		return i / 2
	} else {
		return strings.Index(line, ".")
	}
}

func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}
	count1 := 0
	count2 := 0
	count3 := 0
	count4 := 0
	count5 := 0
	for i, line := range input {
		if string(line[check1(i, line)%len(line)]) == tree {
			count1++
		}
		if string(line[check2(i, line)%len(line)]) == tree {
			count2++
		}
		if string(line[check3(i, line)%len(line)]) == tree {
			count3++
		}
		if string(line[check4(i, line)%len(line)]) == tree {
			count4++
		}
		if string(line[check5(i, line)%len(line)]) == tree {
			count5++
		}
	}
	fmt.Println(count1)
	fmt.Println(count2)
	fmt.Println(count3)
	fmt.Println(count4)
	fmt.Println(count5)
	fmt.Println(count1 * count2 * count3 * count4 * count5)
}
