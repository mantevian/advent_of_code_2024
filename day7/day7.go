package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"mantevian.xyz/advent_of_code_2024/util"
)

func Day7Task1() {
	bytes, _ := os.ReadFile("7.txt")
	text := string(bytes)
	text = strings.ReplaceAll(text, ":", "")
	lines := strings.Split(text, "\n")

	sum := 0

	for _, line := range lines {
		vals := util.Map(strings.Fields(line), util.StringToInt)

		power := len(vals) - 2

		answer := false
		for i := range int(math.Pow(2, float64(power))) {
			result := vals[1]
			for j := range power {
				switch (i >> j) & 1 {
				case 0:
					result += vals[j+2]

				case 1:
					result *= vals[j+2]
				}
			}

			if result == vals[0] {
				answer = true
				break
			}
		}

		if answer {
			sum += vals[0]
		}
	}

	fmt.Println(sum)
}

func Day7Task2() {
	bytes, _ := os.ReadFile("7.txt")
	text := string(bytes)
	text = strings.ReplaceAll(text, ":", "")
	lines := strings.Split(text, "\n")

	sum := 0

	for _, line := range lines {
		vals := util.Map(strings.Fields(line), util.StringToInt)

		power := len(vals) - 2

		answer := false
		for i := range int(math.Pow(2, float64(power*2))) {
			result := vals[1]
			for j := range power {
				switch (i >> (j * 2)) & 3 {
				case 0, 3:
					result += vals[j+2]

				case 1:
					result *= vals[j+2]

				case 2:
					result = util.ConcatInts(result, vals[j+2])
				}
			}

			if result == vals[0] {
				answer = true
				break
			}
		}

		if answer {
			sum += vals[0]
		}
	}

	fmt.Println(sum)
}
