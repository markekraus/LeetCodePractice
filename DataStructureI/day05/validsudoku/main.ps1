# 36. Valid Sudoku
# https://leetcode.com/problems/valid-sudoku/

# time: O(n)
# space: O(n)
function isValidSudoku {
    param ($board)
    $row = [bool[][]]::new(9)
    $col = [bool[][]]::new(9)
    $box = [bool[][]]::new(9)
    foreach ($i in 0..8) {
        $row[$i] = [bool[]]::new(9)
        $col[$i] = [bool[]]::new(9)
        $box[$i] = [bool[]]::new(9)
    }
    foreach ($i in 0..8) {
        foreach ($j in 0..8) {
            if ($board[$i][$j] -eq [char]'.'){continue}
            $boxId = [math]::Truncate($i/3)*3+[math]::Truncate($j/3)
            [int]$num = $board[$i][$j] - [char]'0' - 1
            if ($row[$i][$num] -or $col[$j][$num] -or $box[$boxId][$num]) {
                return $false
            }
            $row[$i][$num], $col[$j][$num], $box[$boxId][$num] = $true, $true, $true
        }
    }
    $true
}

[char[][]]$board = [char[][]]::new(9)
$board[0] = [char[]]@('5', '3', '.', '.', '7', '.', '.', '.', '.')
$board[1] = [char[]]@('6', '.', '.', '1', '9', '5', '.', '.', '.')
$board[2] = [char[]]@('.', '9', '8', '.', '.', '.', '.', '6', '.')
$board[3] = [char[]]@('8', '.', '.', '.', '6', '.', '.', '.', '3')
$board[4] = [char[]]@('4', '.', '.', '8', '.', '3', '.', '.', '1')
$board[5] = [char[]]@('7', '.', '.', '.', '2', '.', '.', '.', '6')
$board[6] = [char[]]@('.', '6', '.', '.', '.', '.', '2', '8', '.')
$board[7] = [char[]]@('.', '.', '.', '4', '1', '9', '.', '.', '5')
$board[8] = [char[]]@('.', '.', '.', '.', '8', '.', '.', '7', '9')
isValidSudoku -board $board # true

[char[][]]$board = [char[][]]::new(9)
$board[0] = [char[]]@('8', '3', '.', '.', '7', '.', '.', '.', '.')
$board[1] = [char[]]@('6', '.', '.', '1', '9', '5', '.', '.', '.')
$board[2] = [char[]]@('.', '9', '8', '.', '.', '.', '.', '6', '.')
$board[3] = [char[]]@('8', '.', '.', '.', '6', '.', '.', '.', '3')
$board[4] = [char[]]@('4', '.', '.', '8', '.', '3', '.', '.', '1')
$board[5] = [char[]]@('7', '.', '.', '.', '2', '.', '.', '.', '6')
$board[6] = [char[]]@('.', '6', '.', '.', '.', '.', '2', '8', '.')
$board[7] = [char[]]@('.', '.', '.', '4', '1', '9', '.', '.', '5')
$board[8] = [char[]]@('.', '.', '.', '.', '8', '.', '.', '7', '9')
isValidSudoku -board $board # false
