package main

import (
    "errors"
    "fmt"
    "strconv"
)

// 队列
// 队列是一种特殊的线性表，特殊之处在于它只允许在表的前端进行删除操作，而在表的后端进行插入操作
// 所以队列又称为先进先出（FIFO—first in first out）

// 循环队列
// 在实际使用队列时，为了使队列空间能重复使用，往往对队列的使用方法稍加改进：无论插入或删除，
// 一旦 rear 指针增 1 或 front 指针增 1 时超出了所分配的队列空间，就让它指向这片连续空间的起始位置
// 实际上是把队列空间想象成一个环形空间，环形空间中的存储单元循环使用，用这种方法管理的队列也就称为循环队列

// 数组实现循环队列

type IntQueue struct {
    array   []int // 存放队列元素的切片（数组无法使用变量来定义长度）
    maxSize int   // 最大队列元素大小
    front   int   // 队头指针
    rear    int   // 队尾指针
}

func NewQueue(maxSize int) *IntQueue {
    return &IntQueue{
        array:   make([]int, maxSize),
        maxSize: maxSize,
        front:   0,
        rear:    0,
    }
}

// 放入队列元素
func (q *IntQueue) Put(elem int) error {
    if q.isFull() {
        return errors.New("queue is full")
    }
    q.array[q.rear] = elem
    // 循环累加，当 rear + 1 等于 maxSize 时变成 0，重新累加
    q.rear = (q.rear + 1) % q.maxSize
    return nil
}

// 取出队列元素
func (q *IntQueue) Take() (int, error) {
    if q.isEmpty() {
        return 0, errors.New("queue is empty")
    }
    elem := q.array[q.front]
    q.front = (q.front + 1) % q.maxSize
    return elem, nil
}

func (q *IntQueue) isEmpty() bool {
    // 队头指针等于队尾指针表示队列为空
    return q.front == q.rear
}

func (q *IntQueue) isFull() bool {
    // 空出一个位置，判断是否等于队头指针
    // 队尾指针指向的位置不能存放队列元素，实际上会比 maxSize 指定的大小少一
    return (q.rear+1)%q.maxSize == q.front
}

func (q *IntQueue) size() int {
    return (q.rear + q.maxSize - q.front) % q.maxSize
}

// 重新定义 String 方法，方便输出
func (q *IntQueue) String() string {
    str := "["
    tempFront := q.front
    for i := 0; i < q.size(); i++ {
        str += strconv.Itoa(q.array[tempFront]) + " "
        // 超过最大大小，从 0 开始
        tempFront = (tempFront + 1) % q.maxSize
    }
    str += "]"
    return str
}

func main() {
    intQueue := NewQueue(5)
    _ = intQueue.Put(1)
    _ = intQueue.Put(2)
    _ = intQueue.Put(3)
    _ = intQueue.Put(4)
    _ = intQueue.Put(5) // 队列已满，无法放入数据，实际上只能放 4 个元素

    fmt.Println("intQueue:", intQueue)

    num, _ := intQueue.Take()
    fmt.Println("取出一个元素:", num)
    num, _ = intQueue.Take()
    fmt.Println("取出一个元素:", num)
    num, _ = intQueue.Take()
    fmt.Println("取出一个元素:", num)
    num, _ = intQueue.Take()
    fmt.Println("取出一个元素:", num)
    num, takeErr := intQueue.Take()
    fmt.Println("取出一个元素:", num)
    if takeErr != nil {
        fmt.Println("出队失败:", takeErr)
    }

    // 取出数据后可以继续放入数据
    _ = intQueue.Put(5)
    fmt.Println("intQueue:", intQueue)
}