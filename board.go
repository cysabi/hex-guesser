package main

import (
	"strings"
)

type Cell struct {
	Empty string
	Wall  string
}

type board struct {
	cells      []string
	boardWidth int
}

func (c *board) init(w, h int) {
	if w == 0 {
		return
	}
	c.boardWidth = w
	c.cells = make([]string, w*h)
	c.wipe()
}

func (c board) get(x, y int) string {
	i := y*c.boardWidth + x
	if i > len(c.cells)-1 || x < 0 || y < 0 || x >= c.width() || y >= c.height() {
		return " "
	}
	return c.cells[i]
}

func (c *board) set(x, y int, v string) {
	i := y*c.boardWidth + x
	if i > len(c.cells)-1 || x < 0 || y < 0 || x >= c.width() || y >= c.height() {
		return
	}
	c.cells[i] = v
}

func (c *board) wipe() {
	for i := range c.cells {
		c.cells[i] = " "
	}
}

func (c board) width() int {
	return c.boardWidth
}

func (c board) height() int {
	h := len(c.cells) / c.boardWidth
	if len(c.cells)%c.boardWidth != 0 {
		h++
	}
	return h
}

func (c board) ready() bool {
	return len(c.cells) > 0
}

func (c board) String() string {
	var b strings.Builder
	for i := 0; i < len(c.cells); i++ {
		b.WriteString(c.cells[i])
	}
	return b.String()
}
