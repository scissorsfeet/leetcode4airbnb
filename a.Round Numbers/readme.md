# 题意
Given an array of numbers A=[x1,x2...,xn] and 
T=Round(x1+x2+...+xn).We want to find a way to 
round each element in A such that after rounding we get
a new array B=[y1,y2,...,yn] such that y1+y2+...+yn=T where
yi=Floor(xi) or Ceil(xi), ceiling or floor of xi.
We also want to minimize sum |xi-yi|.

A是一个浮点数组，将所有元素累加后得到的值进行四舍五入圆整后得到T，
对每个元素单独向上或向下圆整后得到B，并且B中每个元素累加后的和等于T，
即Round(sum(A)) = sum(B),并且要使sum(|xi-yi|)最小

# 思路
1. 将A中所有元素都向下取整后累加得到floorSum,
则Round(sum(A))-floorSum就是所有数字向下取整后需要补充的差值
2. 对A数列进行排序，排序规则是按照xi-floor(xi)递增，则在最后面
的就是差值最大的，采用贪心的思想将其改为向上取整即可