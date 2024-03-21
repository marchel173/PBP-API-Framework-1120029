package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetUserById(id string) User {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Users WHERE id=?",
		id,
	)
	fmt.Println(id)
	if err != nil {
		log.Print(err)
	}

	var user User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Nama, &user.Alamat, &user.Email, &user.Password); err != nil {
			log.Print(err.Error())
		}
	}
	fmt.Println(user)
	return (user)
}

func InsertUser(newUser User) bool {
	db := connect()
	defer db.Close()

	_, errQuery := db.Exec("INSERT INTO users(nama, alamat, email, password) VALUES (?,?,?,?)",
		newUser.Nama,
		newUser.Alamat,
		newUser.Email,
		newUser.Password,
	)

	if errQuery == nil {
		return true
	} else {
		return false
	}
}

func UpdateUser(user User) bool {
	db := connect()
	defer db.Close()

	fmt.Print(user)
	_, errQuery := db.Exec("UPDATE users SET nama = ?, alamat = ?, email = ?, password = ? WHERE id = ?",
		user.Nama,
		user.Alamat,
		user.Email,
		user.Password,
		user.Id,
	)

	if errQuery == nil {
		return true
	} else {
		return false
	}
}

func DeleteUser(id string) bool {
	db := connect()
	defer db.Close()

	_, errQuery := db.Exec("DELETE FROM users WHERE id = ?",
		id,
	)

	if errQuery == nil {
		return true
	} else {
		return false
	}
}

func SendResponse(w http.ResponseWriter, r *http.Request, errMessage string, status int) {
	var response Response
	response.Status = status
	response.Message = errMessage
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
