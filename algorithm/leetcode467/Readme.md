# Unique Substrings in Wraparound String

Consider the string s to be the infinite wraparound string of "abcdefghijklmnopqrstuvwxyz", so s will look like this: "...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd....".

Now we have another string p. Your job is to find out how many unique non-empty substrings of p are present in s. In particular, your input is the string p and you need to output the number of different non-empty substrings of p in the string s.

Note: p consists of only lowercase English letters and the size of p might be over 10000.

Example 1:
Input: "a"
Output: 1

Explanation: Only the substring "a" of string "a" is in the string s.
Example 2:
Input: "cac"
Output: 2
Explanation: There are two substrings "a", "c" of string "cac" in the string s.
Example 3:
Input: "zab"
Output: 6
Explanation: There are six substrings "z", "a", "b", "za", "ab", "zab" of string "zab" in the string s.

p is very big

已知：zab 是 s 的一个子串，问：zab的非空子集有多少是 s 的子串？

3 + 2 + 1 = 6

已知：zabc 是 s 的一个子串，问：zabc的非空子集有多少是 s 的子串？

4 + 3 + 2 + 1 = 10

n项字符串的和

"z", "a", "b", "za", "ab", "zab" 正好是 zab 的子集，除了空集

集合归属问题

如何给p进行划分，才能满足上述条件。连续字符串是s的子串

“zabdddabcdefxyzabcdhijk” 的一个有效划分如下：“zab d d d abcdef xyzabcd hijk”

这个题目依然是一个计数问题。这种计数问题计算机有很多类似的问题，不过这种类型的问题，我是第一次见到，总结它的计数规则优点困难。
我们看看有没有其他的方法可以做这个东西。

计数问题中比较麻烦的就是去除重复项。

1 + 2 + 3 + 1 + 1 + 1 + 1 + 2 + 3 + 4 + 5 + 6 + 1 + 2 + 3 + 4 + 5 + 6 + 7 + 1 + 2 + 3 + 4

d=1

当前距离
之前距离

距离

  zab
      d
      d
      d
   abcdef
xyzabcd
          hijk

  1+2+3
        1
        1-1
        1-1
    1-1+2-2+3+4+5+6
1+2+3-1+4-2+5-3+6-3+7
           1+2+3+4

如何证明我的分析是对的？

去除重复集合。。。这个条件没有考虑

zabczabc的一个有效结果集的元素数量是 10，并不是 20

a b c d e f g h i j k l m n o p q r s t u v w x y z

<!-- z: (z, 1) -->
<!-- a: (z, 1) -->
<!-- b: (z, 1) -->
<!-- d: (d, 1) -->
<!-- d: (d, 1) -->
<!-- d: (d, 1) -->
<!-- a: (a, 1) -->
<!-- b: (a, 1) -->
<!-- c: (a, 1) -->
<!-- d: (a, 1) -->
<!-- e: (a, 1) -->
<!-- f: (a, 1) -->
<!-- x: (x, 1) -->
<!-- y: (x, 1) -->
<!-- z: (x, 1) -->
<!-- a: (x, 1) -->
<!-- b: (x, 1) -->
<!-- c: (x, 1) -->
<!-- d: (x, 1) -->
<!-- h: (h, 1)
i: (h, 1)
j: (h, 1)
k: (h, 1) -->

a: (z, 1)
a: (a, 1)
a: (x, 1)

b: (z, 1)
b: (a, 1)
b: (x, 1)

c: (a, 1)
c: (x, 1)

d: (d, 1)
d: (d, 1)
d: (d, 1)
d: (a, 1)
d: (x, 1)

e: (a, 1)

f: (a, 1)

h: (h, 1)
i: (h, 1)
j: (h, 1)
k: (h, 1)

x: (x, 1)

y: (x, 1)

z: (z, 1)
z: (x, 1)

什么叫重复计数

yzab
 zab
  abc
  abc
  abcd

1+2+3+4
  1-1+2-2+3-3
  1-1+2-2+3
  1-1+2-2+3-3
  1-1+2-2+3-3+4

这个就是划分的总和

y: (y, 1)
z: (y, 1)
a: (y, 1)
b: (y, 1)

<!-- z: (z, 1) -->
<!-- a: (z, 1) -->
<!-- b: (z, 1) -->

<!-- a: (a, 2)
b: (a, 2)
c: (a, 2) -->

-abc

a: (a, 1)
b: (a, 1)
c: (a, 1)

z: (y, 1)
<!-- z: (z, 1) -->
a: (y, 1)
<!-- a: (z, 1) -->
b: (y, 1)
<!-- b: (z, 1) -->

-zab

y: (y, 1)
z: (y, 1)
a: (y, 1)
b: (y, 1)

a: (a, 1)
b: (a, 1)
c: (a, 1)

优先去大后去小

大的里面往往包含了小的计数，优先去大，这样有助于减轻去小的压力。
