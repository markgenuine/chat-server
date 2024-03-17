package main

import (
	"context"
	"flag"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/markgenuine/chat-server/internal/app"
)

func main() {
	flag.Parse()
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	if err := a.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
