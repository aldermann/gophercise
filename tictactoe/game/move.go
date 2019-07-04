package game

import (
    "errors"
    "fmt"
)

func (b Board) MarkMove(x, y int, t uint8) error {
    if !b.IsValidCoordinate(x, y) {
        return fmt.Errorf("(%d,%d) is not a valid coordinate", x, y)
    }
    if b.container[y][x] != Empty{
        return errors.New("a move have been made on that square")
    }
    b.container[y][x] = t
    return nil
}

func (b Board) MoveGen(player uint8) <-chan Board {
    c := make(chan Board)
    go func() {
        for x, row := range b.container {
            for y, cell := range row {
                if cell == Empty {
                    newB := b.Clone()
                    _ = newB.MarkMove(y, x, player)
                    c <- newB
                }
            }
        }
        close(c)
    }()
    return c
}

