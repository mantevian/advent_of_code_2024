package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day13Task1() {
	bytes, _ := os.ReadFile("13.txt")
	lines := strings.Split(string(bytes), "\n")

	var sum int64 = 0
	for i := range len(lines) / 4 {
		fields1 := strings.Fields(lines[i*4])
		fields2 := strings.Fields(lines[i*4+1])
		fields3 := strings.Fields(lines[i*4+2])

		a11, _ := strconv.ParseFloat(fields1[0], 64)
		a12, _ := strconv.ParseFloat(fields2[0], 64)
		a13, _ := strconv.ParseFloat(fields3[0], 64)
		a21, _ := strconv.ParseFloat(fields1[1], 64)
		a22, _ := strconv.ParseFloat(fields2[1], 64)
		a23, _ := strconv.ParseFloat(fields3[1], 64)

		a13 += 10000000000000.0
		a23 += 10000000000000.0

		// fmt.Println(a11, a12, a13, a21, a22, a23)

		// a11 * x + a12 * y = a13
		// a21 * x + a22 * y = a23

		// x = (a13 - a12 * y) / a11
		// a21 * (a13 - a12 * y) / a11 + a22 * y = a23
		// a21 * a13 / a11 - a21 * a12 * y / a11 + a22 * y = a23
		// y * (a22 - a21 * a12 / a11) = a23 - a21 * a13 / a11
		y := (a23 - a21*a13/a11) / (a22 - a21*a12/a11)
		x := (a13 - a12*y) / a11

		if x >= 0 && y >= 0 && math.Abs(x-math.Round(x)) < 0.0000000001 && math.Abs(y-math.Round(y)) < 0.0000000001 {
			sum += int64(math.Round(x)*3 + math.Round(y))
		}

		fmt.Println(x, y)

	}

	fmt.Println(sum)
}
