# 88. Merge Sorted Array
# https://leetcode.com/problems/merge-sorted-array/

# time: O(m+n)
# space: O(1)
function merge {
    param ([int[]][ref]$nums1, [int]$m, [int[]]$nums2, [int]$n)
    $size = $m + $n
    $i = $m -1
    $j = $n -1
    while ($size -gt 0) {
        $size--
        if ($j -ge 0 -and $i -ge 0 -and $nums1[$i] -gt $nums2[$j]) {
            $nums1[$size] = $nums1[$i]
            $i--
            continue
        }
        if ($j -ge 0 -and $i -ge 0 -and $nums2[$j] -gt $nums1[$i]) {
            $nums1[$size] = $nums2[$j]
            $j--
            continue
        }
        if ($i -ge 0) {
            $nums1[$size] = $nums1[$i]
            $i--
            continue
        }
        $nums1[$size] = $nums2[$j]
        $j--
    }
}

[int[]]$nums = @(1, 2, 3, 0, 0, 0)
merge -nums1 ([ref]$nums) -m 3 -nums2 @(2,5,6) -n 3
$nums -join ', '

[int[]]$nums = [int[]]@(1)
merge -nums1 ([ref]$nums) -m 1 -nums2 @() -n 0
$nums -join ', '

[int[]]$nums = @(0)
merge -nums1 ([ref]$nums) -m 0 -nums2 @(1) -n 1
$nums -join ', '
