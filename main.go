package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func hello(done chan bool) {
	log.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	log.Println("hello go routine awake and going to write to done")
	done <- true
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func calcSquares(number int, out chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	out <- sum
}

func calcCubes(number int, out chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	out <- sum
}

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func Greet(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("Hello", name)
}

func main() {
	defer log.Println("Go routine demo completed")

	done := make(chan bool)
	log.Println("Main going to call hello go goroutine")
	go hello(done)
	ok := <-done
	log.Println("Main received data ", ok)

	var a chan int
	if a == nil {
		log.Println("channel a is nil, going to define it")
		a = make(chan int)
		log.Printf("Type of a is %T\n", a)
	}

	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	log.Println("Final output", squares+cubes)

	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		log.Println("Received ", v, ok)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go Greet("World", &wg)
	go Greet("Universe", &wg)

	log.Println("Waiting for Greetings to finish")
	wg.Wait()
}
