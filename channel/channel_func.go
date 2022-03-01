package channel

import (
	"errors"
	"sort"
	"time"
)

func longTimeRequest(i int) <-chan int {
	req := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		req <- i
	}()

	return req
}

func future_promise() int {
	a, b := longTimeRequest(10), longTimeRequest(20)

	sum := func(i1, i2 int) int {
		return i1*i1 + i2*i2
	}(<-a, <-b)

	return sum // block here until both of caculation be finished
}

func longTimeRequestChan(i int, ch chan int) {
	time.Sleep(time.Second * 1)
	ch <- i
}

func future_promise_2() int {
	ch := make(chan int, 2)
	go longTimeRequestChan(10, ch)
	go longTimeRequestChan(20, ch)

	return <-ch + <-ch // block here until both of caculation be finished
}

func getFirstInput() int {
	ch := make(chan int, 5)
	for i := 1; i <= cap(ch); i++ {
		n := i * i
		go func() {
			ch <- n
			time.Sleep(time.Second * 1)
		}()
		time.Sleep(time.Second * 2)
	}
	return <-ch // block here until ch get first value
}

func getResponseWithError() (int, error) {
	ch := make(chan struct {
		i int
		e error
	})

	go func() {
		time.Sleep(time.Second * 2)
		ch <- struct {
			i int
			e error
		}{-1, errors.New("err")}
	}()

	ret := <-ch // block here until ch get something

	return ret.i, ret.e
}

func oneToOneChanInform() (int, int) {
	val := make([]int, 10)
	for i := 0; i < 10; i++ {
		val[i] = i
	}
	done := make(chan struct{})

	go func() {
		sort.Slice(val, func(i, j int) bool {
			return val[i] > val[j]
		})
		done <- struct{}{}
	}()

	<-done // block here until sort finish
	return val[0], val[len(val)-1]
}

func oneToOneChanInform_2() int {
	done := make(chan struct{})
	ans := 0

	go func() {
		time.Sleep(time.Second)
		ans = 1
		<-done // inform main thread it finish
	}()

	done <- struct{}{} // block here until sort finish
	return ans
}

func informGroupByCloseChan() int {
	worker := func(id int, ready <-chan struct{}, done chan<- int) {
		<-ready
		time.Sleep(time.Second * time.Duration(id))
		done <- id
	}

	ready, done := make(chan struct{}), make(chan int)
	go worker(1, ready, done)
	go worker(2, ready, done)
	go worker(3, ready, done)

	time.Sleep(time.Second) // simulate a init duration

	close(ready) // close chan ready to inform all worker, because we can get unlimited values from a closing chan

	ret := 0
	ret += <-done
	ret += <-done
	ret += <-done

	return ret
}

// this func is similar with time.After()
// Normally, we should use time.After()
func afterDuration(d time.Duration) <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()
	return c
}

func informWithTimer() int {
	return <-afterDuration(time.Second * 2)
}

// use chan as mutex
// but we recommand to use sync.mutex, its performance is greater than use chan
func useChanAsMutex() int {
	mutex := make(chan struct{}, 1)

	data := 0
	increase := func() {
		mutex <- struct{}{}
		data++
		<-mutex
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done
	return data
}
