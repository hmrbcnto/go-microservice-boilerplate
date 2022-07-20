package server

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
	router "github.com/hmrbcnto/go-net-http/app/server/routes"
	"github.com/hmrbcnto/go-net-http/helpers"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServerInterface interface {
	ListenAndServe(port string) error
}

type Server struct {
	Db  *mongo.Client
	Mux *mux.Router
}

func New(db *mongo.Client) Server {
	mux := mux.NewRouter()
	mux.Use(helpers.LogRequestHandler)
	return Server{
		Db:  db,
		Mux: mux,
	}
}

// Creating server
func (srv *Server) ListenAndServe(port string) error {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	serverError := make(chan error, 1)

	router.InitializeRoutes(srv.Db, srv.Mux)

	httpServer := http.Server{
		Addr:         port,
		Handler:      srv.Mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		fmt.Printf("Server starting at port: %v", port)
		serverError <- httpServer.ListenAndServe()
	}()

	select {
	case <-serverError:
		log.Println("Error")
		return fmt.Errorf("Server error")
	case <-shutdown:
		log.Println("starting graceful server shutdown")
		defer log.Println("server has shutdown successfully")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			log.Println("server graceful shutdown is not successful")
			httpServer.Close()
			return err
		}
	}
	return nil
}
