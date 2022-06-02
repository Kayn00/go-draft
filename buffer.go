package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Person struct {
	name string
	age  int
}

// 全局变量（简单处理）
var p atomic.Value
var mu sync.Mutex

func update(name string, age int) {
	lp:=&Person{}
	// 更新第一个字段
	lp.name = name
	// 加点随机性
	time.Sleep(time.Millisecond*200)
	// 更新第二个字段
	lp.age = age

	p.Store(lp)
}

func main() {
	now := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(100)
	// 10 个协程并发更新
	for i := 0; i < 100; i++ {
		name, age := fmt.Sprintf("nobody:%v", i), i
		go func() {
			defer wg.Done()
			update(name, age)
		}()
	}
	wg.Wait()
	// 结果是啥？你能猜到吗？
	_p := p.Load().(*Person)
	fmt.Printf("p.name=%s\np.age=%v\n", _p.name, _p.age)
	fmt.Println(time.Since(now))
}