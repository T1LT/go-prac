package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
    
    for _, v := range cb[file] {
        if v {
            count++
        }
    }

    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
    
    if rank < 1 || rank > 8 {
        return count
    }

	for k, _ := range cb {
        for i, v := range cb[k] {
            if i == rank - 1 && v {
                count++
            }
        }
    }

    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0

    for _, v := range cb {
        for _, _ = range v {
            count++
        }
    }

    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0

    for _, v := range cb {
        for _, cell := range v {
            if cell {
                count++
            }
        }
    }

    return count
}
