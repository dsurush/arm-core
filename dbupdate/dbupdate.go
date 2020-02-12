package dbupdate

import (
	"database/sql"
	"fmt"
	"github.com/dsurush/arm-core/dbupdate/models"
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

func GetAllClients(db *sql.DB) (clients []models.Client, err error){
	rows, err := db.Query(getAllClients)
	if err != nil {
		log.Fatalf("1 wrong")
		return nil, err
	}

	defer func() {
		if innerErr := rows.Close(); innerErr != nil {
			clients = nil
		}
	}()

	for rows.Next(){
		client := models.Client{}
		err = rows.Scan(&client.ID, &client.Name, &client.Surname, &client.NumberPhone, &client.Login, &client.Password, &client.Locked)
		if err != nil {
			log.Fatalf("2 wrong")
			return nil, err
		}
		clients = append(clients, client)
	}
	if rows.Err() != nil {
		log.Fatalf("3 wrong")
		return nil, rows.Err()
	}
	return clients, nil
}

func GetAllAccounts(db *sql.DB) (accounts []models.Account, err error){
	rows, err := db.Query(getAllAccounts)
	if err != nil {
		log.Fatalf("1 wrong")
		return nil, err
	}

	defer func() {
		if innerErr := rows.Close(); innerErr != nil {
			accounts = nil
		}
	}()

	for rows.Next(){
		account := models.Account{}
		err = rows.Scan(&account.ID, &account.UserId, &account.Name, &account.AccountNumber, &account.Locked)
		if err != nil {
			log.Fatalf("2 wrong")
			return nil, err
		}
		accounts = append(accounts, account)
	}
	if rows.Err() != nil {
		log.Fatalf("3 wrong")
		return nil, rows.Err()
	}
	return accounts, nil
}

func GetAllATMs(db *sql.DB) (ATMs []models.ATM, err error){
	rows, err := db.Query(getAllATMs)
	if err != nil {
		log.Fatalf("1 wrong")
		return nil, err
	}

	defer func() {
		if innerErr := rows.Close(); innerErr != nil {
			ATMs = nil
		}
	}()

	for rows.Next(){
		ATM := models.ATM{}
		err = rows.Scan(&ATM.ID, &ATM.Name, &ATM.Locked)
		if err != nil {
			log.Fatalf("2 wrong")
			return nil, err
		}
		ATMs = append(ATMs, ATM)
	}
	if rows.Err() != nil {
		log.Fatalf("3 wrong")
		return nil, rows.Err()
	}
	return ATMs, nil
}

func Test() error {
	return nil
}