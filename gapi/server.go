package gapi

import (
	"fmt"

	db "github.com/herzaparam/simplebank.git/db/sqlc"
	"github.com/herzaparam/simplebank.git/pb"
	"github.com/herzaparam/simplebank.git/token"
	"github.com/herzaparam/simplebank.git/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	return server, nil
}
