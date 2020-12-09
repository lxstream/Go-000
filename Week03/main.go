package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func serveDebug(ctx context.Context, stop <-chan bool) error {
	s := http.Server{
		Addr:    "localhost:8000",
		Handler: nil,
	}

	//模拟异常退出
	http.HandleFunc("/err", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("this is an error")
		s.Shutdown(ctx)
	})

	go func() {
		select {
		case <-stop:
			s.Shutdown(ctx)
			fmt.Println("stop signal")
		case <-ctx.Done():
			s.Shutdown(ctx)
			fmt.Println("ctx done", ctx.Err())
		}
	}()

	return s.ListenAndServe()
}

func serveApp(ctx context.Context, stop <-chan bool) error {
	s := http.Server{
		Addr:    "localhost:8001",
		Handler: nil,
	}

	go func() {
		select {
		case <-stop:
			s.Shutdown(ctx)
			fmt.Println("stop signal")
		case <-ctx.Done():
			s.Shutdown(ctx)
			fmt.Println("ctx done", ctx.Err())
		}

	}()

	return s.ListenAndServe()
}

func main() {
	sigs := make(chan os.Signal, 1)
	stop := make(chan bool)

	g, ctx := errgroup.WithContext(context.Background())

	signal.Notify(sigs)

	go func() {
		select {
		case sign := <-sigs:
			fmt.Println(sign)
			stop <- true
		case <-ctx.Done():
			fmt.Println("signal stop")
		}
	}()

	g.Go(func() error {
		return serveDebug(ctx, stop)
	})
	g.Go(func() error {
		return serveApp(ctx, stop)
	})

	fmt.Println("http servers start , awaiting signal")

	if err := g.Wait(); err != nil {
		fmt.Println("error group return err:", err.Error())
	}

	fmt.Println("exiting")
}
