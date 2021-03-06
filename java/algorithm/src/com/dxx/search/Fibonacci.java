package com.dxx.search;

import java.util.Arrays;

/**
 * 斐波那契查找
 * 中间值计算公式: mid = start + F[k - 1] - 1
 */
public class Fibonacci {

    public static int maxSize = 20;  // 斐波那契数组长度

    /**
     * 获取斐波那契数列
     */
    public static int[] fibonacci() {
        int[] fib = new int[maxSize];
        fib[0] = 1;
        fib[1] = 1;
        for (int i = 2; i < maxSize; i++) {
            fib[i] = fib[i - 1] + fib[i - 2];
        }
        return fib;
    }

    public static int fibonacciSearch(int[] nums, int findVal) {
        int start = 0;
        int end = nums.length - 1;
        int k = 0; // 斐波那契分割数的下标
        int mid = 0; // 中间值
        int[] f = fibonacci(); // 斐波那契数列

        // 获取斐波那契分割数的下标
        while (end > f[k] - 1) {
            k++;
        }
        // 创建一个切片，长度为斐波那契分割数的值
        int dstLength = f[k];
        int[] dst;
        // 将要查找的数组复制到 dst 中
        dst = Arrays.copyOf(nums, dstLength);
        // 将最后一个元素填充到 dst 元素为 0 的位置
        // [1, 8, 10, 89, 100, 100, 123] => [1 8 10 89 100 100 123 123]
        for (int i = end + 1; i < dstLength; i++) {
            dst[i] = nums[end];
        }

        // 使用循环来寻找 findVal
        while (start <= end) {
            mid = start + f[k - 1] - 1;
            if (findVal < dst[mid]) { // 向左查找
                end = mid - 1;
                // 全部元素 = 前面的元素 + 后边元素
                // 即f[k] = f[k-1] + f[k-2]
                // 因为前面有 f[k-1]个元素,所以可以继续拆分 f[k-1] = f[k-2] + f[k-3]
                // 即 在 f[k-1] 的前面继续查找 k--
                // 即下次循环 mid = f[k-1-1]-1
                k -= 1;
            } else if (findVal > dst[mid]) { // 向右查找
                start = mid + 1;
                // 全部元素 = 前面的元素 + 后边元素
                // f[k] = f[k-1] + f[k-2]
                // 因为后面我们有f[k-2] 所以可以继续拆分 f[k-1] = f[k-3] + f[k-4]
                // 即在f[k-2] 的前面进行查找 k -=2
                // 即下次循环 mid = f[k - 1 - 2] - 1
                k -= 2;
            } else {
                if (mid <= end) {
                    return mid; // 若相等则说明 mid 即为查找到的位置
                } else {
                    return end; // 若 mid > end 则说明是扩展的数值，返回 end
                }
            }
        }
        return -1;
    }

    public static void main(String[] args) {
        int value = 100;
        int[] nums = new int[]{1, 8, 10, 89, 100, 100, 123};
        int index = fibonacciSearch(nums, value);
        if (index != -1) {
            System.out.printf("找到 %d, 下标为 %d\n", value, index);
        } else {
            System.out.printf("未找到 %d\n", value);
        }
    }
}
