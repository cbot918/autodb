package main

import (
	"autodb/internal"
	"fmt"
	"log"
)

func main() {

	var err error

	// init config
	cfg, err := internal.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// createdb
	err = internal.CreateDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// initial db pool
	db, err := internal.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// migrate sql to database
	err = internal.Migrate(cfg, db)
	if err != nil {
		log.Fatal(err)
	}

	// initial DBMetadata: get result dbm
	dbm, err := internal.NewDBMetadata(cfg, db)
	if err != nil {
		log.Fatal(err)
	}

	internal.PrintJSON(dbm)

	fmt.Println("DBMetadata read success!")
}