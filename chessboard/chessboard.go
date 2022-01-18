package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools

type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"

type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	// panic("Please implement CountInRank()")
	var count int

	for _,hasValue := range cb[rank] {
		if hasValue {
			count += 1
		}
	}

	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	// panic("Please implement CountInFile()")

	var count int

	if file < 1 || file > 8 {
		return 0
	}

	fileIndex := file - 1

	for _,col := range cb {
		for rowIndex, row := range col {
			if rowIndex != fileIndex || !row {
				continue
			}

			count += 1
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	// panic("Please implement CountAll()")

	var count int

	for _,col := range cb {
		for rowIndex := 0; rowIndex < len(col); rowIndex++ {
			count += 1
		}
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	// panic("Please implement CountOccupied()")

	var count int

	for _,col := range cb {
		for _, row := range col {
			if row {
				count += 1
			}
		}
	}

	return count
}
