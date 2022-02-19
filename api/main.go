package main

import (
	"context"
	"getir-case/api/data"
	"getir-case/api/search"
	"getir-case/api/store/cache"
	"getir-case/api/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){

	mongoServer := new(search.MongoDb)
	util.Init()
	holder := cache.NewCacheProvider()
	dataHandler := data.New(holder)

	http.HandleFunc("/getAndSet", dataHandler.InMemory)
	http.HandleFunc("/searchDatabase", mongoServer.ServeMongo)

	httpServer := &http.Server{
		Addr: ":8080",
	}

	util.Info("Application started at 8080")

	go func() {
		if err := httpServer.ListenAndServe();
		err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("error handled: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Printf("stopped\n")
	}

	defer os.Exit(0)
}