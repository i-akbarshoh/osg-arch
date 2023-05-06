package main

import (
	user4 "github.com/i-akbarshoh/osg-arch/internal/controller/user"
	"github.com/i-akbarshoh/osg-arch/internal/pkg/config"
	_ "github.com/i-akbarshoh/osg-arch/internal/pkg/config"
	"github.com/i-akbarshoh/osg-arch/internal/pkg/repository/postgres"
	"github.com/i-akbarshoh/osg-arch/internal/pkg/router"
	user1 "github.com/i-akbarshoh/osg-arch/internal/repository/user"
	user2 "github.com/i-akbarshoh/osg-arch/internal/service/user"
	user3 "github.com/i-akbarshoh/osg-arch/internal/usecase/user"
	"log"
	"net/http"
)

func main() {
	db := postgres.NewDB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()
	userRepo := user1.NewRepository(db)
	userService := user2.NewService(userRepo)
	userUseCase := user3.NewUseCase(userService)
	userController := user4.NewController(userUseCase)

	engine := router.New(userController)
	if err := http.ListenAndServe(config.C.Server.Host+":"+config.C.Server.Port, engine); err != nil {
		log.Println(err)
	}
}
