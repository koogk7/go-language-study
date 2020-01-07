package main

import (
	"fmt"
	"sync"
)

/*
	done 과 select를 이용한 early return 은 잠잭어르 블락된 샌더의 수를 알아야 하며
	샌더의 시그널을 정리해야 함으로 좋지 않다.
	그래서 close channel을 추천한다.
	닫힌 채널에 대한 리시브 오퍼레이션은 항상 즉시 zero value를 가져온다.
	채널을 닫는 것은 샌더에게 보드캐스팅하는 효율적인 방법이다.
*/
func main() {
	done := make(chan struct{}) // 모든 파이프라인들이 공유
	defer close(done)           // done을 닫으면 done을 가져갈려던 샌더들은 모두 제로값을 반환한다.

	in := gen(2)

	c1 := sq(done, in)
	c2 := sq(done, in)

	// Consume the first value from output
	out := merge(done, c1, c2)
	fmt.Println(<-out)
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return // 채널 종료
			}
		}
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
			out <- n // 다른 녀석이 빼갈 때 까지 기다림
		}
		close(out)
	}()
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}
