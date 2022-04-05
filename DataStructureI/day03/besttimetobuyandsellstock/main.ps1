# 121. Best Time to Buy and Sell Stock
# https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

# Time: O(n)
# Space: O(1)
function maxProfit {
    param ([int[]]$prices)
    $total, $local = 0, 0
    for ($i = 0; $i -lt $prices.Length -1; $i++) {
        $local = [math]::Max(0, $local + $prices[$i+1] - $prices[$i])
        $total = [math]::Max($local, $total)
    }
    $total
}

maxProfit -prices @(7, 1, 5, 3, 6, 4) # 5
maxProfit -prices @(7, 6, 4, 3, 1)    # 0