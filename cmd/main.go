package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"tracker/internal/api/http"
	"tracker/internal/auth"
	authImpl "tracker/internal/auth/impl"
	"tracker/internal/base/database"
	"tracker/internal/config"
	"tracker/internal/team"
	teamImpl "tracker/internal/team/impl"
	"tracker/internal/user"
	userImpl "tracker/internal/user/impl"
)

func main() {
	cfg, err := config.LoadConfig()
	str, _ := json.MarshalIndent(cfg, "", "\t")
	fmt.Println(string(str))
	if err != nil {
		panic(err)
	}
	//rdb := database.NewRedisConnection(cfg.Redis)
	mySQlDB, err := database.NewMySQlConnection(cfg.MySQL)
	if err != nil {
		panic(err)
	}
	userRepository := userImpl.NewUserRepository(mySQlDB)
	teamRepository := teamImpl.NewTeamRepository(mySQlDB)

	userUC := userImpl.NewUserUC(userRepository)
	authUC := authImpl.NewAuthUC(userUC, cfg.JWTSecret, cfg.TokenTTL)
	teamUC := teamImpl.NewTeamUC(teamRepository, userUC)

	userHandler := user.NewUserHandler(userUC)
	authHandler := auth.NewAuthHandler(authUC)
	teamHandler := team.NewTeamHandler(teamUC)

	server := http.NewServer(cfg.Server, http.Handlers{
		User: userHandler,
		Auth: authHandler,
		Team: teamHandler,
	})

	server.Start()
	log.Println("server started")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}
	log.Println("Server exiting")
}
