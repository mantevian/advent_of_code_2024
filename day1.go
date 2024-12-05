package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day1Task1() {
	text, _ := os.ReadFile("./1.txt")

	var numbers1 []int
	var numbers2 []int

	lines := strings.Split(string(text), "\n")

	for _, l := range lines {
		fields := strings.Fields(l)

		n1, _ := strconv.Atoi(fields[0])
		n2, _ := strconv.Atoi(fields[1])
		numbers1 = append(numbers1, n1)
		numbers2 = append(numbers2, n2)
	}

	slices.Sort(numbers1)
	slices.Sort(numbers2)

	sum := 0

	for i := range len(numbers1) {
		sum += int(math.Abs(float64(numbers1[i]) - float64(numbers2[i])))
	}

	fmt.Println(sum)
}

func Day1Task2() {
	text, _ := os.ReadFile("./1.txt")

	var numbers1 []int
	var numbers2 []int

	lines := strings.Split(string(text), "\n")

	n := len(lines)

	for _, l := range lines {
		fields := strings.Fields(l)

		n1, _ := strconv.Atoi(fields[0])
		n2, _ := strconv.Atoi(fields[1])
		numbers1 = append(numbers1, n1)
		numbers2 = append(numbers2, n2)
	}

	sum := 0

	for i := range n {
		c := 0

		for j := range n {
			if numbers2[j] == numbers1[i] {
				c++
			}
		}

		sum += numbers1[i] * c
	}

	fmt.Println(sum)
}
