package main

import (
	"flag"
	"collector/app"
	"os"

	"github.com/go-pg/migrations"
	"pg-ext"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	pg_ext.InitMigrationTableIfNeeded(app.Db(), app.Log())
	oldVersion, newVersion, err := migrations.Run(app.Db(), flag.Args()...)
	if err != nil {
		app.Log().Panic(err.Error())
	}
	if newVersion != oldVersion {
		app.Log().Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		app.Log().Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	app.Log().Print(pg_ext.MigrationUsageText)
	flag.PrintDefaults()
	os.Exit(2)
}
