package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"mantevian.xyz/advent_of_code_2024/util"
)

type rule struct {
	a int
	b int
}

func Day5Task1() {
	file, _ := os.Open("./5.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := []rule{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		vals := strings.Split(line, "|")
		rules = append(rules, rule{a: util.StringToInt(vals[0]), b: util.StringToInt(vals[1])})
	}

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		ints := util.Map(strings.Split(line, ","), util.StringToInt)

		result := true
		for _, rule := range rules {
			posA := -1
			posB := -1

			for i, n := range ints {
				if n == rule.a {
					posA = i
				}
				if n == rule.b {
					posB = i
				}
			}

			if posA > -1 && posB > -1 && posB < posA {
				result = false
				break
			}
		}

		if result {
			sum += ints[len(ints)/2]
		}
	}

	fmt.Println(sum)
}

func Day5Task2() {
	file, _ := os.Open("./5.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := []rule{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		vals := strings.Split(line, "|")
		rules = append(rules, rule{a: util.StringToInt(vals[0]), b: util.StringToInt(vals[1])})
	}

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		ints := util.Map(strings.Split(line, ","), util.StringToInt)
		prevCounts := make([]int, len(ints))

		for _, rule := range rules {
			for i, n := range ints {
				if n == rule.b {
					prevCounts[i]++
				}
			}
		}

		result := []int{}
		finishedCount := 0

		for finishedCount < len(ints) {
			i := slices.Index(prevCounts, 0)
			if i == -1 {
				result = append(result, ints[finishedCount])
				finishedCount++
				continue
			}

			for j := range ints {
				for _, rule := range rules {
					if ints[i] == rule.a && ints[j] == rule.b {
						prevCounts[j]--
					}
				}
			}
			result = append(result, ints[i])
			prevCounts[i] = -1
			finishedCount++
		}

		fmt.Println(result)
		sum += result[len(result)/2]
	}

	fmt.Println(sum)
}
