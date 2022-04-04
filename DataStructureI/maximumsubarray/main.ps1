# 53. Maximum Subarray
# https://leetcode.com/problems/maximum-subarray/

# time: O(n)
# space: O(n)
# Dynamic Programing solution
function maxSubArray {
    param ([int[]]$nums)
    if ($nums.Length -eq 1) {
        return $nums[0]
    }
    $table = [int[]]::new($nums.Length)
    $table[0] = $nums[0]
    for ($i = 1; $i -lt $nums.Length; $i++) {
        $table[$i] = [math]::Max($nums[$i]+$table[$i-1],$nums[$i])
    }
    for ($i = 1; $i -lt $nums.Length; $i++) {
        $table[$i] = [math]::Max($table[$i-1],$table[$i])
    }
    $table[$nums.Length - 1]
}

maxSubArray -nums @(-2, 1, -3, 4, -1, 2, 1, -5, 4)
maxSubArray -nums @(1)
maxSubArray -nums @(5, 4, -1, 7, 8)