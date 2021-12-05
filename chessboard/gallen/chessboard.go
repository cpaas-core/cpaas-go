package chessboard

const RankSize = 8

type Rank [RankSize]bool

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
    if (file <= RankSize) {
        for _, rank := range cb {
            if (rank[file-1]) {
                count++
            }
        }
    }
    return
}

func CountAll(cb Chessboard) int {
    for _, rank := range cb {
        return len(cb) * len(rank)
    }
    return 0
}

func CountOccupied(cb Chessboard) (count int) {
    for _, rank := range cb {
        for _, occupied := range rank {
            if (occupied) {
                count++
            }
	}
    }
    return
}
