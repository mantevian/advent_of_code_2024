package util

import (
	"math"
	"strconv"
)

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func MapString[T any](input string, f func(rune) T) []T {
	res := make([]T, len(input))
	for i, r := range input {
		res[i] = f(r)
	}
	return res
}

func StringToInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func RuneToInt(r rune) int {
	n, _ := strconv.Atoi(string(r))
	return n
}

func CountDigits(n int) int {
	result := 0
	for n > 0 {
		result++
		n /= 10
	}
	return result
}

func ConcatInts(a int, b int) int {
	return a*int(math.Pow10(CountDigits(b))) + b
}

func Mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}
