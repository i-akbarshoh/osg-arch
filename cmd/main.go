package cmd

postgresDB := NewPostgres()

userRepo := user.NewRepository(postgresDB)
userService := user2.NewService(userRepo)
userUseCase := user3.NewUseCase(userService, userFileService)
userController := user4.NewController(userUseCase)

userController.GetAll