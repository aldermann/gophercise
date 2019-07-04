package game

func (b Board) checkRow () int {
    for _, row := range b.container {
        allX, allO := true, true
        for _, cell := range row {
            if cell != X {
                allX = false
            }
            if cell != O {
                allO = false
            }
        }
        if allX {
            return -1
        }
        if allO {
            return 1
        }
    }
    return 0
}

func (b Board) checkColumn () int {
    for j := 0; j < b.Col; j++ {
        allX, allO := true, true
        for i := 0; i < b.Row; i++ {
            if b.container[i][j] != X {
                allX = false
            }
            if b.container[i][j] != O {
                allO = false
            }
        }
        if allX {
            return -1
        }
        if allO {
            return 1
        }
    }
    return 0
}

func (b Board) checkDiagonal () int {
    if b.Row != b.Col {
        return 0
    }
    allX, allO := true, true
    for i := 0; i < b.Row; i++ {
        if b.container[i][i] != X {
            allX = false
        }
        if b.container[i][i] != O {
            allO = false
        }
    }
    if allX {
        return -1
    }
    if allO {
        return 1
    }

    allX, allO = true, true
    for i := 0; i < b.Row; i++ {
        if b.container[i][b.Row - i - 1] != X {
            allX = false
        }
        if b.container[i][b.Row - i - 1] != O {
            allO = false
        }
    }

    if allX {
        return -1
    }
    if allO {
        return 1
    }

    return 0
}

func (b Board) GetValue () int {
    c := b.checkColumn()
    r := b.checkRow()
    d := b.checkDiagonal()
    if c != 0 {
        return c
    }
    if r != 0 {
        return r
    }
    if d != 0 {
        return d
    }
    return 0
}

func (b Board) IsFull () bool {
    for _, row := range b.container {
        for _, cell := range row {
            if cell == Empty {
                return false
            }
        }
    }
    return true
}

