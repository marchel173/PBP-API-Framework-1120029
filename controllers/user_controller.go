package controllers

import (
	"log"
	"strconv"

	m "github.com/Martini/models"
)

func GetUserById(id string) m.User {
	db := connect()
	defer db.Close()

	id_int, _ := strconv.Atoi(id)

	rows, err := db.Query("SELECT * FROM Users WHERE id=?",
		id_int,
	)
	if err != nil {
		log.Print(err)
	}

	var user m.User
	for rows.Next() {
		if err := rows.Scan(&user.Id,
			&user.Name,
			&user.Age,
			&user.Address,
			&user.Password); err != nil {
			log.Print(err.Error())
		}
	}
	return (user)
}

func GetUserByLogin(cekLogin m.Login) m.User {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE name=? AND password=?", cekLogin.Name, cekLogin.Password)
	if err != nil {
		log.Print(err)
	}

	var user m.User
	for rows.Next() {
		if err := rows.Scan(&user.Id,
			&user.Name,
			&user.Age,
			&user.Address,
			&user.Password); err != nil {
			log.Print(err.Error())
		}
	}
	return user
}

func AddUser(cekRegister m.Register) bool {
	db := connect()
	defer db.Close()

	_, errQuery := db.Exec("INSERT INTO users (name, age, address, password) values (?,?,?,?)", cekRegister.Name, cekRegister.Age, cekRegister.Address, cekRegister.Password1)

	if errQuery == nil {
		return true
	} else {
		return false
	}
}

func UpdateUser(id string, newData m.Update) bool {
	db := connect()
	defer db.Close()

	id_int, _ := strconv.Atoi(id)

	result, errQuery := db.Exec("UPDATE users SET name = ?, age = ?, address = ? WHERE id = ?", newData.Name, newData.Age, newData.Address, id_int)
	num, _ := result.RowsAffected()
	if num != 0 {
		if errQuery == nil {
			return true
		} else {
			return false
		}
	}
	return false
}

func DeleteUser(id string) bool {
	db := connect()
	defer db.Close()

	id_int, _ := strconv.Atoi(id)

	result, errQuery := db.Exec("DELETE FROM users WHERE id = ?", id_int)
	num, _ := result.RowsAffected()
	if num != 0 {
		if errQuery == nil {
			return true
		} else {
			return false
		}
	}
	return false
}
