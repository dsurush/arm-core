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

func AddATM(address string, locked bool, db *sql.DB) (err error){
	_, err = db.Exec(addATM, address, locked)
	if err != nil {
		log.Fatalf("Запись недобавлена: %e", err)
 		//return err
	}
	return nil
}