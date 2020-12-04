package day04

import (
	"errors"
	"fmt"
	"github.com/knalli/aoc"
	"regexp"
	"strconv"
	"strings"
)

var eyeColor *regexp.Regexp
var nineDigits *regexp.Regexp

func init() {
	eyeColor, _ = regexp.Compile("^[a-f0-9]{6}$")
	nineDigits, _ = regexp.Compile("^[0-9]{9}$")
}

func solve1(lines []string) error {
	if passports, err := parsePassports(lines); err != nil {
		return err
	} else {
		count := 0
		for _, p := range passports {
			if validPassport(p) {
				count++
			}
		}
		aoc.PrintSolution(fmt.Sprintf("There are %d valid passports.", count))
		return nil
	}
}

func solve2(lines []string) error {
	if passports, err := parsePassports(lines); err != nil {
		return err
	} else {
		count := 0
		for _, p := range passports {
			if validPassport2(p) {
				count++
			}
		}
		aoc.PrintSolution(fmt.Sprintf("There are %d valid passports.", count))
		return nil
	}
}

func parsePassports(lines []string) ([]*passportType, error) {

	blocks := strings.Split(strings.Join(lines, "\n"), "\n\n")

	result := make([]*passportType, 0)

	for _, b := range blocks {
		// ensure newlines are normal whitespaces also
		block := strings.Replace(b, "\n", " ", -1)

		tokens := make([]tokenType, 0)
		for _, item := range strings.Split(block, " ") {
			split := strings.Split(item, ":")
			if len(split) != 2 {
				return nil, errors.New("invalid token item")
			}
			tokens = append(tokens, tokenType{key: split[0], val: split[1]})
		}
		result = append(result, &passportType{tokens: tokens})
	}

	return result, nil
}

func validPassport(p *passportType) bool {
	for _, k := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if !p.existTokenKey(k) {
			return false
		}
	}
	return true
}

func validPassport2(p *passportType) bool {

	if !validPassport(p) {
		return false
	}

	if val, err := p.getTokenValue("byr"); err != nil {
		return false
	} else {
		if n, err := strconv.Atoi(val); err != nil {
			return false
		} else if !(1920 <= n && n <= 2002) {
			return false
		}
	}

	if val, err := p.getTokenValue("iyr"); err != nil {
		return false
	} else {
		if n, err := strconv.Atoi(val); err != nil {
			return false
		} else if !(2010 <= n && n <= 2020) {
			return false
		}
	}

	if val, err := p.getTokenValue("eyr"); err != nil {
		return false
	} else {
		if n, err := strconv.Atoi(val); err != nil {
			return false
		} else if !(2020 <= n && n <= 2030) {
			return false
		}
	}

	if val, err := p.getTokenValue("hgt"); err != nil {
		return false
	} else {
		if strings.HasSuffix(val, "cm") {
			if n, err := strconv.Atoi(strings.TrimSuffix(val, "cm")); err != nil {
				return false
			} else if !(150 <= n && n <= 193) {
				return false
			}
		} else if strings.HasSuffix(val, "in") {
			if n, err := strconv.Atoi(strings.TrimSuffix(val, "in")); err != nil {
				return false
			} else if !(59 <= n && n <= 76) {
				return false
			}
		} else {
			return false
		}
	}

	if val, err := p.getTokenValue("hcl"); err != nil {
		return false
	} else {
		if strings.HasPrefix(val, "#") {
			if !eyeColor.MatchString(strings.TrimPrefix(val, "#")) {
				return false
			}
		} else {
			return false
		}
	}

	if val, err := p.getTokenValue("ecl"); err != nil {
		return false
	} else {
		match := false
		for _, color := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if val == color {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}

	if val, err := p.getTokenValue("pid"); err != nil {
		return false
	} else {
		if !nineDigits.MatchString(val) {
			return false
		}
	}

	return true
}

type passportType struct {
	tokens []tokenType
}

func (p *passportType) existTokenKey(key string) bool {
	for _, t := range p.tokens {
		if t.key == key {
			return true
		}
	}
	return false
}

func (p *passportType) getTokenValue(key string) (string, error) {
	for _, t := range p.tokens {
		if t.key == key {
			return t.val, nil
		}
	}
	return "", errors.New("invalid key")
}

type tokenType struct {
	key string
	val string
}
