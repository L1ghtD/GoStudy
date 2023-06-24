package main

import (
	"fmt"
	"sync"
	"time"
)

/*
生产者消费者模型
*/

type Queue struct {
	queue []string
	cron  *sync.Cond
}

func main() {
	// 初始化
	q := Queue{
		queue: []string{},
		cron:  sync.NewCond(&sync.Mutex{}),
	}
	go func() {
		for {
			q.Enqueue("a")
			time.Sleep(2 * time.Second)
		}
	}()
	for {
		q.Dequeue()
		time.Sleep(1 * time.Second)
	}
}

func (q *Queue) Enqueue(item string) {
	// cron 加锁,结构体自带方法
	q.cron.L.Lock()
	defer q.cron.L.Unlock()
	q.queue = append(q.queue, item)
	fmt.Printf("Putting #{item} to queue, notify all\n")
	// 通知 cron.wait() 等待的线程
	q.cron.Broadcast()
}

func (q *Queue) Dequeue() string {
	q.cron.L.Lock()
	defer q.cron.L.Unlock()
	if len(q.queue) == 0 {
		fmt.Println("no data available, wait...")
		// 阻塞，当 Enqueue 发出 Broadcast 信号时解除
		q.cron.Wait()
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}
