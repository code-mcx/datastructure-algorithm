package main

import "fmt"

// 背包问题（01背包）
// 有 n 件物品和一个容量为 v 的背包。第 i 件物品的重量是 w[i]，价值是 v[i]。
// 求解将哪些物品装入背包可使这些物品的重量总和不超过背包容量，且价值总和最大。
// 背包问题可以分 01 背包、完全背包、多重背包等。

// 01 背包问题要求每种物品仅有一件，可以选择放或不放。

// 假设现有如下物品
//
// | 物品名称 | 重量(kg) | 价值 |
// | ------- | -------- | ---- |
// | 音响     | 1      | 500  |
// | 电脑     | 2      | 5000 |
// | 手机     | 1      | 3000 |
//
// 给定一个容量为 3 的背包，不超过背包容量时，放入物品使得总价值最大，最大总价值是多少。

// 思路
// 将问题转换成填充一个二维表格。定义一个二维数组 value，value\[i\]\[j\] 表示放入前 i 件物品容量为 j 的最大价值总和。
//
// 第 0 件物品的任意容量总价值都是 0，容量为 0 时，放入任意一件物品的总价值也是 0。
// 如果当前物品的容量大于当前背包的容量，那么最大价值总和只能是前 i - 1 件物品放入
// 当前背包时的最大价值总和，最大价值为 value[i -1][j]。如果当前物品的容量
// 小于或等于当前背包的容量，那么最大价值总和就是第 i 件物品放入当前背包的价值 v[i]
// 加上前 i - 1 件物品放入剩余容量的最大值和 value[i - 1][v - w[i]]，
// 即 v[i] + value[i - 1][j - w[i]]，有可能放入前 i 件物品的最大价值和
// 比放入前 i - 1件最大价值和小，这里就要取两者中的最大值。
//
// 于是就有状态转移方程：
// 1. i == 0, j == 0 时，value[0][j] = value[i][0] = 0
// 2. w[i] > j 时，value[i][j] = value[i -1][j]
// 3. w[i] <= j 时，max{value[i -1][j], v[i] + value[i - 1][j - w[i]]}

func findMaxValue(w, v []int, c int) int {
    n := len(w)
    // 二维数组比实际物品数多出一行一列
    value := make([][]int, n + 1)
    // 初始化二维数组
    for i := 0; i < len(value); i++ {
        value[i] = make([]int, c + 1)
    }

    //// 初始第一列，默认为 0，也可以不初始化
    //for i := 0; i < len(value); i++ {
    //    value[i][0] = 1
    //}
    //// 初始第一行
    //for j := 0; j < len(value[0]); j++ {
    //    value[0][j] = 1
    //}

    // 遍历所有物品
    for i := 1; i <= n; i++ {
        // 遍历所有容量
        for j := 1; j <= c; j++ {
            // 实际物品下标从 0 开始
            if w[i - 1] > j {
                value[i][j] = value[i - 1][j]
            } else {
                value[i][j] = max(value[i - 1][j], v[i - 1] + value[i - 1][j - w[i - 1]])
            }
        }
    }

    fmt.Println("填表如下:")
    for i := 0; i < len(value); i++ {
        for _, val := range value[i] {
            fmt.Printf("%5d ", val)
        }
        fmt.Println()
    }
	return value[n][c]
}

func findMaxValue2(w, v []int, c int) int {
    // 数组比实际物品总数多出一列
    value := make([]int, c + 1)

    // 遍历所有物品
    for i := 0; i < len(w); i++ {
        // 遍历所有容量
        // 因为每个物品只能使用一次，给每一行填表时，采用倒序遍历
        for j := c; j >= w[i]; j-- {
            // 滑动数组优化空间复杂度
            value[j] = max(value[j], v[i] + value[j - w[i]])
        }
    }
	return value[c]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    // 物品重量(kg)
    w := []int{
        1, 2, 1,
    }
    // 物品价值
    v := []int{
        500, 5000, 3000,
    }
	// 背包容量
    c := 3
    // max := findMaxValue(w, v, c)
    max := findMaxValue2(w, v, c)
	fmt.Printf("最大价值总和为: %d", max)
}
