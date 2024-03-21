package models

type User struct {
	Id       int    `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Age      int    `form:"age" json:"age"`
	Address  string `form:"address" json:"address"`
	Password string `form:"password" json:"password"`
}

type Login struct {
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
}

type Register struct {
	Name      string `form:"name" json:"name"`
	Age       int    `form:"age" json:"age"`
	Address   string `form:"address" json:"address"`
	Password1 string `form:"password1" json:"password1"`
	Password2 string `form:"password2" json:"password2"`
}

type Update struct {
	Name    string `form:"name" json:"name"`
	Age     int    `form:"age" json:"age"`
	Address string `form:"address" json:"address"`
}

type Delete struct {
	Id int `form:"id" json:"id"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
