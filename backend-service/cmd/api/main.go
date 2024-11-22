package main

import (
	"flag"
	"ongambl/internal/jsonlog"
	"ongambl/internal/repository"
	"ongambl/internal/repository/dbrepo"
	"os"
	"strings"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type config struct {
	port int
	env  string
	db   struct {
		maxOpensConns int
		maxIdleConns  int
		maxIdleTime   string
	}
	cors struct {
		trustedOrigins []string
	}
	jwt struct {
		secret       string
		issuer       string
		audience     string
		cookieDomain string
	}
}

type application struct {
	auth   Auth
	cfg    config
	logger *jsonlog.Logger
	wg     sync.WaitGroup
	Domain string
	DSN    string
	DB     repository.DatabaseRepo
	Rabbit *amqp.Connection
}

func main() {
	var cfg config
	var app application

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "env|dev|main|")

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=ongambl sslmode=disable timezone=UTC connect_timeout=5", "Database Source Name")
	flag.StringVar(&app.Domain, "domain", "localhost", "domain")
	flag.StringVar(&cfg.jwt.secret, "jwt-secret", "verysecret", "signing secret")
	flag.StringVar(&cfg.jwt.issuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&cfg.jwt.audience, "jwt-audience", "example.com", "signing audience")
	flag.StringVar(&cfg.jwt.cookieDomain, "cookie-domain", "localhost", "cookie domain")

	flag.IntVar(&cfg.db.maxOpensConns, "db-max-opens-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max edle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connections idle time")

	flag.Func("cors-trusted-origins", "Trusted CORS origins http://localhost:3000", func(s string) error {
		cfg.cors.trustedOrigins = strings.Fields(s)
		return nil
	})
	cfg.env = "docker"
	if cfg.env == "docker" {
		app.DSN = "host=postgres port=5432 user=postgres password=password dbname=ongambl sslmode=disable timezone=UTC connect_timeout=5"
	}

	flag.Parse()
	cfg.cors.trustedOrigins = append(cfg.cors.trustedOrigins, "http://localhost:3000")

	app.logger = jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	app.cfg = cfg

	rabbitConn, err := app.connectToRabbit()
	if err != nil {
		app.logger.PrintFatal(err, map[string]string{})
		return
	}
	defer rabbitConn.Close()

	app.Rabbit = rabbitConn

	app.auth = Auth{
		Issuer:        app.cfg.jwt.issuer,
		Audience:      app.cfg.jwt.audience,
		Secret:        app.cfg.jwt.secret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "__Auth-refresh_token",
		CookieDomain:  app.cfg.jwt.cookieDomain,
	}

	conn, err := app.connectToDB()
	if err != nil {
		app.logger.PrintFatal(err, map[string]string{})
		return
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Conncetion().Close()

	err = app.server()
	if err != nil {
		app.logger.PrintFatal(err, nil)
	}

}
