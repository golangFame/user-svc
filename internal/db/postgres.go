package db

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"github.com/BzingaApp/user-svc/enums"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
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

	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(fmt.Sprintf("%s:%s", conf.GetString(enums.POSTGRESQL_HOST), conf.GetString(enums.POSTGRESQL_PORT))),
		pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithUser(conf.GetString(enums.POSTGRESQL_USER)),
		pgdriver.WithPassword(conf.GetString(enums.POSTGRESQL_PASSWORD)),
		pgdriver.WithDatabase(conf.GetString(enums.POSTGRESQL_DB)),
		pgdriver.WithApplicationName(conf.GetString(enums.APP_NAME)),
		pgdriver.WithTimeout(10*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
		pgdriver.WithConnParams(map[string]interface{}{
			"timezone": conf.GetString(enums.TIMEZONE), //becomes set timezone to ...
		}),
	)
	sqldb := sql.OpenDB(pgconn)

	db = bun.NewDB(sqldb, pgdialect.New(), bun.WithDiscardUnknownColumns())

	if conf.GetString(enums.MODE) == "development" {
		db.AddQueryHook(&QueryHook{&database.Log})
	}
	_, err := db.Exec(fmt.Sprintf("set timezone to '%s'", conf.GetString(enums.TIMEZONE)))
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to connect to db due to - %s", err))
		return
	}

	return
}
