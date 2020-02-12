package dbinit

import (
	"database/sql"
)

func Init(db *sql.DB) (err error) {
	ddls := []string{foreignKeyOn, clientsDDL, accountsDDL, ATMsDDL, servicesDDL}
	for _, ddl := range ddls{
		_, err := db.Exec(ddl)
		if err != nil {
			return err
		}
	}
	return nil
}
