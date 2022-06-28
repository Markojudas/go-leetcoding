package main

import "fmt"

type MyQueue struct {
	list []int
}

func Constructor() MyQueue {

	mq := MyQueue{
		list: []int{},
	}

	return mq
}

func (mq *MyQueue) Push(x int) {

	mq.list = append(mq.list, x)
}

func (mq *MyQueue) Pop() int {

	elem := mq.Peek()

	if len(mq.list) > 1 {
		mq.list = mq.list[1:]
	} else {
		mq.list = []int{}
	}
	return elem
}

func (mq *MyQueue) Peek() int {

	return mq.list[0]
}

func (mq *MyQueue) Empty() bool {

	return len(mq.list) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1) //null
	obj.Push(2) // null

	fmt.Println(obj.Peek())  //1
	fmt.Println(obj.Pop())   //1
	fmt.Println(obj.Empty()) //false
}
