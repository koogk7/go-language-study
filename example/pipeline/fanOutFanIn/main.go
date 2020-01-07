package main

import (
	"fmt"
	"sync"
)

/*
	Fan-out : 채널이 닫힐 때까지 같은 채널에서 여러개의 function이 read
	Fan-in : 하나의 function이 여러개의 채널을 읽는 것
		--> sq
	gen 		---> mere 구조
		--> sq
*/

func main() {
	in := gen(2, 3)

	c1 := sq(in)
	c2 := sq(in)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
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

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
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
