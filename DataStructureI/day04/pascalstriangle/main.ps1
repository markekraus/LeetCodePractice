# 118. Pascal's Triangle
# https://leetcode.com/problems/pascals-triangle/

# time: O(n2)
# space: O(1)
function generate {
    param ([int]$numRows)
    $result = [int[][]]::new($numRows)
    for ($row = 0; $row -lt $numRows; $row++) {
        $result[$row] = [int[]]::new($row + 1)
        $result[$row][0] = 1
        $result[$row][$row] = 1
        for ($col = 1; $row -gt 1 -and $col -lt $row; $col++) {
            $result[$row][$col] = $result[$row - 1][$col -1] + $result[$row - 1][$col]
        }
    }
    $result
}

generate -numRows 5 | ConvertTo-Json -Compress
generate -numRows 1 | ConvertTo-Json -Compress