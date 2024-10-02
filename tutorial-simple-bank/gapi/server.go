package gapi

import (
	"fmt"

	db "github.com/tutorial/simple-bank/db/sqlc"
	"github.com/tutorial/simple-bank/pb"
	"github.com/tutorial/simple-bank/token"
	"github.com/tutorial/simple-bank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// GRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	return server, nil
}
