package gapi

import (
	"fmt"

	db "github.com/pyotarou/simplebank/db/sqlc"
	"github.com/pyotarou/simplebank/pb"
	"github.com/pyotarou/simplebank/token"
	"github.com/pyotarou/simplebank/util"
)

type Server struct {
	pb.UnimplementedSimplebankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
