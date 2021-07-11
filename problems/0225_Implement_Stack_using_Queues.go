package problems

/**
 * LeetCode: https://leetcode-cn.com/problems/implement-stack-using-queues/
 */

type MyStack struct {
	QueueR, QueueW []int
}


/** Initialize your data structure here. */
func StackConstructor() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.QueueW = append(this.QueueW, x)
	for i := 0; i < len(this.QueueR); i++ {
		this.QueueW = append(this.QueueW, this.QueueR[i])
	}
	this.QueueR, this.QueueW = this.QueueW, nil
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	val := this.QueueR[0]
	this.QueueR = this.QueueR[1:]
	return val
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.QueueR[0]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.QueueR) == 0 {
		return true
	}
	return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
