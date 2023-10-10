package types

import "database/sql"

type Users struct {
	Id                            int            `db:"id"`
	USER                          sql.NullString `db:"USER"`
	Name                          string         `db:"name"`
	CURRENT_CONNECTIONS           int64          `db:"CURRENT_CONNECTIONS"`
	Email                         string         `db:"email"`
	TOTAL_CONNECTIONS             int64          `db:"TOTAL_CONNECTIONS"`
	MAX_SESSION_CONTROLLED_MEMORY int64          `db:"MAX_SESSION_CONTROLLED_MEMORY"`
	MAX_SESSION_TOTAL_MEMORY      int64          `db:"MAX_SESSION_TOTAL_MEMORY"`
}
