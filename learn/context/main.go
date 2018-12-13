package main

import (
	"context"
	"fmt"
	"time"
)

func processHandler(ctx context.Context) {
	//ch1 := make(chan string)
	ct, cancle := context.WithCancel(ctx)
	go func() {
		for {
			time.Sleep(time.Second * 1)
			select {
			case <-ct.Done():
				fmt.Println("done")
				return
			default:
				fmt.Println("work")
				fmt.Println(ct.Deadline())
			}
		}
	}()
	time.Sleep(time.Second * 10)
	cancle()
	time.Sleep(time.Second * 1)
}

func main() {
	ctx := context.TODO()
	processHandler(ctx)
	fmt.Println("end")
}
