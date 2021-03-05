双指针 

二分查找

二叉树
  inorder preoder postorder  递归 / 迭代
  层序遍历

队列

BFS 的两个主要方案：遍历或找出最短路径


去理解,最基本的两种的存储结构,在每种程序设计语言中,都设计了相对应的"data type"(数据类型)去描述它.
例如我们在Go中可以定义一个Array Type,它就是用于描述顺序存储结构的.
Go语言中,每种可以被定义的数据类型,它都是怎么被设计来描述啥的:
顺序存储结构可以用哪些Type来描述: [3]int 
链式存储结构可以用哪些Type来描述: 
Go内置库:
https://golang.org/pkg/container/list/  双向链表
https://golang.org/pkg/container/heap/  堆操作
https://golang.org/pkg/container/ring/  循环列表



### 关于数据结构(Hold住基本面)
如何合理地组织数据、高效地处理数据，这就是“数据结构”主要研究的问题.

数据结构主要研究非数值计算问题，非数值计算问题无法用数学方程建立数学模型.

数据在计算机中的存储结构: `顺序存储结构` 和 `链式存储结构`.

当然,我们更熟悉的是,在分析具体问题时,我们通常可以抽象成一些和实际存储无关的逻辑结构:
- 线性结构
    - 一般线性表: 
        - 线性表
    - 特殊线性表: 
        - 栈
        - 队列
        - 字符串
    - 线性表的推广: 
        - 数组
        - 广义表

- 非线性结构
    - 树
        - 二叉树
    - 图
        - 有向图
        - 无向图


### Thanks
[fucking-algorithm/学习数据结构和算法的高效方法.md](https://github.com/labuladong/fucking-algorithm/blob/master/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/%E5%AD%A6%E4%B9%A0%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E5%92%8C%E7%AE%97%E6%B3%95%E7%9A%84%E9%AB%98%E6%95%88%E6%96%B9%E6%B3%95.md)