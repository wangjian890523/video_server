package dbops

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginname string, pwd string) error {
	stmIns, err = dbConn.Prepare("INSET INTO　users(login_name, pwd) VALUES (?,?)")
	if err != nil {
		return err
	}

	stmIns.Exec()
	stmIns.Close()
	return nil
}

func GetUserCredential(loginname string) (string, error) {
	stmOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		retturn "", err
	}
	var pwd string
	stmOut.QueryRow(loginName).Scan(&pwd)
	stmOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name =?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	stmDelD.Exec(loginName, pwd)
	stmDel.Close()

	return nil
}
