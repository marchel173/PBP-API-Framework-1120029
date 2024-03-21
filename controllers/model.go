package controllers

type User struct {
	Id       int    `form:"id" json:"id"`
	Nama     string `form:"nama" json:"nama"`
	Alamat   string `form:"alamat" json:"alamat"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type Response struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
}
