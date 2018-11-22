package main

import (
	"log"
	"time"
	"math/rand"
)

func main() {
	works := make(chan int, 3)

	for i:=0;i<3;i++{
		go func() {
			for {
				word := <-works
				log.Println(word)
				time.Sleep(time.Second*5)
			}
		}()
	}

	for{
		works<-rand.Intn(10)
	}
}
