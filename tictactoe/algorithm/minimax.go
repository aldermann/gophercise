package algorithm

import (
    "tictactoe/game"
)

const (
    Inf = int(1e8)
)
func CalculateBestMove(board game.Board, alpha, beta, depth int, maximizingPlayer bool, rootCall bool) (int, game.Board) {
    val := board.GetValue()
    if val != game.Draw || board.IsFull() || depth == 0 {
        return val, board
    }
    resVal, resMove := 0, game.Board{}
    if maximizingPlayer {
        resVal = -Inf
        for nextMove := range board.MoveGen(game.O) {
            v, _ := CalculateBestMove(nextMove, alpha, beta, depth-1, false, false)
            if v > resVal {
                resVal = v
                resMove = nextMove
            }
            if resVal > alpha {
                alpha = resVal
            }
            if alpha >= beta || resVal == game.Win {
                break
            }
        }
    } else {
        resVal = Inf
        for nextMove := range board.MoveGen(game.X) {
            v, _ := CalculateBestMove(nextMove, alpha, beta, depth-1, true, false)
            if v < resVal {
                resVal = v
                resMove = nextMove
            }
            if resVal < beta {
                beta = resVal
            }
            if alpha >= beta || resVal == game.Lose{
                break
            }
        }
    }
    if rootCall {
        return resVal, resMove.Clone()
    } else {
        return resVal, game.Board{}
    }
}
