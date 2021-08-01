package concurrency

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var chain string

/*
 * 运行结果：无法编译，A() 和 C() 中的 mu.Lock() 操作死锁
 */
func Quiz01() {
	chain = "main"
    MuA()
	fmt.Println(chain)
}

func MuA() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	MuB()
}

func MuB() {
	chain = chain + " --> B"
    MuC()
}

func MuC() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}
