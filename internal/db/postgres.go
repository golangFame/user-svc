package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/BzingaApp/user-svc/enums"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type QueryHook struct {
	log *log.Logger
}

//goland:noinspection ALL
func (h *QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

//goland:noinspection ALL
func (h *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	h.log.Debugf("%v %v", time.Since(event.StartTime), event.Query)
	err := event.Err
	if err != nil {
		h.log.Error("err:- ", err)
	}
}

func newPostgressDB(database *DB) (db *bun.DB) {

	conf := database.Conf

	if conf.GetString(enums.MODE) == enums.DEVELOPMENT {
		sshTunnelTheSQL(conf)
	}

	if sqlDB, err := sql.Open("postgres+ssh", conf.GetString(enums.POSTGRESQL_DB_URL)); err == nil {

		log.Info("Successfully connected to the sqlDB")

		err = sqlDB.Ping()

		if err != nil {
			log.Errorln("unable to ping sqlDB")
			panic(err)
		}
		log.Info("successfully pinged the DB")
		db = bun.NewDB(sqlDB, pgdialect.New(), bun.WithDiscardUnknownColumns())
		return
	} else {
		panic(fmt.Sprintf("Failed to connect to the sqlDB: %s\n", err.Error()))
	}

	if conf.GetString(enums.MODE) == "development" {
		db.AddQueryHook(&QueryHook{&database.Log})
	}
	_, err := db.Exec(fmt.Sprintf("set timezone to '%s'", conf.GetString(enums.TIMEZONE)))
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to connect to sqlDB due to - %s", err))
		return
	}
	return
}
