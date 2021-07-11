package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/implement-queue-using-stacks/
 */

type MyQueue struct {
	StackR, StackW []int
}

/** Initialize your data structure here. */
func QueueConstructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.StackW = append(this.StackW, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	val := this.Peek()
	this.StackR = this.StackR[:len(this.StackR)-1]
	return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	val := 0
	if len(this.StackR) == 0 {
		for i := len(this.StackW)-1; i >= 0; i-- {
			this.StackR = append(this.StackR, this.StackW[i])
		}
		this.StackW = nil
	}
	endIdx := len(this.StackR) - 1
	val = this.StackR[endIdx]
	return val
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.StackR) + len(this.StackW) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
