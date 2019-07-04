package algorithm

import (
    "tictactoe/game"
)

func CalculateBestMove(board game.Board, alpha, beta, depth int, maximizingPlayer bool, rootCall bool) (int, game.Board) {
    val := board.GetValue()
    if val != 0 || board.IsFull() || depth == 0 {
        return val, board
    }
    resVal, resMove := 0, game.Board{}
    if maximizingPlayer {
        resVal = -1000000
        for nextMove := range board.MoveGen(game.O) {
            v, _ := CalculateBestMove(nextMove, alpha, beta, depth-1, false, false)
            if v >= resVal {
                resVal = v
                resMove = nextMove
            }
            if resVal > alpha {
                alpha = resVal
            }
            if alpha >= beta {
                break
            }
        }
    } else {
        resVal = 1000000
        for nextMove := range board.MoveGen(game.X) {
            v, _ := CalculateBestMove(nextMove, alpha, beta, depth-1, true, false)
            if v <= resVal {
                resVal = v
                resMove = nextMove
            }
            if resVal < beta {
                beta = resVal
            }
            if alpha >= beta {
                break
            }
        }
    }
    if rootCall {
        return resVal, resMove
    } else {
        return resVal, game.Board{}
    }
}
