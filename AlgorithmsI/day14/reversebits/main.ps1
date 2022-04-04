function reverseBits {
    param ([uint32]$num)
    [uint32]$result = 0
    [uint32]$power = 31
    while ($num -ne 0) {
        $result += ($num -band 1) -shl $power
        $power--
        $num = $num -shr 1
    }
    $result
}

