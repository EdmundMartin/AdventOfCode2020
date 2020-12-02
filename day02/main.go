package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PasswordRules struct {
	Min        int
	Max        int
	TargetChar string
	TestString string
}

func (p *PasswordRules) String() string {
	return fmt.Sprintf("Min: %d, Max: %d, Target: %s, Raw: %s",
		p.Min, p.Max, p.TargetChar, p.TestString)
}

func (p *PasswordRules) MeetsCondition() bool {
	r := []rune(p.TargetChar)[0]
	count := 0
	for _, ch := range p.TestString {
		if ch == r {
			count++
		}
	}
	if count >= p.Min && count <= p.Max {
		return true
	}
	return false
}

func (p *PasswordRules) MeetsVariantCondition() bool {
	r := []rune(p.TargetChar)[0]
	test := []rune(p.TestString)
	count := 0
	if test[p.Min-1] == r {
		count++
	}
	if test[p.Max-1] == r {
		count++
	}
	return count == 1
}

func NewPasswordRules(contents string) *PasswordRules {
	res := strings.Split(contents, " ")
	miniMax := strings.Split(res[0], "-")
	min, _ := strconv.Atoi(miniMax[0])
	max, _ := strconv.Atoi(miniMax[1])
	testChar := strings.Replace(res[1], ":", "", -1)
	return &PasswordRules{
		Min:        min,
		Max:        max,
		TargetChar: testChar,
		TestString: res[len(res)-1],
	}
}

func main() {
	fileHandle, _ := os.Open("day02/input.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	count := 0
	variantCondition := 0
	for fileScanner.Scan() {
		contents := fileScanner.Text()
		rule := NewPasswordRules(contents)
		passes := rule.MeetsCondition()
		if passes {
			count++
		}
		variantPasses := rule.MeetsVariantCondition()
		if variantPasses {
			variantCondition++
		}
	}
	fmt.Printf("Total First Condition Passed: %d\n", count)
	fmt.Printf("Total Second Condition Passed: %d\n", variantCondition)
}
