package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} { // функция для объединения каналов (merge channels)
	var wg sync.WaitGroup
	out := make(chan interface{})

	for _, ch := range channels { // итерируемся по каналам и для каждого канала запускаем горутину
		wg.Add(1) // увеличиваем счетчик ожидания
		go func(ch <-chan interface{}) {
			defer wg.Done() // уменьшаем счетчик ожидания
			<-ch            // блокируем горутину и ждем сигнал (в данном случае сигнал закрытия канала из функции sig)
		}(ch) // передаем каждый канал в горутину
	}
	go func() {
		wg.Wait()  // ждем пока все горуины отрботают
		close(out) // закрываем канал
	}()
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		// блокируется горутина main и ждем сигнал
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("fone after %v", time.Since(start))
}
