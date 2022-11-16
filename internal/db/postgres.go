package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/BzingaApp/user-svc/enums"
	"github.com/oiime/logrusbun"
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
	//h.log.Infof(" %v %v", time.Since(event.StartTime), event.Query) //FIXME doesn't look nice
	fmt.Printf("[BUN]%v %v\n", time.Since(event.StartTime), event.Query)
	err := event.Err
	if err != nil {
		h.log.Errorln("query failed")
		h.log.Errorln(err)
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
	} else {
		panic(fmt.Sprintf("Failed to connect to the sqlDB: %s\n", err.Error()))
	}

	if conf.GetString(enums.MODE) == enums.DEVELOPMENT {

		//bundebug.NewQueryHook(bundebug.WithVerbose(true))

		db.AddQueryHook(&QueryHook{&database.Log})
	} else {
		db.AddQueryHook(logrusbun.NewQueryHook(logrusbun.QueryHookOptions{
			Logger:          &database.Log,
			LogSlow:         time.Millisecond * 40,
			ErrorLevel:      log.ErrorLevel,
			SlowLevel:       log.WarnLevel,
			MessageTemplate: "{{.Operation}}[{{.Duration}}]: {{.Query}}",
			ErrorTemplate:   "{{.Operation}}[{{.Duration}}]: {{.Query}}: {{.Error}}",
		}))
		//db.AddQueryHook(bundebug.NewQueryHook()) // only prints failed queries
	}
	_, err := db.Exec(fmt.Sprintf("set timezone to '%s'", conf.GetString(enums.TIMEZONE)))
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to connect to sqlDB due to - %s", err))
		return
	}
	return
}
