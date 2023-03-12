package main

import (
	"context"
	"fmt"
	"time"
)

// func main() {
// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	go func() {
// 		time.Sleep(time.Second * 4)
// 		cancel()
// 	}()

// 	bookhotel(ctx)
// }

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	bookhotel(ctx)
}

func bookhotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Tempo excedido para bookar o quarto")
	case <-time.After(time.Second * 11):
		fmt.Println("Quarto reservado com suceesso")
	}
}
