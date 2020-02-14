package dbupdate

import (
	"database/sql"
	"fmt"
	"github.com/dsurush/arm-core/dbupdate/cmodels"
	"log"
)

func AddClient(name, surname, login, password, phoneNumber string, db *sql.DB) (err error){
	locked := true
	password = makeHash(password)
	_, err = db.Exec(addClientDML, name, surname, login, password, phoneNumber, locked)
	if err != nil {
		log.Fatalf("Пользователь недобавлен: %s", err)
		return err
	}
	return nil
}

func AddATM(address string, locked bool, db *sql.DB) (err error){
	_, err = db.Exec(addATM, address, locked)
	if err != nil {
		log.Fatalf("Запись недобавлена: %e", err)
 		return err
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
			fmt.Errorf("cant find last AccountWithUserName Number %e", err)
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

func GetAllClients(db *sql.DB) (clients []cmodels.Client, err error){
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
		client := cmodels.Client{}
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

func GetAllAccounts(db *sql.DB) (accounts []cmodels.AccountWithUserName, err error){
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
		account := cmodels.AccountWithUserName{}
		err = rows.Scan(&account.Account.ID, &account.Account.UserId, &account.Account.Name, &account.Account.AccountNumber, &account.Account.Locked, &account.Client.ID, &account.Client.Name, &account.Client.Surname, &account.Client.NumberPhone, &account.Client.Login, &account.Client.Password, &account.Client.Locked)
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

func GetAllATMs(db *sql.DB) (ATMs []cmodels.ATM, err error){
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
		ATM := cmodels.ATM{}
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
/////////////////////----------CLIENT---------//////////////////////
func QueryError(text string) (err error){
	return fmt.Errorf(text)
}

func Login(login, password string, db *sql.DB) (loginPredicate bool, err error){
	var dbLogin, dbPassword string
	err = db.QueryRow(loginSQL, login).Scan(&dbLogin, &dbPassword)
	if err != nil {
//		fmt.Printf("%s, %e\n", loginSQL, err)
		return false, err
	}
	err = QueryError("Несовпадение пароля")
	if makeHash(password) != dbPassword {
		//fmt.Println(makeHash(password), " ", dbPassword)
		return true, err
	}
	//fmt.Println(makeHash(password), " ", dbPassword)
	return true, nil
}

func SearchByLogin(login string, db *sql.DB) (id int64, surname string){
	err := db.QueryRow(searchClientByLogin, login).Scan(&id, &surname)
	if err != nil {
		log.Fatalf("Ошибка в %s", searchClientByLogin)
	}
	return id, surname
}

func SearchAccountById(id int64, db *sql.DB) (Accounts []cmodels.AccountForUser, err error){
	var account cmodels.AccountForUser

	rows, err := db.Query(searchAccountByIDSql, id)
	if err != nil {
		fmt.Errorf("ne chitayutsya %e\n", err)
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&account.ID, &account.Name, &account.AccountNumber, account.Locked)
		Accounts = append(Accounts, account)
	}
	return Accounts, nil
}

func Test() error {
	return nil
}