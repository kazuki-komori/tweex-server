package main

import (
	"fmt"
	"os"

	"github.com/kazuki-komori/tweex-server/database"
	"github.com/kazuki-komori/tweex-server/usecase"
	"github.com/kazuki-komori/tweex-server/web"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		fmt.Print("failed to create db=%w", err)
		os.Exit(1)
	}
	userRepo := database.NewUserRepository(*db)
	userUC := usecase.NewUserUsecase(userRepo)
	web.NewServer(userUC)
}
