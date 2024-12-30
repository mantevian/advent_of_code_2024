package main

import (
	"os"
	"strings"
)

type vec6 struct {
	x int
	y int
}

type cell6 struct {
	pos      vec6
	obstacle bool
	visited  bool
}

type world6 struct {
	cells     []cell6
	size      vec6
	guard     *vec6
	direction *vec6
}

func (w *world6) getCell(pos vec6) *cell6 {
	return &w.cells[pos.y*w.size.x+pos.x]
}

func newWorld(input string) world6 {
	lines := strings.Split(input, "\n")

	cells := []cell6{}

	var g vec6

	for y, line := range lines {
		for x, c := range line {
			cells = append(cells, cell6{vec6{x, y}, c == '#', false})
			if c == '^' {
				g = vec6{x, y}
			}
		}
	}

	return world6{cells: cells, size: vec6{len(lines[0]), len(lines)}, guard: &g, direction: &vec6{0, -1}}
}

func (v *vec6) turnRight() {
	if v.x == 0 && v.y == -1 {
		v.x = 1
		v.y = 0
		return
	}

	if v.x == 1 && v.y == 0 {
		v.x = 0
		v.y = 1
		return
	}

	if v.x == 0 && v.y == 1 {
		v.x = -1
		v.y = 0
		return
	}

	if v.x == -1 && v.y == 0 {
		v.x = 0
		v.y = -1
		return
	}
}

func (w world6) inBounds(cell vec6) bool {
	return cell.x >= 0 && cell.x < w.size.x && cell.y >= 0 && cell.y < w.size.y
}

// 0 -1 -> 1 0 -> 0 1 -> -1 0

func Day6Task1() {
	bytes, _ := os.ReadFile("./6.txt")

	w := newWorld(string(bytes))

	for w.inBounds(*w.guard) {
		w.getCell(*w.guard).visited = true

		forward := vec6{w.guard.x + w.direction.x, w.guard.y + w.direction.y}
		if w.inBounds(forward) && w.getCell(forward).obstacle {
			w.direction.turnRight()
		}

		w.guard.x += w.direction.x
		w.guard.y += w.direction.y
	}

	count := 0
	for _, c := range w.cells {
		if c.visited {
			count++
		}
	}

	println(count)
}
