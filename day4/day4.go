package day4

import (
	"os"
	"strings"
)

func Day4Task1() {
	bytes, _ := os.ReadFile("./4.txt")
	text := string(bytes)

	lines := strings.Split(text, "\n")

	c := 0

	l := len(lines)

	for x := range l {
		for y := range l {
			if lines[x][y] != 'X' {
				continue
			}

			for xmod := -1; xmod <= 1; xmod++ {
				for ymod := -1; ymod <= 1; ymod++ {
					if x+xmod*3 >= 0 && x+xmod*3 < l && y+ymod*3 >= 0 && y+ymod*3 < l {
						if lines[x+xmod][y+ymod] == 'M' && lines[x+xmod*2][y+ymod*2] == 'A' && lines[x+xmod*3][y+ymod*3] == 'S' {
							c++
						}
					}
				}
			}
		}
	}

	println(c)
}

func Day4Task2() {
	bytes, _ := os.ReadFile("./4.txt")
	text := string(bytes)

	lines := strings.Split(text, "\n")

	c := 0

	l := len(lines)

	for x := 1; x < l-1; x++ {
		for y := 1; y < l-1; y++ {
			if lines[x][y] != 'A' {
				continue
			}

			if lines[x-1][y-1] == 'M' && lines[x+1][y+1] == 'S' && lines[x-1][y+1] == 'M' && lines[x+1][y-1] == 'S' {
				c++
			}

			if lines[x-1][y-1] == 'S' && lines[x+1][y+1] == 'M' && lines[x-1][y+1] == 'S' && lines[x+1][y-1] == 'M' {
				c++
			}

			if lines[x-1][y-1] == 'M' && lines[x+1][y+1] == 'S' && lines[x-1][y+1] == 'S' && lines[x+1][y-1] == 'M' {
				c++
			}

			if lines[x-1][y-1] == 'S' && lines[x+1][y+1] == 'M' && lines[x-1][y+1] == 'M' && lines[x+1][y-1] == 'S' {
				c++
			}
		}
	}

	println(c)
}
