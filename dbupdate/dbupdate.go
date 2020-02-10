package dbupdate

import (
	"database/sql"
	"fmt"
	"log"
)

func AddClient(name, surname, login, password, phoneNumber string, db *sql.DB) (err error){
	locked := true
	_, err = db.Exec(addClientDML, name, surname, login, password, phoneNumber, locked)
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

func AddAccount(user_id int64, name string, locked bool, db *sql.DB) (err error){
	var count int
	//Number = 1_000_000_000_000_000
	err = db.QueryRow(`select count(*) from accounts`).Scan(&count)
	if err != nil {
		fmt.Errorf("cant %e", err)
		return err
	}
	var accountNumber int64
	var lastAccountNumber int64
	accountNumber = 1_000_000_000_000_000
	if count != 0 {
		err := db.QueryRow(`select max(accountNumber) from accounts`).Scan(&lastAccountNumber)
		if err != nil {
			fmt.Errorf("cant find last Account Number %e", err)
			return err
		}
		accountNumber = lastAccountNumber + 1
	}
	_, err = db.Exec(addAccount, user_id, name, accountNumber, locked)
	if err != nil {
		fmt.Errorf("cant insert %e", err)
		return err
	}

	fmt.Println("Success")
	return nil
}

func AddService(serviceName string, price int64, db *sql.DB) (err error){
	_, err = db.Exec(addService, serviceName, price)
	if err != nil {
		fmt.Errorf("Error in %s, err: %e", addService, err)
		return err
	}
	return nil
}

func Test() error {
	return nil
}