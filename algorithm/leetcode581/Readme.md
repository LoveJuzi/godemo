# Shortest Unsorted Continuous Subarray

Easy

Share
Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order, too.

You need to find the shortest such subarray and output its length.

Example 1:
Input: [2, 6, 4, 8, 10, 9, 15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
Note:
Then length of the input array is in range [1, 10,000].
The input array may contain duplicates, so ascending order here means <=.

这道题我的想法就是先排序，然后比对

无序求解的方式也比较简单，但是需要的参数比较多次

这种题就是数学上的图形题，想想给定的数组，然后将起在直角坐标系中绘制出来，然后观察特征就行了。


