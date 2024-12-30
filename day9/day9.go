package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day9Task1() {
	bytes, _ := os.ReadFile("./9.txt")
	text := strings.TrimSpace(string(bytes))
	blocks := []int{}

	for i, c := range text {
		l, _ := strconv.Atoi(string(c))

		if i%2 == 0 {
			for range l {
				blocks = append(blocks, i/2)
			}
		} else {
			for range l {
				blocks = append(blocks, -1)
			}
		}
	}

	p1 := 0
	p2 := len(blocks) - 1

	res := []int64{}

	for p1 != p2+1 {
		if blocks[p1] != -1 {
			res = append(res, int64(blocks[p1]))
			p1++
			continue
		}

		if blocks[p2] == -1 {
			p2--
			continue
		}

		res = append(res, int64(blocks[p2]))
		p2--
		p1++
	}

	var sum int64 = 0
	for i, n := range res {
		sum += int64(i) * n
	}

	fmt.Println(sum)
}

type block struct {
	start  int
	length int
}

func Day9Task2() {
	bytes, _ := os.ReadFile("./9.txt")
	text := strings.TrimSpace(string(bytes))
	blocks := []block{}
	spaces := []block{}

	p := 0
	for i, c := range text {
		l, _ := strconv.Atoi(string(c))

		if i%2 == 0 {
			blocks = append(blocks, block{start: p, length: l})
		} else {
			spaces = append(spaces, block{start: p, length: l})
		}

		p += l
	}

outer:
	for i := len(blocks) - 1; i >= 0; i-- {
		for j := range len(spaces) {
			if spaces[j].start > blocks[i].start {
				break
			}

			if spaces[j].length < blocks[i].length {
				continue
			}

			blocks[i].start = spaces[j].start
			spaces[j].length -= blocks[i].length
			spaces[j].start += blocks[i].length
			continue outer
		}
	}

	var sum int64 = 0
	for i, b := range blocks {
		for j := range b.length {
			sum += int64(i) * (int64(b.start) + int64(j))
		}
	}

	fmt.Println(sum)
}
