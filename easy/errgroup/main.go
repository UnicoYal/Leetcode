package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			fmt.Println("first: context done")
		default:
			fmt.Println("first: started")
			time.Sleep(time.Second)
		}
		return nil
	})

	g.Go(func() error {
		fmt.Println("second: started")
		return errors.New("second: something went wrong")
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("third: context done")
		default:
			fmt.Println("third: started")
			time.Sleep(time.Second)
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
