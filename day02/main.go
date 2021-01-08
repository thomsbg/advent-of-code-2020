package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type record struct {
	min      int
	max      int
	chr      rune
	password string
}

func parse(line string) (record, error) {
	rec := record{}
	n, err := fmt.Sscanf(line, "%d-%d %c: %s", &rec.min, &rec.max, &rec.chr, &rec.password)
	if n != 4 {
		return rec, fmt.Errorf("only parsed %d of 4 fields", n)
	}
	return rec, err
}

func (r *record) check1() bool {
	count := strings.Count(string(r.password), string(r.chr))
	return count >= r.min && count <= r.max
}

func (r *record) check2() bool {
	return (rune(r.password[r.min-1]) == r.chr) != (rune(r.password[r.max-1]) == r.chr)
}

func readInput() ([]record, error) {
	var input []record
	err := eachInputLine(func(line string) error {
		rec, err := parse(line)
		if err != nil {
			return err
		}
		input = append(input, rec)
		return nil
	})
	return input, err
}

func eachInputLine(cb func(line string) error) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := cb(scanner.Text()); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}
	numValid := 0
	for _, rec := range input {
		if rec.check2() {
			numValid++
		}
	}
	fmt.Println(numValid)
}
