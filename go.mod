module abc10946/go-tetris

go 1.20

require github.com/hajimehoshi/ebiten/v2 v2.6.5

replace github.com/abc10946/go-tetris/tetrimino => ./tetrimino.go

replace github.com/abc10946/go-tetris/types => ./types.go

require (
	github.com/ebitengine/purego v0.5.0 // indirect
	github.com/jezek/xgb v1.1.0 // indirect
	golang.org/x/exp v0.0.0-20190731235908-ec7cb31e5a56 // indirect
	golang.org/x/image v0.12.0 // indirect
	golang.org/x/mobile v0.0.0-20230922142353-e2f452493d57 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
)
