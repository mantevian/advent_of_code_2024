package day14

import (
	"fmt"
	"os"
	"strings"

	"mantevian.xyz/advent_of_code_2024/util"
)

func Day14Task1() {
	bytes, _ := os.ReadFile("./14.txt")
	lines := strings.Split(string(bytes), "\n")

	quads := []int{0, 0, 0, 0}

	width := 101
	height := 103

	for _, l := range lines {
		trimmed := strings.ReplaceAll(l, "p", "")
		trimmed = strings.ReplaceAll(trimmed, "=", "")
		trimmed = strings.ReplaceAll(trimmed, "v", "")
		trimmed = strings.ReplaceAll(trimmed, ",", " ")

		data := strings.Fields(trimmed)

		px := util.StringToInt(data[0])
		py := util.StringToInt(data[1])
		vx := util.StringToInt(data[2])
		vy := util.StringToInt(data[3])

		cx := px
		cy := py

		for range 100 {
			cx = util.Mod(cx+vx, width)
			cy = util.Mod(cy+vy, height)
		}

		// fmt.Println(cx, cy)

		if cx < width/2 && cy < height/2 {
			quads[0]++
			continue
		}

		if cx > width/2 && cy < height/2 {
			quads[1]++
			continue
		}

		if cx > width/2 && cy > height/2 {
			quads[2]++
			continue
		}

		if cx < width/2 && cy > height/2 {
			quads[3]++
		}
	}

	fmt.Println(quads)
	fmt.Println(quads[0] * quads[1] * quads[2] * quads[3])
}
