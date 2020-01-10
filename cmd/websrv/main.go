package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
)

var conf Config

func init() {
	conf = mustGetConfig()

	log.SetFlags(log.Ltime)
	log.SetPrefix(conf.ServiceName)
}

func main() {
	svc := &WebService{}

	rtr := httprouter.New()
	rtr.HandlerFunc("GET", "/", svc.IndexEndpoint)
	rtr.HandlerFunc("POST", "/login", svc.Login)
	rtr.HandlerFunc("POST", "/logout", svc.Logout)

	addr := fmt.Sprintf("%s:%d", conf.ServerHost, conf.ServerPort)
	srv := &http.Server{
		Addr:              addr,
		Handler:           rtr,
		ReadTimeout:       20 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("srv.ListenAndServe => %v", err)
		}
	}()

	log.Printf("Server listening at %q", addr)
	sg := <-sigint
	log.Printf("Interrupt received: %v", sg)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("srv.Shutdown => %v", err)
	} else {
		log.Println("Server gracefully exited.")
	}
}
