package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1000 * time.Millisecond //10 or 1 * time.Millisecond

func main() {

	//並列処理Go routineの処理が終わるごとに100ミリ秒遅延させる
	_ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Make sure to cancel the context when main is done
	intStream := generator(_ctx)

	go func() {
		for i := 0; i < 50; i++ {
			num := <-intStream
			fmt.Println("Received:", num)
			time.Sleep(100 * time.Millisecond) // Simulate some processing time
		}
	}()

	//Contextで制御する
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("脑子进煎鱼了")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func generator(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
