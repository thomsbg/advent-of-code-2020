package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var exists = struct{}{}
var requiredKeys = map[string]struct{}{
	"byr": exists,
	"iyr": exists,
	"eyr": exists,
	"hgt": exists,
	"hcl": exists,
	"ecl": exists,
	"pid": exists,
}

// Passport tktk
type Passport map[string]string

func parseLine(p *Passport, line *string) error {
	tokens := strings.Split(*line, " ")
	for _, token := range tokens {
		parts := strings.Split(token, ":")
		(*p)[parts[0]] = (*p)[parts[1]]
	}
	return nil
}

func (p *Passport) isValid() bool {
	for k := range requiredKeys {
		if _, ok := (*p)[k]; !ok {
			return false
		}
	}
	return true
}

func readInput() (*[]Passport, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var passports []Passport
	passport := make(Passport)
	parsed := false
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			passports = append(passports, passport)
			passport = make(Passport)
			parsed = false
		} else {
			err := parseLine(&passport, &line)
			if err != nil {
				return &passports, err
			} else {
				parsed = true
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return &passports, err
	}
	if parsed {
		passports = append(passports, passport)
	}
	return &passports, nil
}

func main() {
	passports, err := readInput()
	if err != nil {
		panic(err)
	}
	count := 0
	for _, passport := range *passports {
		if passport.isValid() {
			count++
		}
	}
	fmt.Println(count)
}
