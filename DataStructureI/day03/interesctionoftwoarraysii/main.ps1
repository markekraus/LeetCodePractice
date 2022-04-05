# 350. Intersection of Two Arrays II
# https://leetcode.com/problems/intersection-of-two-arrays-ii/

# time: O(n+m)
# space: O(n)
function intersect {
    param ([int[]]$nums1, [int[]]$nums2)
    $lookup = @{}
    foreach ($num in $nums1) {
        $lookup[$num] += 1
    }
    foreach ($num in $nums2) {
        $temp = $lookup[$num]
        if ($temp -gt 0) {
            $lookup[$num]--
            $num
        }
    }
}

(intersect -nums1 @(1, 2, 2, 1) -nums2 @(2,2)) -join ', '
(intersect -nums1 @(4, 9, 5) -nums2 @(9,4)) -join ', '
