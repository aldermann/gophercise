package game

import (
    "errors"
    "fmt"
    "math/rand"
)

func (b Board) MarkMove(x, y int, t uint8) error {
    if !b.IsValidCoordinate(x, y) {
        return fmt.Errorf("(%d,%d) is not a valid coordinate", x, y)
    }
    if b.container[y][x] != Empty {
        return errors.New("a move have been made on that square")
    }
    b.container[y][x] = t
    return nil
}

func (b Board) UndoMove(x, y int) error {
    if !b.IsValidCoordinate(x, y) {
        return fmt.Errorf("(%d,%d) is not a valid coordinate", x, y)
    }
    if b.container[y][x] == Empty {
        return errors.New("the cell is already empty")
    }
    b.container[y][x] = Empty
    return nil
}

type pair struct {
    x, y int
}

func (b Board) MoveGen(player uint8) <-chan Board {
    c := make(chan Board)
    cellList := make([]pair, 0, b.Col*b.Row)
    for y, row := range b.container {
        for x, cell := range row {
            if cell == Empty {
                cellList = append(cellList, pair{x, y})
            }
        }
    }
    rand.Shuffle(len(cellList), func(i, j int) { cellList[i], cellList[j] = cellList[j], cellList[i] })
    go func() {
        for _, cell := range cellList{
            newB := b.Clone()
            _ = newB.MarkMove(cell.x, cell.y, player)
            c <- newB
        }
        close(c)
    }()
    return c
}
