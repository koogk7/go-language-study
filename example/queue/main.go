package main

import (
	"fmt"
	"k8s.io/client-go/util/workqueue"
	"time"
)

func main() {
	stop := make(chan int)
	que := workqueue.New()
	go fillQueue(que)
	go readFromQueue(que, stop)
	<-stop
}

// 인터페이스는 포인터로 선언하지 않아도, 포인터 타입을 받을 수 있네!
func fillQueue(queue workqueue.Interface) {
	time.Sleep(time.Second)
	queue.Add("this")
	queue.Add("is")
	queue.Add("a")
	queue.Add("complete")
	queue.Add("sentence")
}

func readFromQueue(queue workqueue.Interface, stop chan<- int) {
	time.Sleep(3 * time.Second)
	for queue.Len() > 0 {
		item, _ := queue.Get()
		fmt.Printf("Got item : %s\n", item)
		queue.Done(item)
	}
	stop <- 1
}
