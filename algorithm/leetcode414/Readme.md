# Third Maximum Number

Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).

Example 1:
Input: [3, 2, 1]

Output: 1

Explanation: The third maximum is 1.

Example 2:
Input: [1, 2]

Output: 2

Explanation: The third maximum does not exist, so the maximum (2) is returned instead.

Example 3:
Input: [2, 2, 3, 1]

Output: 1

Explanation: Note that the third maximum here means the third maximum distinct number.
Both numbers with value 2 are both considered as second maximum.

这道题给初学者来说就是使用一个三个元素的数组就能解决，但是，为了更好的体现计算机的一些算法应用，这道题目其实是求中位数算法的一个变体。

快速排序的划分方法就是一个非常好用的划分算法，中位数算法就是依赖这个算法进行划分的。
