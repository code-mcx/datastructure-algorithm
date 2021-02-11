package com.mcx.linkedlist.joseph;

/**
 * 约瑟夫问题
 * 设编号为 1,2,...n 的 n 个人围坐一圈约定编号为 k(1<=k<=n) 的人
 * 从 1 开始报数，数到 m 的那个人出列它的下一位又从 1 开始报数数到 m
 * 的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列
 */
public class Main {

    public static class Person {
        private int no;
        private Person prev;
        private Person next;

        public Person() {}
        public Person(int no) {
            this.no = no;
        }
    }

    public static class PersonLinkedList {
        private Person first; // 第一个小孩
        private int length; // 小孩的数量

        public PersonLinkedList(int count) {
            if (count <= 0) {
                System.err.println("链表至少需要一个元素");
                return;
            }
            this.length = count;
            Person prev = this.first;
            // 初始化小孩，构成双向循环链表
            for (int i = 1; i <= count; i++) {
                Person person = new Person(i);
                if (i == 1) {
                    // 初始化 First 节点
                    this.first = person;
                    this.first.next = this.first;
                    this.first.prev = this.first;

                    // 将 prev 指向 First 节点，继续下一次循环
                    prev = this.first;
                    continue;
                }

                prev.next = person;
                person.prev = prev;

                // 新增加的节点的下一个节点指向第一节点
                person.next = this.first;
                // 第一个节点的上一个节点指向新增加的节点
                this.first.prev = person;

                prev = person;
            }
        }

        public void showPersons() {
            if (this.first == null) {
                return;
            }
            if (this.first.next == this.first) {
                System.out.printf("num:%d\n", this.first.no);
                return;
            }
            Person current = this.first;
            while (true) {
                System.out.printf("num:%d\n", current.no);
                current = current.next;

                if (current == this.first) {
                    break;
                }
            }
        }

        public void count(int start, int num) {
            if (start < 1 || start > this.length) {
                System.out.printf("start 不能小于 1 或者不能大于 %d\n", this.length);
                return;
            }
            if (num > this.length) {
                System.out.printf("num 不能大于元素个数: %d\n", this.length);
                return;
            }

            Person current = this.first;

            // 循环 start - 1 次
            for (int i = 1; i <= start - 1; i++) {
                current = current.next;
            }

            while (true) {
                // 表示只有一个节点
                if (current.prev == current && current.next == current) {
                    break;
                }

                // 循环 num - 1 次
                for (int i = 1; i <= num - 1; i++) {
                    current = current.next;
                }

                // 删除元素
                current.prev.next = current.next;
                current.next.prev = current.prev;

                System.out.printf("出队人的编号: %d\n", current.no);
                current = current.next;
            }
            System.out.printf("最后留下人的编号: %d\n", current.no);
        }
    }

    public static void main(String[] args) {
        PersonLinkedList personLinkedList = new PersonLinkedList(5);
        personLinkedList.showPersons();

        personLinkedList.count(1, 3);
    }
}
