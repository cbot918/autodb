package main

import (
	"fmt"
	"log"

	"github.com/cbot918/autodb/internal"
	"github.com/cbot918/autodb/internal/gentype"
)

func main() {
	var err error

	cfg, db, cfgErr, dbErr := internal.Init()
	if cfgErr != nil {
		fmt.Println("load config error")
		return
	}
	if dbErr != nil {
		fmt.Println("db open error")
		return
	}
	defer db.Close()

	// initial DBMetadata: get result dbm
	dbm, err := internal.NewDBMetadata(cfg, db)
	if err != nil {
		log.Fatal(err)
	}

	internal.PrintJSON(dbm)

	gt := gentype.NewGentype(cfg, dbm)

	err = gt.InitContent()
	if err != nil {
		log.Fatal(err)
	}

	// err = gt.Create()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	q1 := "select * from t_goods order by goods_id limit 1;"
	err = TestRowBind(db, q1)
	if err != nil {
		log.Fatal(err)
	}

	q2 := "select * from t_goods order by goods_id limit 5;"
	err = TestRowsBind(db, q2)
	if err != nil {
		log.Fatal(err)
	}

}
