package main

import (
	"awesomeProject/controllers"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"
)

func main() {
	conn, err := pgx.Connect(context.Background(),
		"postgres://idjbfebr:ONurXwA2nXGWBnhcMmYNiGk75zMBawJa@abul.db.elephantsql.com:5432/idjbfebr")
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to connect to database")
	}
	defer conn.Close(context.Background())

	r := gin.Default()

	r.LoadHTMLGlob("view/templates/*")
	controllers.InitializeControllers(r)

	log.Info().Msg("running server")
	err = r.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to run server")
	}
}
