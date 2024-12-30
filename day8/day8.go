package main

import (
	"fmt"
	"os"
	"strings"

	"mantevian.xyz/advent_of_code_2024/util"
)

type vec struct {
	x int
	y int
}

type cell struct {
	pos  vec
	freq rune
	anti bool
}

func (c cell) String() string {
	if c.anti {
		return "#"
	}
	return string(c.freq)
}

type world struct {
	size  vec
	cells []cell
}

func (w world) String() string {
	res := ""
	for y := range w.size.y {
		for x := range w.size.x {
			res += w.cells[y*w.size.x+x].String()
		}
		res += "\n"
	}
	return res
}

func (w *world) getCell(pos vec) *cell {
	return &(w.cells[pos.y*w.size.x+pos.x])
}

func (w *world) inBounds(pos vec) bool {
	return pos.x >= 0 && pos.y >= 0 && pos.x < w.size.x && pos.y < w.size.y
}

func Day8Task1() {
	bytes, _ := os.ReadFile("./8.txt")
	text := string(bytes)
	lines := strings.Split(text, "\n")
	lines = util.Map(lines, strings.TrimSpace)

	cells := []cell{}

	for y, l := range lines {

		for x, c := range l {
			cells = append(cells, cell{pos: vec{x, y}, freq: c, anti: false})
		}
	}

	w := world{size: vec{x: len(lines[0]), y: len(lines)}, cells: cells}

	for _, cell1 := range w.cells {
		for _, cell2 := range w.cells {
			if cell1.pos.x == cell2.pos.x && cell1.pos.y == cell2.pos.y {
				continue
			}

			if cell1.freq == '.' || cell2.freq == '.' || cell1.freq != cell2.freq {
				continue
			}

			d := vec{x: cell2.pos.x - cell1.pos.x, y: cell2.pos.y - cell1.pos.y}

			p1 := vec{cell1.pos.x - d.x, cell1.pos.y - d.y}
			p2 := vec{cell2.pos.x + d.x, cell2.pos.y + d.y}

			if w.inBounds(p1) {
				w.getCell(p1).anti = true
			}

			if w.inBounds(p2) {
				w.getCell(p2).anti = true
			}
		}
	}

	fmt.Println(w)

	count := 0
	for _, cell := range w.cells {
		if cell.anti {
			count++
		}
	}

	fmt.Println(count)
}

func Day8Task2() {
	bytes, _ := os.ReadFile("./8.txt")
	text := string(bytes)
	lines := strings.Split(text, "\n")
	lines = util.Map(lines, strings.TrimSpace)

	cells := []cell{}

	for y, l := range lines {

		for x, c := range l {
			cells = append(cells, cell{pos: vec{x, y}, freq: c, anti: false})
		}
	}

	w := world{size: vec{x: len(lines[0]), y: len(lines)}, cells: cells}

	for _, cell1 := range w.cells {
		for _, cell2 := range w.cells {
			if cell1.pos.x == cell2.pos.x && cell1.pos.y == cell2.pos.y {
				continue
			}

			if cell1.freq == '.' || cell2.freq == '.' || cell1.freq != cell2.freq {
				continue
			}

			d := vec{x: cell2.pos.x - cell1.pos.x, y: cell2.pos.y - cell1.pos.y}

			p1 := vec{cell1.pos.x, cell1.pos.y}
			p2 := vec{cell2.pos.x, cell2.pos.y}

			i1 := 0
			for w.inBounds(p1) {
				w.getCell(p1).anti = true
				i1++
				p1 = vec{cell1.pos.x - i1*d.x, cell1.pos.y - i1*d.y}
			}

			i2 := 0
			for w.inBounds(p2) {
				w.getCell(p2).anti = true
				i2++
				p2 = vec{cell1.pos.x + i2*d.x, cell1.pos.y + i2*d.y}
			}
		}
	}

	fmt.Println(w)

	count := 0
	for _, cell := range w.cells {
		if cell.anti {
			count++
		}
	}

	fmt.Println(count)
}
