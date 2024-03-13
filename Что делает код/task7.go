package main

import (
	"fmt"
	// "math/rand"
	"sync"
	// "time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(vs))
	go func() {
		for _, v := range vs {
			c <- v
			// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			wg.Done()
		}
	}()

	go func() {
		wg.Wait()
		fmt.Println("hello")
		close(c)
	}()

	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			default:
				_, closeA := <-a
				_, closeB := <-b
				if !closeA && !closeB {
					close(c)
					break
				}
			}
		}
	}()
	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
