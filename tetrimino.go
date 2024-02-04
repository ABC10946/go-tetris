package main

import (
	"fmt"
)

// functions for tetrimino operation

func (g *GameStage) MoveLeft() error {
	if g.operatingTetrimino.x < 1 {
		return fmt.Errorf("cannot move left")
	}
	g.operatingTetrimino.x--
	return nil
}

func (g *GameStage) MoveRight() error {
	if g.operatingTetrimino.x > g.width - 5 {
		return fmt.Errorf("cannot move right")
	}
	g.operatingTetrimino.x++
	return nil
}

func (g *GameStage) init(width int, height int) {
	g.width = width
	g.height = height
	g.field = make([][]int, height)
	for i := range g.field {
		g.field[i] = make([]int, width)
	}

	// fill field with 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			g.field[i][j] = 0
		}
	}

	// create a tetrimino
	g.holdingTetrimino = new(Tetrimino)
	g.holdingTetrimino.MakeTetrimino(0)

	g.operatingTetrimino = new(Tetrimino)
	g.operatingTetrimino.MakeTetrimino(1)
}

func (g *GameStage) clearField() {
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			g.field[i][j] = 0
		}
	}
}

func (g *GameStage) putTetrimino() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.operatingTetrimino.shapes[g.operatingTetrimino.rotation][i][j] == 1 {
				g.field[g.operatingTetrimino.y + i][g.operatingTetrimino.x + j] = 1
			}
		}
	}
}

func (g *GameStage) removeTetrimino(t *Tetrimino) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.operatingTetrimino.shapes[g.operatingTetrimino.rotation][i][j] == 1 {
				g.field[t.y + i][t.x + j] = 0
			}
		}
	}
}

func (t *Tetrimino) Rotate() {
	t.rotation = (t.rotation + 1) % 4
}

func (t *Tetrimino) MakeTetrimino(tetriminoType TetriminoType) {
	t.x = 4
	t.y = 0
	t.tetriminoType = tetriminoType
	t.rotation = 0

	if tetriminoType == 0 {
		// I
		t.shapes[0] = [][]int{
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		t.shapes[1] = [][]int{
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 1, 0},
		}
		t.shapes[2] = [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
		}
		t.shapes[3] = [][]int{
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
		}
	}
	if tetriminoType == 1 {
		// J
		t.shapes[0] = [][]int{
			{1, 0, 0, 0},
			{1, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		t.shapes[1] = [][]int{
			{0, 1, 1, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		}
		t.shapes[2] = [][]int{
			{0, 0, 0, 0},
			{1, 1, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 0},
		}
		t.shapes[3] = [][]int{
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0},
			{0, 0, 0, 0},
		}
	}
	if tetriminoType == 2 {
		// L
		t.shapes[0] = [][]int{
			{0, 0, 1, 0},
			{1, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		t.shapes[1] = [][]int{
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		}
		t.shapes[2] = [][]int{
			{0, 0, 0, 0},
			{1, 1, 1, 0},
			{1, 0, 0, 0},
			{0, 0, 0, 0},
		}
		t.shapes[3] = [][]int{
			{1, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		}
	}
	if tetriminoType == 3 {
		// O
		t.shapes[0] = [][]int{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		}
		t.shapes[1] = t.shapes[0]
		t.shapes[2] = t.shapes[0]
		t.shapes[3] = t.shapes[0]
	}
	if tetriminoType == 4 {
		// S
		t.shapes[0] = [][]int{
			{0, 1, 1, 0},
			{1, 1, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		t.shapes[1] = [][]int{
			{0, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 0},
		}
		t.shapes[2] = t.shapes[0]
		t.shapes[3] = t.shapes[1]
	}
	if tetriminoType == 5 {
		// T
		t.shapes[0] = [][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		t.shapes[1] = [][]int{
			{0, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		}
		t.shapes[2] = [][]int{
			{0, 0, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		}
		t.shapes[3] = [][]int{
			{0, 1, 0, 0},
			{1, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		}
	}
	if tetriminoType == 6 {
		// Z
		t.shapes[0] = [][]int{
			{1, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		t.shapes[1] = [][]int{
			{0, 0, 1, 0},
			{0, 1, 1, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		}
		t.shapes[2] = t.shapes[0]
		t.shapes[3] = t.shapes[1]
	}
}
