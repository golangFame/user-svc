package config

import "github.com/BzingaApp/user-svc/enums"

type argvMeta struct {
	desc       string
	defaultVal string
}

var confList = map[string]argvMeta{

	enums.APP_NAME: {
		"app name",
		"user-svc",
	},
	enums.ENV: {
		"environment",
		enums.DEV,
	},
	enums.MODE: {
		"app run configuration",
		enums.DEVELOPMENT,
	},
	enums.PORT: {
		"app listen port",
		"1604",
	},
	enums.POSTGRESQL_DB: {
		"postgresql db name",
		"bz_pg_dev",
	},
	enums.POSTGRESQL_HOST: {
		"postgresql host",
		"localhost",
	},
	enums.POSTGRESQL_PORT: {
		"postgresql port",
		"5432",
	},
	enums.POSTGRESQL_USER: {
		"postgresql username",
		"nashaath.mohamed@bzinga.com",
	},
	enums.POSTGRESQL_PASSWORD: {
		"postgresql password",
		"Nashaath",
	},
	enums.REDIS_SERVER: {
		"redis server",
		"localhost:6379",
	},
	enums.TIMEZONE: {
		"timezone to be used",
		"Asia/Calcutta",
	},
}
