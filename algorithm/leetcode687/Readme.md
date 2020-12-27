# Longest Univalue Path

Given a binary tree, find the length of the longest path where each node in the path has the same value. This path may or may not pass through the root.

The length of path between two nodes is represented by the number of edges between them.

Example 1:

Input:

              5
             / \
            4   5
           / \   \
          1   1   5
Output: 2

Example 2:

Input:

              1
             / \
            4   5
           / \   \
          4   4   5
Output: 2

Note: The given binary tree has not more than 10000 nodes. The height of the tree is not more than 1000.

这道题目是用来计数的，用来数边的个数。

这道题目不是平凡问题，需要对树进行划分

划分方法：根，左子树，右子树

关键因素：只有一个，就是根和左子树以及右子树的根的键值比较。而且只是恒等比较。

这道题目属于简单题目，不论是划分，还是关键因素，都是比较容易总结

根值 == 左子树根值， PATH 长度 + 1 + 左子树对应根值的最大长度。

根植 == 右子树根值， PATH 长度 + 1 + 右子树对应根值的最大长度。

当前最大PATH值，和新计算的值进行比较，更新最大PATH值。

左子树空  PATH = 0
右子树空  PATH = 0

根值 != 左子树根值 PATH = 0
根值 != 右子树根值 PATH = 0
