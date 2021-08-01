package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var rwMu sync.RWMutex
var count int

/*
 * 运行结果：1
 * sync.RWMutex 解析：https://dongxiem.github.io/2020/06/07/golang-sync-bao-yuan-ma-pou-xi-2-sync.rwmutex/
 */
func Quiz02() {
	go RWMuA()
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}

func RWMuA() {
	rwMu.RLock()
	defer rwMu.RUnlock()
	RWMuB()
}

func RWMuB() {
	time.Sleep(5 * time.Second)
	RWMuC()
}

func RWMuC() {
	rwMu.RLock()
	defer rwMu.RUnlock()
}
