package main

import "os"

import "strconv"

// Config is derived from env vars.
type Config struct {
	ServiceName      string
	ServerHost       string
	ServerPort       int
	DatabaseHost     string
	DatabasePort     int
	DatabaseName     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseSSLMode  string
}

func mustGetConfig() Config {
	conf, err := getConfig()
	if err != nil {
		panic(err)
	}
	return conf
}

func getConfig() (Config, error) {
	conf := Config{}

	conf.ServiceName = os.Getenv("WEBSVC_NAME")
	if conf.ServiceName == "" {
		conf.ServiceName = "server"
	}

	conf.ServerHost = os.Getenv("WEBSVC_SRV_HOST")
	if conf.ServerHost == "" {
		conf.ServerHost = "127.0.0.1"
	}

	srvport := os.Getenv("WEBSVC_SRV_PORT")
	if srvport != "" {
		conf.ServerPort, _ = strconv.Atoi(srvport)
	} else {
		conf.ServerPort = 4200
	}

	conf.DatabaseHost = os.Getenv("WEBSVC_DB_HOST")
	if conf.DatabaseHost == "" {
		conf.DatabaseHost = "127.0.0.1"
	}

	dbport := os.Getenv("WEBSVC_DB_PORT")
	if dbport != "" {
		conf.DatabasePort, _ = strconv.Atoi(dbport)
	} else {
		conf.DatabasePort = 5432
	}

	conf.DatabaseName = os.Getenv("WEBSVC_DB_NAME")
	if conf.DatabaseName == "" {
		conf.DatabaseName = "postgres"
	}

	conf.DatabaseUsername = os.Getenv("WEBSVC_DB_USER")
	if conf.DatabaseUsername == "" {
		conf.DatabaseUsername = "postgres"
	}

	conf.DatabasePassword = os.Getenv("WEBSVC_DB_PASS")
	if conf.DatabasePassword == "" {
		conf.DatabasePassword = "postgres"
	}

	conf.DatabaseSSLMode = os.Getenv("WEBSVC_DB_SSLMODE")
	if conf.DatabaseSSLMode == "" {
		conf.DatabaseSSLMode = "disable"
	}

	return conf, nil
}
