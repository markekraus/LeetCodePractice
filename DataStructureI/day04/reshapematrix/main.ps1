# 566. Reshape the Matrix
# https://leetcode.com/problems/reshape-the-matrix/

# time: O(n*m)
# space: O(1)
function matrixReshape {
    param ([int[][]]$mat, [int]$r, [int]$c)
    if ($r * $c -ne $mat.Length * $mat[0].Length) {
        return $mat
    }
    $result = [int[][]]::new($r)
    $cr = 0
    $cc = 0
    for ($i = 0; $i -lt $r; $i++) {
        $result[$i] = [int[]]::new($c)
        for ($j = 0; $j -lt $c; $j++) {
            if ($cc -ge $mat[0].Length) {
                $cr++
                $cc = 0
            }
            $result[$i][$j] = $mat[$cr][$cc]
            $cc++
        }
    }
    $result
}

$matrix = [int[][]]::new(2)
$matrix[0] = [int[]]@(1,2)
$matrix[1] = [int[]]@(3,4)
matrixReshape -mat $matrix -r 1 -c 4 | ConvertTo-Json -Compress
matrixReshape -mat $matrix -r 2 -c 4 | ConvertTo-Json -Compress