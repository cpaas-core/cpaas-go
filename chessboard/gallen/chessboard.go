package chessboard

const rankSize = 8

type Rank [rankSize]bool

type Chessboard map[string]Rank

func CountInRank(cb Chessboard, rank string) (count int) {
    for _, occupied := range cb[rank] {
        if (occupied) {
            count++
        }
    }
    return
}

func CountInFile(cb Chessboard, file int) (count int) {
    if (file < 0 || file > rankSize) {
       return 0
    }

    for _, rank := range cb {
        if (rank[file-1]) {
            count++
        }
    }
    return
}

func CountAll(cb Chessboard) int {
    return len(cb) * rankSize
}

func CountOccupied(cb Chessboard) (count int) {
    for rank, _ := range cb {
        count += CountInRank(cb, rank)
    }
    return
}
