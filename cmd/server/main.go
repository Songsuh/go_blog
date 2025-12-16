package main

import (
	"context"
	"fmt"
	"github.com/Songsuh/go_blog/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server := server.NewServer()
	go func() {
		server.Run(ctx)
	}()
	for range signalChan {
		server.Stop(ctx)
		cancel()
	}
	fmt.Printf("server stopped")
}
