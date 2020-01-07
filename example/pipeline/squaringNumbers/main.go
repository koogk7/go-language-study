package main

import "fmt"

/*
	상대편이 준비될 때까지 채널에서 대기하며, 별도의 lock을 걸지 않고 데이터를 동기화 하는데 사용
	데이터가 없는 상태에서 채널이 닫히면 nil 값을 가져온다.
	채널에서 데이터를 가져오기 위해 대기하는 녀석이 있는데, 넣어주는 녀석이 없으면 deadlock!
*/

func main() {
	c := gen(2, 3)
	out := sq(c)

	fmt.Println(<-out)
	fmt.Println(<-out)

	// sq는 inbound, outbound가 같기 때문에 이런 식으로 조합 가능
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}
}

// 채널에 데이터를 담아서 리턴
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

// in <- chan int 를 받으므로, 받은 채널에서는 데이터를 꺼낼 수만 있다.
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
