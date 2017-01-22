package main

import (
	"database/sql"
	"os"

	"github.com/WeCanHearYou/wchy-api/context"
	"github.com/WeCanHearYou/wchy-api/handlers"
	"github.com/WeCanHearYou/wchy-api/services"
	"github.com/WeCanHearYou/wchy-api/util"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")
var ctx context.WchyContext
var settings context.WchySettings
var db *sql.DB
var buildtime string

func init() {
	log.Info("Application is starting...")
	db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))

	ctx = context.WchyContext{
		Health: services.NewPostgresHealthCheckService(db),
		Settings: context.WchySettings{
			BuildTime: buildtime,
		},
	}
}

func main() {
	address := ":" + util.GetEnvOrDefault("PORT", "3000")
	engine := handlers.GetMainEngine(ctx)
	engine.Run(address)
}