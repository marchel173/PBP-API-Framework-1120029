package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	controllers "Eksplorasi-Framework-API-Go-Martini/controllers"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

func main() {
	m := martini.Classic()
	// GET USER BY ID
	m.Get("/user/:id", func(w http.ResponseWriter, r *http.Request, p martini.Params) {
		var user controllers.User

		user = controllers.GetUserById(p["id"])
		controllers.SendResponse(w, r, "Get user sukses!", 200)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})
	// POST USER
	m.Post("/user", binding.Bind(controllers.User{}), func(w http.ResponseWriter, r *http.Request, newUser controllers.User) {
		var check bool
		check = controllers.InsertUser(newUser)
		if check {
			controllers.SendResponse(w, r, "Insert sukses!", 200)
		} else {
			controllers.SendResponse(w, r, "Insert gagal...", 400)
		}
	})
	// PUT USER / UPDATE USER
	m.Put("/user/:id", binding.Bind(controllers.User{}), func(w http.ResponseWriter, r *http.Request, userUpdate controllers.User, p martini.Params) {
		var check bool
		userUpdate.Id, _ = strconv.Atoi(p["id"])
		check = controllers.UpdateUser(userUpdate)
		if check {
			controllers.SendResponse(w, r, "Update sukses!", 200)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(userUpdate)
		}
	})
	// DELETE USER
	m.Delete("/user/:id", binding.Bind(controllers.User{}), func(w http.ResponseWriter, r *http.Request, p martini.Params) {
		var check bool
		var id = p["id"]
		check = controllers.DeleteUser(id)
		if check {
			controllers.SendResponse(w, r, "Delete sukses!", 200)
		} else {
			controllers.SendResponse(w, r, "Id User tidak ditemukan, tidak dapat di delete...", 400)
		}
	})

	m.RunOnAddr(":8080")
}
