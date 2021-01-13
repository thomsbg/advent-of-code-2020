package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *Passport) parse(line string) error {
	tokens := strings.Split(line, " ")
	for _, token := range tokens {
		parts := strings.Split(token, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid token: %v", token)
		}
		switch string(parts[0]) {
		case "byr":
			byr, err := strconv.Atoi(parts[1])
			if err != nil {
				return fmt.Errorf("expected int for byr: %v", err)
			} else {
				p.byr = byr
			}
		case "iyr":
			iyr, err := strconv.Atoi(parts[1])
			if err != nil {
				return fmt.Errorf("expected int for iyr: %v", err)
			} else {
				p.iyr = iyr
			}
		case "eyr":
			eyr, err := strconv.Atoi(parts[1])
			if err != nil {
				return fmt.Errorf("expected int for eyr: %v", err)
			} else {
				p.eyr = eyr
			}
		case "hgt":
			p.hgt = parts[1]
		case "hcl":
			p.hcl = parts[1]
		case "ecl":
			p.ecl = parts[1]
		case "pid":
			p.pid = parts[1]
		case "cid":
			p.cid = parts[1]
		default:
			return fmt.Errorf("invalid token key: %s", parts[0])
		}
	}
	return nil
}

func (p *Passport) isValid() bool {
	return p.byr != 0 &&
		p.iyr != 0 &&
		p.eyr != 0 &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != ""
}

func readInput() (*[]Passport, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var passports []Passport
	passport := Passport{}
	parsed := false
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			passports = append(passports, passport)
			passport = Passport{}
			parsed = false
		} else {
			err := passport.parse(line)
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
