# 1. Two Sum
# https://leetcode.com/problems/two-sum/

# time: O(n)
# space: O(n)
function twoSum {
    param ([int[]]$nums, [int]$target)
    if ($nums.Length -eq 2) {
        return @(0,1)
    }
    $visited = @{}
    for ($i = 0; $i -lt $nums.Length; $i++) {
        if ($visited.ContainsKey($nums[$i])) {
            return @($visited[$nums[$i]], $i)
        }
        $visited[$target - $nums[$i]] = $i
    }
    @()
}

(twoSum -nums @(2, 7, 11, 15) -target 9) -join ', '
(twoSum -nums @(3, 2, 4) -target 6)  -join ', '
(twoSum -nums @(3, 3) -target 6)  -join ', '