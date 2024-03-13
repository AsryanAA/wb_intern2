// программа выведет числа от 0 до 9, после упадет с ошибкой fatal error: all goroutines are asleep - deadlock! (забыли закрыть канал) 

package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// close(ch) // чтобы исправить
	}()

	for n := range ch {
		println(n)
	}
}
