package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go_blog/internal/server"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	newServer := server.NewServer()
	go func() {
		newServer.Run(ctx)
	}()
	for range signalChan {
		_ = newServer.Stop(ctx)
		cancel()
	}
	fmt.Printf("newServer stopped")
}
