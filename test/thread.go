package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func A() {
	fmt.Println("start!", GetGID())
	group := sync.WaitGroup{}
	count := 10000
	group.Add(count)
	for i := 0; i < count; i++ {
		go func(group *sync.WaitGroup) {
			defer group.Done()
			fmt.Println("this is A", GetGID())
			time.Sleep(time.Second)
		}(&group)
	}
	group.Wait()
	fmt.Println("end!", GetGID())
}

func B() {
	fmt.Println("start!")
	ch := make(chan bool, 1)
	go func(ch2 chan bool) {
		defer close(ch2)
		fmt.Println("this is B")
		ch2 <- true
	}(ch)
	fmt.Println("wait!")
	<-ch
	fmt.Println("end!")
}

func C() {
	fmt.Println("start!")
	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)
	cond.L.Lock()
	var done = false

	go func() {
		fmt.Println("this is C")
		cond.Signal()
		done = true
	}()

	if !done {
		cond.Wait()
		cond.L.Unlock()
	}

	fmt.Println("end!")
}

func main() {
	A()
	//B()
	//C()
}
