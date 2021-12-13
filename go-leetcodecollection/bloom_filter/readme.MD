# Bloom-Filter

## 简介

Bloom Filter（BF）是一种空间效率很高的随机数据结构，它利用位数组很简洁地表示一个集合，并能判断一个元素是否属于这个集合。

它的优点是空间效率和查询时间都远远超过一般的算法，缺点是有一定的误识别率和删除困难。


## 重点
一个Bloom Filter有以下参数：
- m	bit数组的宽度（bit数）
- n	加入其中的key的数量
- k	使用的hash函数的个数
- f	False Positive的比率

## 参考链接
[CSDN](https://blog.csdn.net/hguisu/article/details/7866173)

[掘金](https://juejin.cn/post/6855839313859461133)