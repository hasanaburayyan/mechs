package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"strings"
)

var tmpl *template.Template

var functionMap = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string{
	return strings.TrimSpace(s)[:3]
}

type ValidUser struct {
	Username string
}

func init() {
	tmpl = template.Must(template.New("").Funcs(functionMap).ParseGlob("templates/*"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/", Index)
	mux.GET("/login", Login)
	mux.GET("/auth", CheckAuth)
	mux.POST("/auth", Auth)
	mux.GET("/redirectme", RedirectExample)
	mux.GET("/redirected", Redirected)

	http.ListenAndServe(":8080", mux)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func RedirectExample(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/redirected", http.StatusSeeOther)
}

func CheckAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	authCookie, err := r.Cookie("Auth-Token")
	fmt.Println("Auth Cookie", authCookie)
	if err != nil {
		tmpl.ExecuteTemplate(w, "authstatus.gohtml", "YOU DO NOT HAVE AUTHORIZATION TO VIEW THIS PAGE PLEASE RETURN TO LOGIN PAGE!")
		return
	}
	tmpl.ExecuteTemplate(w, "authstatus.gohtml", fmt.Sprintf("Your Auth Token: %s is still good!\nIt Will Expire at: %v\n", authCookie.Value, authCookie.Expires))
}

func Redirected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl.ExecuteTemplate(w, "redirected.gohtml", nil)
}

func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	bs, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	if username == "Hasan" && password == "abc" {
		http.SetCookie(w, &http.Cookie{
			Name: "Auth-Token",
			Value: "abcxyz123",
			MaxAge: 10,
		})
		http.SetCookie(w, &http.Cookie{
			Name: "Password",
			Value: fmt.Sprintf("%x", bs),
			MaxAge: 60,
		})
		tmpl.ExecuteTemplate(w, "home.gohtml", ValidUser{
			Username: "Hasan",
		})
	} else {
		tmpl.ExecuteTemplate(w, "login.gohtml", "INVALID USERNAME/PASSWORD COMBO")
	}
}
