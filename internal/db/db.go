package db

import (
	"github.com/BzingaApp/user-svc/internal/genesis"
)

type DB struct {
	*genesis.Service
}

type PostgresDB struct {
	DB
}
