package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/pyotarou/simplebank/api"
	db "github.com/pyotarou/simplebank/db/sqlc"
	"github.com/pyotarou/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)

	go func() {
		err = server.Start(config.ServerAddress)
		if err != nil {
			log.Fatal("cannot start server:", err)
		}
	}()

	<-ctx.Done()
}
