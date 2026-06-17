func isValidSudoku(board [][]byte) bool {
    rows := [9][9]bool{}
    cols := [9][9]bool{}
    boxes := [9][9]bool{}

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if board[r][c] == '.' {
                continue
            }

            num := board[r][c] - '1'  // '1'-'9' → 0-8
            box := (r/3)*3 + c/3      // which of 9 boxes

            if rows[r][num] || cols[c][num] || boxes[box][num] {
                return false
            }

            rows[r][num] = true
            cols[c][num] = true
            boxes[box][num] = true
        }
    }
    return true
}
