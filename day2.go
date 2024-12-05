package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkReport(input []int) bool {
	checkedAsc := false
	asc := false
	result := true

	for i := 1; i < len(input); i++ {
		curr := input[i]
		prev := input[i-1]

		if curr == prev {
			result = false
			break
		}

		if !checkedAsc {
			asc = curr > prev
			checkedAsc = true
		}

		if asc && (curr < prev || curr > prev+3) {
			result = false
			break
		}

		if !asc && (curr > prev || curr < prev-3) {
			result = false
			break
		}
	}

	return result
}

func Day2Task1() {
	text, _ := os.ReadFile("./2.txt")

	lines := strings.Split(string(text), "\n")

	c := 0

	for _, l := range lines {
		fields := strings.Fields(l)

		report := []int{}

		for _, f := range fields {
			n, _ := strconv.Atoi(f)
			report = append(report, n)
		}

		if checkReport(report) {
			c++
		}
	}

	fmt.Println(c)
}

func Day2Task2() {
	text, _ := os.ReadFile("./2.txt")

	lines := strings.Split(string(text), "\n")

	c := 0

	for _, l := range lines {
		fields := strings.Fields(l)

		report := []int{}

		for _, f := range fields {
			n, _ := strconv.Atoi(f)
			report = append(report, n)
		}

		safe := checkReport(report)

		if safe {
			c++
			continue
		}

		for i := range len(report) {
			reportWithOneRemoved := []int{}
			for j := range len(report) {
				if i != j {
					reportWithOneRemoved = append(reportWithOneRemoved, report[j])
				}
			}

			if checkReport(reportWithOneRemoved) {
				safe = true
				break
			}
		}

		if safe {
			c++
		}
	}

	fmt.Println(c)
}
