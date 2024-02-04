package main

type GameStage struct {
	width, height int
	holdingTetrimino *Tetrimino
	operatingTetrimino *Tetrimino
	field [][]int
}


type TetriminoType int

type Tetrimino struct {
	x, y int
	tetriminoType TetriminoType
	rotation int
	shapes [4][][]int
}