package router

import (
	"github.com/gorilla/mux"
	"github.com/hmrbcnto/go-net-http/infastructure/db/mongo/user_repo"
	handlers "github.com/hmrbcnto/go-net-http/services/handlers/user"
	"github.com/hmrbcnto/go-net-http/services/usecases"
	"go.mongodb.org/mongo-driver/mongo"
)

type router interface {
	InitRoutes(mux *mux.Router)
}

func InitializeRoutes(db *mongo.Client, mux *mux.Router) {
	// Generate http handlers
	userRepo := user_repo.NewRepo(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHttpHandler := handlers.NewUserHandler(userUsecase)

	routers := []router{userHttpHandler}

	for _, router := range routers {
		router.InitRoutes(mux)
	}
}
