package main

import (
	"fmt"
	"html/template"
	"net/http"

	authcontroller "gamex/controllers/authcontroller"
	usercontroller "gamex/controllers/usercontroller"
)

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("templates/index.html")
// 	if err != nil {
// 		fmt.Println("not connected!")
// 	}

// 	t.Execute(w, nil)
// }

func hoodi(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/payment_category.html")
	if err != nil {
		fmt.Println("not connected!")
	}

	t.Execute(w, nil)
}

func adminBracket(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/admin_bracket.html")
	if err != nil {
		fmt.Println("not connected!")
	}

	t.Execute(w, nil)
}

func info1(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/payment_info.html")
	if err != nil {
		fmt.Println("not connected!")
	}

	t.Execute(w, nil)
}
func info2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/payment_info2.html")
	if err != nil {
		fmt.Println("not connected!")
	}

	t.Execute(w, nil)
}
func info3(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/payment_info3.html")
	if err != nil {
		fmt.Println("not connected!")
	}

	t.Execute(w, nil)
}

// var router = mux.NewRouter()
// var store = sessions.NewCookieStore([]byte("something-very-secret"))

// func init() {

// 	store.Options = &sessions.Options{
// 		Domain:   "localhost",
// 		Path:     "/",
// 		MaxAge:   3600 * 1, // 1 hour
// 		HttpOnly: true,
// 	}
// }

func main() {

	//router := mux.NewRouter()
	//http.HandleFunc("/", authcontroller.Main)
	http.HandleFunc("/comment", authcontroller.Index2)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)
	http.HandleFunc("/home", authcontroller.HomePage)
	http.HandleFunc("/noAccess", authcontroller.Access)
	http.HandleFunc("/hudi", hoodi)
	http.HandleFunc("/device", authcontroller.Device)
	http.HandleFunc("/admin/bracket", adminBracket)
	http.HandleFunc("/saveComment", authcontroller.SaveComment)
	http.HandleFunc("/info1", info1)
	http.HandleFunc("/info2", info2)
	http.HandleFunc("/info3", info3)

	http.HandleFunc("/admin", usercontroller.Index)
	http.HandleFunc("/admin/user/add", usercontroller.Add)
	http.HandleFunc("/admin/user/edit", usercontroller.Edit)
	http.HandleFunc("/admin/user/delete", usercontroller.Delete)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	//http.Handle("/", router)

	fmt.Println("Server: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
