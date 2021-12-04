package main

import (
	"context"
	"fmt"
	"github.com/baqtiste/fizzbuzz/db"
	"github.com/baqtiste/fizzbuzz/handler"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	database, err := db.Initialize(dbUser, dbPassword, dbName)

	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	defer database.Connection.Close()

	httpHandler := handler.NewHandler(database)
	server := &http.Server{Handler: httpHandler}

	go func() {
		server.Serve(listener)
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

func Fizzbuzz(int1 int, int2 int, limit int, str1 string, str2 string) string {
	var output string

	for i:=1; i<=limit; i++ {
		switch {
		case i % (int1*int2) == 0:
			output += str1+str2
		case i % int1 == 0:
			output += str1
		case i % int2 == 0:
			output += str2
		default:
			output += strconv.Itoa(i)
		}

		if i != limit {
			output += ","
		}
	}

	return output
}