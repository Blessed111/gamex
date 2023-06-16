package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"gamex/config"
	"gamex/entities"
	"gamex/libraries"
	"gamex/models"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// var debug = true
// if(debug){
// 	fmt.Println(err)
// }

type UserInput struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

var userModel = models.NewUserModel()
var validation = libraries.NewValidation()
var commentModel = models.NewCommentModel()

func HomePage(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	var user models.UserModel
	//currentUser, _ := userModel.FindUser()
	// var current_user
	// for _, element := range user {
	// 	if(user)==
	// }
	if r.Method == http.MethodGet {
		data := map[string]interface{}{
			"isnotAuthorized": session.Values["loggedIn"] != true,
			"full_name":       session.Values["full_name"],
			"role":            session.Values["role"],
			"email":           session.Values["email"],
			"password":        session.Values["password"],
			"username":        session.Values["username"],
			"user":            user,
		}
		temp, _ := template.ParseFiles("templates/index.html")
		temp.Execute(w, data)
	}
}

func Device(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			if r.Method == http.MethodGet {
				data := map[string]interface{}{
					"isnotAuthorized": session.Values["loggedIn"] != true,
					"full_name":       session.Values["full_name"],
					"username":        session.Values["username"],
				}

				temp, _ := template.ParseFiles("templates/payment.html")
				temp.Execute(w, data)
			}
		}

	}
}

func Device1(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			if r.Method == http.MethodGet {
				data := map[string]interface{}{
					"isnotAuthorized": session.Values["loggedIn"] != true,
					"full_name":       session.Values["full_name"],
					"username":        session.Values["username"],
				}

				temp, _ := template.ParseFiles("templates/new_payment.html")
				temp.Execute(w, data)
			}
		}

	}
}

func Index(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			var comment entities.Comment

			if r.Method == http.MethodGet {
				data := map[string]interface{}{
					"full_name": session.Values["full_name"],
					"title":     session.Values["title"],
					"username":  session.Values["username"],
					"role":      session.Values["role"],
					"comment":   comment,
				}
				temp, _ := template.ParseFiles("templates/index2.html")
				temp.Execute(w, data)

			} else if r.Method == http.MethodPost {
				r.ParseForm()

				comment := entities.Comment{
					Title: r.Form.Get("title"),
				}

				commentModel.Createcomment(comment)

				data := map[string]interface{}{
					"full_name": session.Values["full_name"],
					"title":     session.Values["title"],
					"comment":   comment,
				}
				temp, _ := template.ParseFiles("templates/index2.html")
				temp.Execute(w, data)
			}
		}
	}

}

// func Index3(w http.ResponseWriter, r *http.Request) {
// 	session, _ := config.Store.Get(r, config.SESSION_ID)

// 	comment, _ := commentModel.FindAllComment()
// 	//var user entities.User
// 	if r.Method == http.MethodGet {
// 		data := map[string]interface{}{
// 			"isnotAuthorized": session.Values["loggedIn"] != true,
// 			"full_name":       session.Values["full_name"],
// 			"role":            session.Values["role"],
// 			"user":            session.Values["email"],
// 			"comment":         comment,
// 		}
// 		temp, _ := template.ParseFiles("templates/index2.html")
// 		temp.Execute(w, data)
// 	} else if r.Method == http.MethodPost {
// 		// melakukan proses registrasi

// 		// mengambil inputan form
// 		r.ParseForm()

// 		comment := entities.Comment{
// 			Title: r.Form.Get("comment"),
// 		}

// 		data := map[string]interface{}{

// 			"comment": comment,
// 			"title":   session.Values["title"],
// 		}
// 		temp, _ := template.ParseFiles("templates/index2.html")
// 		temp.Execute(w, data)
// 	} else {
// 		var comment entities.Comment
// 		commentModel.Createcomment(comment)

// 		data := map[string]interface{}{
// 			"comment": comment,
// 			"title":   session.Values["title"],
// 		}
// 		temp, _ := template.ParseFiles("templates/index2.html")
// 		temp.Execute(w, data)
// 	}

// }

func Index2(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if r.Method == http.MethodGet {
		users, _ := userModel.FindAll()
		comments, _ := commentModel.FindAllComment()

		data := map[string]interface{}{
			"comments": comments,
			"users":    users,
			"username": session.Values["username"],
		}

		temp, err := template.ParseFiles("templates/index2.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, data)

	} else if r.Method == http.MethodPost {
		r.ParseForm()

		comment := entities.Comment{
			Title: r.Form.Get("title"),
		}
		// insert ke database
		commentModel.Createcomment(comment)

		data := map[string]interface{}{
			"com": "Comment added",
		}
		temp, _ := template.ParseFiles("templates/index2.html")
		temp.Execute(w, data)

	}

}

type Comment struct {
	Comment string
}

type CommentModel struct {
}

func SaveComment(w http.ResponseWriter, r *http.Request) {
	comment := r.FormValue("title")

	if comment == "" {
		data := map[string]interface{}{
			"err": "Fill in the text box",
		}
		tmp := template.Must(template.ParseFiles("templates/index2.html"))
		tmp.Execute(w, data)
	} else {
		db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/hotel")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `comment` (`comment`) VALUES('%s')", comment))
		if err != nil {
			panic(err)
		}
		defer insert.Close()
		result, err := db.Query("SELECT comment FROM comment")
		if err != nil {
			panic(err.Error())
		}

		for result.Next() {
			// doc, err := html.Parse(strings.NewReader(s))
			if err != nil {
				panic(err)
			}
			// list := getElementById(doc, "list")
			var CommentModel CommentModel
			comments, _ := CommentModel.FindAllComment()
			//comment := r.FormValue("comment")
			data := map[string]interface{}{
				"comments": comments,
			}
			tmp := template.Must(template.ParseFiles("templates/index2.html"))
			tmp.Execute(w, data)
		}
	}
}

