# 74. Search a 2D Matrix
# https://leetcode.com/problems/search-a-2d-matrix/

# time: O(1) to O(log n)
# space: O(1)
# Binary search 2-d array as a 1-d array
function searchMatrix {
    param ([int[][]]$matrix, [int]$target)
    $n = $matrix.Length
    $m = $matrix[0].Length
    if ($matrix[0][0] -gt $target -or $matrix[$n-1][$m-1] -lt $target) {
        return $false
    }
    $l = 0
    $r = $n*$m-1
    while ($l -le $r) {
        $mid = ($l + $r) -shr 1
        $row = [math]::Truncate($mid / $m)
        $col = $mid % $m
        $value = $matrix[$row][$col]
        if ($value -eq $target) {
            return $true
        }
        if ($value -gt $target) {
            $r = $mid - 1
            continue
        }
        $l = $mid + 1
    }
    $false
}

$matrix = [int[][]]::new(3)
$matrix[0] = [int[]]@(1, 3, 5, 7)
$matrix[1] = [int[]]@(10, 11, 16, 20)
$matrix[2] = [int[]]@(23, 30, 34, 60)

searchMatrix -matrix $matrix -target 3 # true
searchMatrix -matrix $matrix -target 13 # false