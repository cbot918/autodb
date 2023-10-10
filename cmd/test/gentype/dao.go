package main

import (
	"database/sql"

	"github.com/cbot918/autodb/cmd/test/gentype/types"
	"github.com/cbot918/autodb/internal"
)

func TestRowBind(db *sql.DB, q string) error {
	g := types.T_goods{}
	err := internal.RowBind(db.QueryRow(q), &g)
	if err != nil {
		return err
	}

	internal.PrintJSON(g)

	return nil
}

func TestRowsBind(db *sql.DB, q string) error {
	gs := []types.T_goods{}
	rows, err := db.Query(q)
	if err != nil {
		return err
	}
	for rows.Next() {
		g := types.T_goods{}
		err = internal.RowsBind(rows, &g)
		if err != nil {
			return err
		}
		gs = append(gs, g)
	}

	internal.PrintJSON(gs)

	return nil
}