func (*CommentModel) FindAllComment() ([]Comment, error) {
	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/hotel")
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select comment from comment order by id desc")
		if err2 != nil {
			return nil, err2
		} else {
			var comments []Comment
			for rows.Next() {
				var comment Comment
				rows.Scan(&comment.Comment)
				comments = append(comments, comment)
			}
			return comments, nil
		}
	}
}

// func Index2(w http.ResponseWriter, r *http.Request) {
// 	session, _ := config.Store.Get(r, config.SESSION_ID)
// 	//comment, _ := commentModel.FindAllComment()
// 	user, _ := userModel.FindAll()

// 	data := map[string]interface{}{
// 		"comment":  user,
// 		"username": session.Values["full_name"],
// 	}

// 	temp, err := template.ParseFiles("templates/index2.html")
// 	if err != nil {
// 		panic(err)
// 	}
// 	temp.Execute(w, data)
// }

// func Index2(w http.ResponseWriter, r *http.Request) {
// 	session, _ := config.Store.Get(r, config.SESSION_ID)
// 	if r.Method == http.MethodGet {

// 		temp, _ := template.ParseFiles("templates/index2.html")
// 		temp.Execute(w, nil)

// 	} else if r.Method == http.MethodPost {

// 		r.ParseForm()

// 		comment := entities.Comment{
// 			Title: r.Form.Get("title"),
// 		}

// 		data := map[string]interface{}{
// 			"comment": comment,
// 		}
// 		temp, _ := template.ParseFiles("templates/index2.html")
// 		temp.Execute(w, data)

// 		commentModel.Createcomment(comment)
// 	}
// }

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		session, _ := config.Store.Get(r, config.SESSION_ID)
		if session.Values["loggedIn"] == true {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
		// js := json.NewDecoder(r.Body).Decode(&r)
		temp, _ := template.ParseFiles("templates/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {

		r.ParseForm()
		UserInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		errorMessages := validation.Struct(UserInput)

		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
			}

			temp, _ := template.ParseFiles("templates/login.html")
			temp.Execute(w, data)

		} else {
			if UserInput.Username == "admin" && UserInput.Password == "admin" {
				http.Redirect(w, r, "/admin", http.StatusSeeOther)
			}
			var user entities.User
			userModel.Where(&user, "username", UserInput.Username)

			var message error
			if user.Username == "" {
				message = errors.New("Wrong username or password!")
			} else {

				errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
				if errPassword != nil {
					message = errors.New("Wrong username or password!")
				}
			}

			if message != nil {

				data := map[string]interface{}{
					"error": message,
				}

				temp, _ := template.ParseFiles("templates/login.html")
				temp.Execute(w, data)
			} else {
				// set session
				session, _ := config.Store.Get(r, config.SESSION_ID)

				session.Values["loggedIn"] = true
				session.Values["email"] = user.Email
				session.Values["username"] = user.Username
				session.Values["full_name"] = user.Full_name
				session.Values["role"] = user.Role
				session.Values["id"] = user.Id
				session.Save(r, w)

				http.Redirect(w, r, "/home", http.StatusSeeOther)
			}
		}

	}

}

// func noAccess(w http.ResponseWriter, r *http.Request) {
// 	session, _ := config.Store.Get(r, config.SESSION_ID)
// 	if r.Method == http.MethodGet {
// 		data := map[string]interface{}{
// 			"isnotAuthorized": session.Values["loggedIn"] != true,
// 		}
// 		temp, _ := template.ParseFiles("templates/no_access.html")
// 		temp.Execute(w, data)
// 	}
// }

func Access(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/no_access.html")
	if err != nil {
		fmt.Println("not connected!")
	}

	t.Execute(w, nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	// delete session
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("templates/register.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {

		r.ParseForm()

		user := entities.User{
			Full_name: r.Form.Get("full_name"),
			Email:     r.Form.Get("email"),
			Username:  r.Form.Get("username"),
			Password:  r.Form.Get("password"),
			Cpassword: r.Form.Get("cpassword"),
		}

		errorMessages := validation.Struct(user)

		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
				"user":       user,
			}

			temp, _ := template.ParseFiles("templates/register.html")
			temp.Execute(w, data)
		} else {

			// hashPassword
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)

			userModel.Create(user)

			data := map[string]interface{}{
				"user1": "Registration is successful",
			}
			temp, _ := template.ParseFiles("templates/register.html")
			temp.Execute(w, data)
		}
	}

}
