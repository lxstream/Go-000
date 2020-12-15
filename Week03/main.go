package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func serveDebug(ctx context.Context) error {
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
		case <-ctx.Done():
			s.Shutdown(ctx)
			fmt.Println("ctx done", ctx.Err())
		}
	}()

	return s.ListenAndServe()
}

func serveApp(ctx context.Context) error {
	s := http.Server{
		Addr:    "localhost:8001",
		Handler: nil,
	}

	go func() {
		select {
		case <-ctx.Done():
			s.Shutdown(ctx)
			fmt.Println("ctx done", ctx.Err())
		}

	}()

	return s.ListenAndServe()
}

func main() {

	ctx, cancal := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	defer cancal()

	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)

		for {
			fmt.Println("signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case sign := <-sig:
				fmt.Println(sign)
				cancal()
				return nil
			}
		}
	})

	g.Go(func() error {
		return serveDebug(ctx)
	})
	g.Go(func() error {
		return serveApp(ctx)
	})

	fmt.Println("http servers start , awaiting signal")

	if err := g.Wait(); err != nil {
		fmt.Println("error group return err:", err.Error())
	}

	fmt.Println("exiting")
}
