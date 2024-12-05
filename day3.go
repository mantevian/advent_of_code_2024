package main

import (
	"fmt"
	"os"
	"strconv"
)

func readInt(input string, pos int) (length int, value int, err error) {
	currPos := pos
	for input[currPos] >= '0' && input[currPos] <= '9' {
		currPos++
	}
	result, err := strconv.Atoi(input[pos:currPos])
	return currPos - pos, result, err
}

func tryMul(input string, pos int) (ok bool, result int) {
	n := pos

	if input[n] != 'm' {
		return false, 0
	}

	if input[n:n+4] != "mul(" {
		return false, 0
	}

	n += 4

	length, lhs, err := readInt(input, n)
	if err != nil || length > 3 {
		return false, 0
	}

	n += length

	if input[n] != ',' {
		return false, 0
	}

	n++

	length, rhs, err := readInt(input, n)
	if err != nil || length > 3 {
		return false, 0
	}

	n += length

	if input[n] != ')' {
		return false, 0
	}

	return true, lhs * rhs
}

func tryDo(input string, pos int) bool {
	if input[pos] != 'd' {
		return false
	}

	return input[pos:pos+4] == "do()"
}

func tryDont(input string, pos int) bool {
	if input[pos] != 'd' {
		return false
	}

	return input[pos:pos+7] == "don't()"
}

func Day3Task1() {
	bytes, _ := os.ReadFile("./3.txt")
	text := string(bytes)

	l := len(text)
	n := 0

	sum := 0

	for n < l {
		ok, result := tryMul(text, n)

		if ok {
			sum += result
		}

		n++
	}

	fmt.Println(sum)
}

func Day3Task2() {
	bytes, _ := os.ReadFile("./3.txt")
	text := string(bytes)

	l := len(text)
	n := 0

	sum := 0

	enabled := true

	for n < l {
		if tryDo(text, n) {
			enabled = true
		}

		if tryDont(text, n) {
			enabled = false
		}

		if enabled {
			ok, result := tryMul(text, n)

			if ok {
				sum += result
			}
		}

		n++
	}

	fmt.Println(sum)
}
