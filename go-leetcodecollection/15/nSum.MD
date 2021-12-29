[解决思路](https://mp.weixin.qq.com/s/fSyJVvggxHq28a0SdmZm6Q)


> nSum问题

对于nSum问题,可以抽象成一个递归

前提: 调用之前必须先将nums从低到高排序

递归公式: f(n)=fn(n-1) + nums[n]

终止条件: n<2或 n==2

base case: 

当n==2时,执行2Sum函数

2Sum两种常用解法

1.基于low，high 的双指针解法

2.用mapping



