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
	enums.POSTGRESQL_DB_URL: {
		"db url ",
		"postgres://nashaath.mohamed@bzinga.com:Nashaath@bz-dev-db-pvt-pgsql-central-india-1.postgres.database.azure.com:5432/bz_pg_dev",
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
	enums.SSH_HOST: {
		"bastion host",
		"20.219.153.186",
	},
	enums.SSH_USER: {
		"bastion host username",
		"nashaath.mohamed",
	},
	enums.SSH_PORT: {
		"bastion host ssh port",
		"22",
	},
	enums.SSH_PRIVATE_KEY_FILE_PATH: {
		"ssh private key location",
		"c:/users/Hiro/.ssh/github.com-hiroBzinga",
	},
}
