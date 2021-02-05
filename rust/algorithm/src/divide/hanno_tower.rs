/// 汉诺塔
/// 大梵天创造世界的时候做了三根金刚石柱子，在一根柱子上从下往上按照大小
/// 顺序摞着 64 片黄金圆盘大梵天命令婆罗门把圆盘从下面开始按大小顺序重新
/// 摆放在另一根柱子上。并且规定，在小圆盘上不能放大圆盘，在三根柱子之间
/// 一次只能移动一个圆盘

/// 思路
/// 假设有 A、B、C 三个柱子，要把 A 柱子上的所有盘放到 C 柱子上
/// 如果只有一个盘，直接将盘从 A 柱子移动到 C 柱子
/// 如果有两个或以上的盘，将盘分成两份，最下面一个盘一份和上面所有的面盘一份
/// 先将上面所有的盘移动到 B 柱子，然后将下面的盘从 A 柱子移动到 C 柱子，再
/// 将 B 柱子上的所有盘移动到 C 柱子上
fn hanno_tower(num: i32, a: char, b: char, c: char) {
    if num == 1 { // 只有一个盘
        println!("第 1 个盘从 {} 到 {}", a, c);
        return;
    }
    // 先将除了最后一个盘之外的所有盘从 a 移动到 b
    hanno_tower(num - 1, a, c, b);
    // 最后一个盘从 a 移动到 c
    println!("第 {} 个盘从 {} 到 {}", num, a, c);
    // 再将 b 柱子上的所有盘移动到 c 柱子上
    hanno_tower(num - 1, b, a, c);
}

#[test]
fn test_hanno_tower() {
    hanno_tower(3, 'A', 'B', 'C');
}
