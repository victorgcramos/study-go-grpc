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

	"github.com/gorilla/mux"
	"github.com/victorgcramos/zi/gateway/router"
	"github.com/victorgcramos/zi/pkg/config"
	"github.com/victorgcramos/zi/pkg/user"
)

func (g *Gateway) handleUserVersion(w http.ResponseWriter, _ *http.Request) {
	b, err := g.UserClient.Version(context.Background())
	if err != nil {
		router.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	router.RespondWithJSON(w, 200, b)
}

type Gateway struct {
	UserClient *user.ServiceClient
}

func (g *Gateway) setupRouter() *mux.Router {
	r := router.NewRouter()
	r.AddRoute("/version", g.handleUserVersion, http.MethodGet)

	return r.GetRouter()
}

func setupGateway() error {
	log.Printf("main: starting HTTP router")
	//	Setup config
	// TODO: Load config for each router
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	// Setup user client
	c := user.InitServiceClient(&cfg)
	g := Gateway{UserClient: c}

	//	Setup HTTP Server
	r := g.setupRouter()
	//	Start Server
	listen := make(chan error)
	go func() {
		s := &http.Server{
			// TODO: Use correct env vars
			Addr:         "127.0.0.1:" + cfg.Port,
			Handler:      r,
			ReadTimeout:  100000 * time.Second,
			WriteTimeout: 100000 * time.Second,
		}
		fmt.Println("Start of day")
		listen <- s.ListenAndServe()
	}()

	// Tell user we are ready to go.

	// Setup OS signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGINT)
	for {
		select {
		case sig := <-sigs:
			fmt.Printf("Terminating with %v", sig)
			goto done
		case err := <-listen:
			fmt.Printf("listen error: %v", err)
			goto done
		}
	}
done:
	return nil
}

func main() {
	err := setupGateway()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
