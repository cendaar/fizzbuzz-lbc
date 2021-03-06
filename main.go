package main

import (
	"context"
	"fmt"
	"github.com/cendaar/fizzbuzz/db"
	"github.com/cendaar/fizzbuzz/handler"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	client, err := db.Initialize()
	if err != nil {
		log.Fatalln(err)
	}

	httpHandler := handler.NewHandler(client)
	server := &http.Server{Handler: httpHandler}

	go func() {
		_ = server.Serve(listener)
	}()

	defer Stop(server)

	log.Printf("Started server on %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}