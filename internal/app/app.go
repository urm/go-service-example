package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/urm/go-service-example/internal/adapters/msg"
	"github.com/urm/go-service-example/internal/adapters/storage"
	"github.com/urm/go-service-example/internal/adapters/webapi"
	"github.com/urm/go-service-example/internal/domain/tasks"
)

func Run() {
	storage := storage.New("connstring")
	sender := msg.New()
	tasks := tasks.New(storage, sender)
	server := webapi.New(tasks)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: server.Router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	<-ctx.Done()

	log.Println("Shutdown Server ...")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
