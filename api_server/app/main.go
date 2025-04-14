package main

import (
	"app/db"
	"app/handlers"
	"app/middlewares"
	apis "app/openapi"
	"app/services"
	"app/utils/routers"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()

	dbCon := db.Init()
	e := echo.New()

	// NOTE: service層のインスタンス
	authService := services.NewAuthService(dbCon)
	todoService := services.NewTodoService(dbCon)

	// NOTE: Handlerのインスタンス化
	authHandler := handlers.NewAuthHandler(authService)
	todosHandler := handlers.NewTodosHandler(todoService)
	mainHandler := handlers.NewMainHandler(authHandler, todosHandler)
	
	mainStrictHandler := apis.NewStrictHandler(mainHandler, []apis.StrictMiddlewareFunc{middlewares.AuthMiddleware})

	appliedMiddlewareEcho := routers.ApplyMiddlewares(e)
	apis.RegisterHandlers(appliedMiddlewareEcho, mainStrictHandler)

	appliedMiddlewareEcho.Logger.Fatal(appliedMiddlewareEcho.Start(":" + os.Getenv("SERVER_PORT")))
}

func loadEnv() {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}
	godotenv.Load(envFilePath)
}
