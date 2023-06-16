package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"gamex/config"
	"gamex/entities"
	"gamex/libraries"
	"gamex/models"

	"golang.org/x/crypto/bcrypt"
)

var validation = libraries.NewValidation()
var userModel = models.NewUserModel()

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if (session.Values["loggedIn"] != true) || (session.Values["role"] == "user") {
		http.Redirect(w, r, "/noAccess", http.StatusSeeOther)
	} else {
		user, _ := userModel.FindAll()
		//currentuser := userModel.currentUser(id)
		data := map[string]interface{}{
			"user":    user,
			"offline": session.Values["loggedIn"] != true,
		}

		temp, err := template.ParseFiles("templates/admin.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, data)
	}
}

// func Add(response http.ResponseWriter, request *http.Request) {

// 	if request.Method == http.MethodGet {
// 		temp, err := template.ParseFiles("templates/admin_add.html")
// 		if err != nil {
// 			panic(err)
// 		}
// 		temp.Execute(response, nil)
// 	} else if request.Method == http.MethodPost {

// 		request.ParseForm()

// 		var user entities.User
// 		user.Full_name = request.Form.Get("fullname")
// 		user.Email = request.Form.Get("email")
// 		user.Username = request.Form.Get("username")
// 		user.Password = request.Form.Get("password")

// 		var data = make(map[string]interface{})

// 		vErrors := validation.Struct(user)

// 		if vErrors != nil {
// 			data["user"] = user
// 			data["validation"] = vErrors
// 		} else {
// 			data["user1"] = "User data successfully saved"
// 			userModel.Create(user)
// 		}

// 		temp, _ := template.ParseFiles("templates/admin_add.html")
// 		temp.Execute(response, data)
// 	}
// }

func Add(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("templates/admin_add.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		// melakukan proses registrasi

		// mengambil inputan form
		r.ParseForm()

		user := entities.User{
			Full_name: r.Form.Get("fullname"),
			Email:     r.Form.Get("email"),
			Username:  r.Form.Get("username"),
			Password:  r.Form.Get("password"),
			Cpassword: r.Form.Get("cpassword"),
			Role:      r.Form.Get("role"),
		}

		errorMessages := validation.Struct(user)

		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
				"user":       user,
			}
			fmt.Println(errorMessages)
			temp, _ := template.ParseFiles("templates/admin_add.html")
			temp.Execute(w, data)
		} else {

			// hashPassword
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)

			
			userModel.CreateUser(user)

			data := map[string]interface{}{
				"user1": "User added successfully",
			}
			temp, _ := template.ParseFiles("templates/admin_add.html")
			temp.Execute(w, data)
		}
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var user entities.User
		userModel.Find(id, &user)

		data := map[string]interface{}{
			"user": user,
		}

		temp, err := template.ParseFiles("templates/admin_update.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var user entities.User
		user.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		user.Full_name = request.Form.Get("fullname")
		user.Email = request.Form.Get("email")
		user.Username = request.Form.Get("username")
		user.Role = request.Form.Get("role")

		var data = make(map[string]interface{})

		//vErrors := validation.Struct(user)

		// if vErrors != nil {
		// 	data["user"] = user
		// 	data["validation"] = vErrors
		// } else {
		data["user1"] = "User data successfully updated"
		userModel.Update(user)
		//fmt.Println(vErrors)
		//}
		http.Redirect(response, request, "/admin", http.StatusSeeOther)

	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	userModel.Delete(id)

	http.Redirect(response, request, "/admin", http.StatusSeeOther)
}
