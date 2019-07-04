package game

const (
    Empty = iota
    X = iota
    O = iota
)

type Board struct {
    container [][]uint8
    Row, Col  int
}

func MakeBoard (row, col int) Board{
    c := make([][]uint8, row)
    for i := range c {
        c[i] = make([]uint8, col)
        for j := range c[i] {
            c[i][j] = Empty
        }
    }
    return Board{container: c, Row: row, Col: col}
}

func (b Board) Clone () Board {
    newBoard := MakeBoard(b.Row, b.Col)
    for i, row := range b.container {
        for j, cell := range row {
            _ = newBoard.MarkMove(j, i, cell)
        }
    }
    return newBoard
}

func (b Board) IsValidCoordinate (x, y int) bool {
    if x < 0 || x >= b.Col {
        return false
    }
    if y < 0 || y >= b.Row {
        return false
    }
    return true
}