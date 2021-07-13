package problems

import "fmt"

/**
 * LeetCode: https://leetcode-cn.com/problems/lru-cache/
 */

type LRUCache struct {
	size, capacity int
	head, tail     *DLinkedNode
	cache          map[int]*DLinkedNode
}

type DLinkedNode struct {
	key, value int
	pre, next  *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	cache.head.pre = cache.tail
	cache.tail.next = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	ret := -1
	if node, ok := this.cache[key]; ok {
		ret = this.cache[key].value
		this.moveToHead(node)
	}
	return ret
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		if this.size == this.capacity {
			node := this.tail.next
			this.removeNode(node)
			delete(this.cache, node.key)
			this.size--
		}
		node := &DLinkedNode{key: key, value: value}
		this.addToHead(node)
		this.cache[key] = node
		this.size++
	}
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head
	node.pre = this.head.pre
	node.pre.next = node
	node.next.pre = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) Watch() {
	fmt.Println("---------------------")
	fmt.Printf("size=%d, cap=%d\n", this.size, this.capacity)
	for node := this.tail; node != nil; node = node.next {
		fmt.Printf("k = %d, v = %d\n", node.key, node.value)
	}
	fmt.Println("---------------------")
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
