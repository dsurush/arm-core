package dbupdate

import (
	"database/sql"
	"log"
)

func AddClient(name, surname, login, password string, db *sql.DB) (err error){
	locked := true
	_, err = db.Exec(addClientDML, name, surname, login, password, locked)
	if err != nil {
		log.Fatalf("Пользователь недобавлен: %s", err)
	}
	return nil
}