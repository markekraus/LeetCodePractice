# LeetCode Practice

This is a collection of my attempts at solutions to LeetCode problems.
Please note that these may not be the best or optimal solutions.
Some of the solutions may be incomplete and/or not working.
Some of the solutions are missing.
I will try to get these all competed and updated to working order, but PRs are welcome!

Most of the solutions are available in Go.
Some also have PowerShell (PS) solutions.

PowerShell examples may not follow PowerShell best practices.
These solutions are intended to be done quickly in a coding interview.
PowerShell is a very verbose language when done properly.
As such, the PowerShell solutions optimize function over form.

## To Do

* Cleanup solutions
* Document time and space complexity for each solution
* Add PowerShell solutions
* Document the LeetCode problem title and URL in each solution
* Document how each solutions works

## Solutions

* Algorithms I [https://leetcode.com/study-plan/algorithm/](https://leetcode.com/study-plan/algorithm/)
  * Day 1 - Binary Search
    * [704. Binary Search](https://leetcode.com/problems/binary-search/) - [Go](AlgorithmsI/day01/binarysearch/main.go)
    * [278. First Bad Version](https://leetcode.com/problems/first-bad-version/) - [Go](AlgorithmsI/day01/firstbadversion/main.go)
    * [35. Search Insert Position](https://leetcode.com/problems/search-insert-position/) - [Go](AlgorithmsI/day01/searchinsertposition/main.go)
  * Day 2 - Two Pointers
    * [977. Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/) - [Go](AlgorithmsI/day02/squaresofasortedarray/main.go)
    * [189. Rotate Array](https://leetcode.com/problems/rotate-array/) - [Go](AlgorithmsI/day02/rotatearray/main.go)
  * Day 3 - Two Pointers
    * [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/) - [Go](AlgorithmsI/day03/movezeroes/main.go)
    * [167. Two Sum II - Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) - [Go](AlgorithmsI/day03/twosumiiarrayissorted/main.go)
  * Day 4 - Two Pointers
    * [344. Reverse String](https://leetcode.com/problems/reverse-string/) - [Go](AlgorithmsI/day04/reversestring/main.go)
    * [557. Reverse Words in a String III](https://leetcode.com/problems/reverse-words-in-a-string-iii/) - [Go](AlgorithmsI/day04/reversewordsinstringiii/main.go)
  * Day 5 - Two Pointers
    * [876. Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/) - [Go](AlgorithmsI/day05/middleofthelinkedlist/main.go)
    * [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) - [Go](AlgorithmsI/day05/removenthnodefromendoflist/main.go)
  * Day 6 - Sliding Window
    * [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) - [Go](AlgorithmsI/day06/lengthOfLongestSubstring/main.go)
    * [567. Permutation in String](https://leetcode.com/problems/permutation-in-string/) - [Go](AlgorithmsI/day06/permutationinstring/main.go)
  * Day 7 - Breadth-First Search / Depth-First Search
    * [733. Flood Fill](https://leetcode.com/problems/flood-fill/) - [Go](AlgorithmsI/day07/floodfill/main.go)
    * [695. Max Area of Island](https://leetcode.com/problems/max-area-of-island/) - [Go](AlgorithmsI/day07/maxareaofisland/main.go)
  * Day 8 - Breadth-First Search / Depth-First Search
    * [617. Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/) - [Go](AlgorithmsI/day08/mergetwobinarytrees/main.go)
    * [116. Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/) - [Go](AlgorithmsI/day08/populatingnextrightpointersineachnode/main.go)
  * Day 9 - Breadth-First Search / Depth-First Search
    * [542. 01 Matrix](https://leetcode.com/problems/01-matrix/) - [Go](AlgorithmsI/day09/01matrix/main.go)
    * [994. Rotting Oranges](https://leetcode.com/problems/rotting-oranges/) - [Go](AlgorithmsI/day09/rottingoranges/main.go)
  * Day 10 - Recursion / Backtracking
    * [21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) - [Go](AlgorithmsI/day10/mergetwosortedlists/main.go)
    * [206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/) - [Go](AlgorithmsI/day10/reverselinkedlist/main.go)
  * Day 11 - Recursion / Backtracking
    * [77. Combinations](https://leetcode.com/problems/combinations/) - [Go](AlgorithmsI/day11/combinations/main.go)
    * [46. Permutations](https://leetcode.com/problems/permutations/) - [Go](AlgorithmsI/day11/permutations/main.go)
    * [784. Letter Case Permutation](https://leetcode.com/problems/letter-case-permutation/) - [Go](AlgorithmsI/day11/lettercasepermutation/main.go)
  * Day 12 - Dynamic Programming
    * [70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) - [Go](AlgorithmsI/day12/climbingstairs/main.go)
    * [198. House Robber](https://leetcode.com/problems/house-robber/) - [Go](AlgorithmsI/day12/houserobber/main.go)
    * [120. Triangle](https://leetcode.com/problems/triangle/) - [Go](AlgorithmsI/day12/triangle/main.go)
  * Day 13 - Bit Manipulation
    * [231. Power of Two](https://leetcode.com/problems/power-of-two/) - [Go](AlgorithmsI/day13/poweroftwo/main.go)
    * [191. Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/) - [Go](AlgorithmsI/day13/numberofonebits/main.go)
  * Day 14 - Bit Manipulation
    * [190. Reverse Bits](https://leetcode.com/problems/reverse-bits/) - [Go](AlgorithmsI/day14/reversebits/main.go) | [PS](AlgorithmsI/day14/reversebits/main.ps1)
    * [136. Single Number](https://leetcode.com/problems/single-number/) - [Go](AlgorithmsI/day14/singlenum/main.go) | [PS](AlgorithmsI/day14/singlenum/main.ps1)
* Data Structure I
  * Day 1 - Array
    * [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) - [Go](DataStructureI/day01/containsduplicate/main.go) | [PS](DataStructureI/day01/containsduplicate/main.ps1)
    * [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) - [Go](DataStructureI/day01/maximumsubarray/main.go) | [PS](DataStructureI/day01/maximumsubarray/main.ps1)
  * Day 2 - Array
    * [1. Two Sum](https://leetcode.com/problems/two-sum/) - [Go](DataStructureI/day02/twosum/main.go) | [PS](DataStructureI/day02/twosum/main.ps1)
    * [88. Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/) - [Go](DataStructureI/day02/mergesortedarray/main.go) | [PS](DataStructureI/day02/mergesortedarray/main.ps1)
  * Day 3 - Array
    * [350. Intersection of Two Arrays II](https://leetcode.com/problems/intersection-of-two-arrays-ii/) - [Go](DataStructureI/day03/interesctionoftwoarraysii/main.go) | [PS](DataStructureI/day03/interesctionoftwoarraysii/main.ps1)
    * [121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) - [Go](DataStructureI/day03/besttimetobuyandsellstock/main.go) | [PS](DataStructureI/day03/besttimetobuyandsellstock/main.ps1)
  * Day 4 - Array
    * [566. Reshape the Matrix](https://leetcode.com/problems/reshape-the-matrix/) - [Go](DataStructureI/day04/reshapematrix/main.go) | [PS](DataStructureI/day04/reshapematrix/main.ps1)
    * [118. Pascal's Triangle](https://leetcode.com/problems/pascals-triangle/) - [Go](DataStructureI/day04/pascalstriangle/main.go) | [PS](DataStructureI/day04/pascalstriangle/main.ps1)
