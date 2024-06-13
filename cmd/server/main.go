package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/yuki-maruyama/learn-gh-action/router"
	"github.com/yuki-maruyama/learn-gh-action/util"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	logger := slog.New(util.NewLogHandler())
	slog.SetDefault(logger)

	r := router.NewRouter()

	server := &http.Server{
		Addr:    ":8888",
		Handler: r,
	}
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	log.Println("server start running at :8888")
	log.Fatal(server.ListenAndServe())
}
