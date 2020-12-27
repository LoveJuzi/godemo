# X of a Kind in a Deck of Cards

Share
In a deck of cards, each card has an integer written on it.

Return true if and only if you can choose X >= 2 such that it is possible to split the entire deck into 1 or more groups of cards, where:

Each group has exactly X cards.
All the cards in each group have the same integer.
 
Example 1:

Input: deck = [1,2,3,4,4,3,2,1]
Output: true
Explanation: Possible partition [1,1],[2,2],[3,3],[4,4].
Example 2:

Input: deck = [1,1,1,2,2,2,3,3]
Output: false
Explanation: No possible partition.
Example 3:

Input: deck = [1]
Output: false
Explanation: No possible partition.
Example 4:

Input: deck = [1,1]
Output: true
Explanation: Possible partition [1,1].
Example 5:

Input: deck = [1,1,2,2,2,2]
Output: true
Explanation: Possible partition [1,1],[2,2],[2,2].
 
Constraints:

1 <= deck.length <= 10^4
0 <= deck[i] < 10^4

这道题目是一道划分相关的题目

题目没有提及元素的有序性，现在，如果不排序，我们能否完成题目的要求？

我的答案是可以的，但是需要使用哈希表进行数据统计

对所有统计值进行统计，查看是否有共同的公约数。

优先查找最大公约数，然后使用辗转法则，进一步减少最大公约数，如果最大公约是1，表示无解
