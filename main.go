package main

import (
	"encoding/json"
	"net/http"

	controllers "github.com/Martini/controllers"
	models "github.com/Martini/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

func main() {
	m := martini.Classic()

	m.Get("/user/:id", func(w http.ResponseWriter, r *http.Request, p martini.Params) {
		var id string = p["id"]
		var user models.User = controllers.GetUserById(id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	m.Post("/user/login", binding.Bind(models.Login{}), func(w http.ResponseWriter, r *http.Request, cekLogin models.Login) {
		var user models.User = controllers.GetUserByLogin(cekLogin)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	m.Post("/user/register", binding.Bind(models.Register{}), func(w http.ResponseWriter, r *http.Request, cekRegister models.Register) {
		var cek bool
		if cekRegister.Password1 == cekRegister.Password2 {
			cek = controllers.AddUser(cekRegister)
			if cek {
				controllers.PrintSuccess(200, "Registered", w)
			} else {
				controllers.PrintError(400, "Failed to Register", w)
			}
		} else {
			controllers.PrintError(400, "Failed to Register", w)
		}
	})

	m.Put("/user/edit/:id", binding.Bind(models.Update{}), func(w http.ResponseWriter, r *http.Request, p martini.Params, update models.Update) {
		var id string = p["id"]

		var cek bool = controllers.UpdateUser(id, update)
		if cek {
			controllers.PrintSuccess(200, "Updated", w)
		} else {
			controllers.PrintError(400, "Failed to Delete", w)
		}
	})

	m.Delete("/user/delete/:id", binding.Bind(models.Delete{}), func(w http.ResponseWriter, r *http.Request, p martini.Params) {
		var id string = p["id"]
		var cek bool = controllers.DeleteUser(id)
		if cek {
			controllers.PrintSuccess(200, "Deleted", w)
		} else {
			controllers.PrintError(400, "Failed to Delete", w)
		}
	})

	http.Handle("/", m)
	// m.Run() // port :3000
	m.RunOnAddr(":8080") // port :8080
}
