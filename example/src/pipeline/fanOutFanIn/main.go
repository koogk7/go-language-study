package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Fan-out : 채널이 닫힐 때까지 같은 채널에서 여러개의 function이 read
	Fan-in : 하나의 function이 여러개의 채널을 읽는 것
		--> sq
	gen 		---> mere 구조
		--> sq
	unbuffered 채널에 값을 주는 애는, 받는 애가 나오기 전까지 그 라인에 묶여있다.
*/

func main() {
	//in := gen(2, 3)
	var wg sync.WaitGroup
	var arr = []int{1, 2}

	wg.Add(1)
	gen(arr, &wg)
	time.Sleep(time.Second * 5)
	wg.Wait()
	fmt.Println("End")

	//c1 := sq(in)
	//c2 := sq(in)
	//
	//for n := range merge(c1, c2) {
	//	fmt.Println(n)
	//}
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	//Todo [refactor] 클로저 제거해보기
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func gen(nums []int, wg *sync.WaitGroup) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n // 다른 녀석이 빼갈 때 까지 기다림
		}
		close(out)
		wg.Done()
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
			n = n * 1
		}
		close(out)
	}()
	return out
}
