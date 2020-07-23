# Long Pressed Name

Easy

648

98

Add to List

Share
Your friend is typing his name into a keyboard.  Sometimes, when typing a character c, the key might get long pressed, and the character will be typed 1 or more times.

You examine the typed characters of the keyboard.  Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.

 

Example 1:

Input: name = "alex", typed = "aaleex"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.
Example 2:

Input: name = "saeed", typed = "ssaaedd"
Output: false
Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.
Example 3:

Input: name = "leelee", typed = "lleeelee"
Output: true
Example 4:

Input: name = "laiden", typed = "laiden"
Output: true
Explanation: It's not necessary to long press any character.
 

Constraints:

1 <= name.length <= 1000
1 <= typed.length <= 1000
The characters of name and typed are lowercase letters.

这道题目是字符串比较问题

字符串比较问题是两个集合间的关系问题

A   B 

A 表示原始字符串

B 表示键入的字符串

现在要验证 B 中是否有 A 的字序列串

这种题目是平凡问题

为什么是平凡问题，我们仅仅需要进行一次有效的遍历就可以解决这个问题 

如果能正常遍历A，那么就是一个完成名字。

两个字符串是否相等算法之上的一个修正算法
