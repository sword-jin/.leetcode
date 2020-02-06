/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue: []int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	this.shift()
	x := this.queue[0]
	this.queue = this.queue[1:]
	return x
}

func (this *MyStack) shift() {
	for i := 0; i < len(this.queue)-1; i++ {
		x := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, x)
	}
}

/** Get the top element. */
func (this *MyStack) Top() int {
	this.shift()
	x := this.queue[0]
	this.queue = this.queue[1:]
	this.queue = append(this.queue, x)
	return x
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

