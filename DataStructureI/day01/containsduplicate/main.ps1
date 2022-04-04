# 217. Contains Duplicate
# https://leetcode.com/problems/contains-duplicate/

# time worst: O(n*log(n))
# time average: O(n)
# space: O(n)
function containsDuplicate {
    param ([int[]]$nums)
    if ($nums.Count -eq 1) {
        return $false
    }
    $visited = @{}
    $visited[$nums[0]] = $true
    for ($i = 1; $i -lt $nums.Count; $i++) {
        if ($visited[$nums[$i]]) {
            return $true
        }
        $visited[$nums[$i]] = $true
    }
    $false
}

containsDuplicate -nums @(1, 2, 3, 1)
containsDuplicate -nums @(1, 2, 3, 4)
containsDuplicate -nums @(1, 1, 1, 3, 3, 4, 3, 2, 4, 2)