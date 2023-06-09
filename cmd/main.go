package main

import (
	"log"
	"net/http"

	project4 "github.com/i-akbarshoh/osg-arch/internal/controller/project"
	comment1 "github.com/i-akbarshoh/osg-arch/internal/repository/comment"
	comment2 "github.com/i-akbarshoh/osg-arch/internal/service/comment"
	user4 "github.com/i-akbarshoh/osg-arch/internal/controller/user"
	"github.com/i-akbarshoh/osg-arch/internal/pkg/config"
	_ "github.com/i-akbarshoh/osg-arch/internal/pkg/config"
	"github.com/i-akbarshoh/osg-arch/internal/pkg/repository/postgres"
	"github.com/i-akbarshoh/osg-arch/internal/pkg/router"
	"github.com/i-akbarshoh/osg-arch/internal/repository/project"
	task1 "github.com/i-akbarshoh/osg-arch/internal/repository/task"
	user1 "github.com/i-akbarshoh/osg-arch/internal/repository/user"
	project2 "github.com/i-akbarshoh/osg-arch/internal/service/project"
	task2 "github.com/i-akbarshoh/osg-arch/internal/service/task"
	user2 "github.com/i-akbarshoh/osg-arch/internal/service/user"
	project3 "github.com/i-akbarshoh/osg-arch/internal/usecase/project"
	user3 "github.com/i-akbarshoh/osg-arch/internal/usecase/user"
	attendance1 "github.com/i-akbarshoh/osg-arch/internal/repository/attendance"
	attendance2 "github.com/i-akbarshoh/osg-arch/internal/service/attendance"
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
	taskRepo := task1.New(db)
	taskService := task2.New(taskRepo)
	attenRepo := attendance1.New(db)
	attenService := attendance2.New(attenRepo)
	userUseCase := user3.NewUseCase(userService, attenService)
	userController := user4.NewController(userUseCase)
	projectRepo := project.New(db)
	projectService := project2.New(projectRepo)
	commentRepo := comment1.New(db)
	commentService := comment2.New(commentRepo)
	projectUseCase := project3.New(projectService, taskService, commentService)
	projectController := project4.New(projectUseCase)

	engine := router.New(userController, projectController)
	if err := http.ListenAndServe(config.C.Server.Host+":"+config.C.Server.Port, engine); err != nil {
		log.Println(err)
	}
}
