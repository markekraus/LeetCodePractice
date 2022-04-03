function singleNumber {
    param ([int[]]$nums)
    for ($i = 1; $i -lt $nums.Count; $i++) {
        $nums[0] = $nums[0] -bxor $nums[$i]
    }
    $nums[0]
}

singleNumber -nums @(2,2,1)
singleNumber -nums @(4,1,2,1,2)