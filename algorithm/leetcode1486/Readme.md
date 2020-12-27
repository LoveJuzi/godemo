# XOR Operation in an Array

Given an integer n and an integer start.

Define an array nums where nums[i] = start + 2*i (0-indexed) and n == nums.length.

Return the bitwise XOR of all elements of nums.

Example 1:

Input: n = 5, start = 0
Output: 8
Explanation: Array nums is equal to [0, 2, 4, 6, 8] where (0 ^ 2 ^ 4 ^ 6 ^ 8) = 8.
Where "^" corresponds to bitwise XOR operator.

Example 2:

Input: n = 4, start = 3
Output: 8
Explanation: Array nums is equal to [3, 5, 7, 9] where (3 ^ 5 ^ 7 ^ 9) = 8.

Example 3:

Input: n = 1, start = 7
Output: 7

Example 4:

Input: n = 10, start = 5
Output: 2

Constraints:

1 <= n <= 1000
0 <= start <= 1000
n == nums.length

x = 0 ^ 2 = ?
x = x ^ 4 = ?
x = x ^ 6 = ?
x = x ^ 8 = ?

nums[i] = start + 2*i

0000  0
0010  2
----
0010
0100  4
----
0110
0110  6
----
0000
1000  8
----
1000

0011  3
0101  5
----
0110
0111  7
----
0001
1001  9
----
1000

a >> 1
a << 1

这种类型的题目，主要考察的是计算机的基本的算术运算符，比如左移和右移运算符，以及构造异或运算符。

以及异或本身的概念

1. 如何取最低位置的二进制值
2. 如何取其他位置的二进制值
3. 一共需要取31个位置的二进制值
4. 如何保存计算的结果值
5. 向左移动低位补0，向有移动高位补0
6. 逆序结果，从而调换低位和高位
