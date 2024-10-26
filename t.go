package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func MainFunc(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	loggout := (err != http.ErrNoCookie)
	if loggout {
		fmt.Fprintf(w, "<a href=/logout>logout<a>")
		fmt.Fprintf(w, "ваш логин %v", session)

	} else {
		fmt.Fprintf(w, "<a href=/login>login<a>")
		fmt.Fprintf(w, "вы должны зарегистрироваться")
	}
}
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		data, _ := ioutil.ReadFile("index.html")

		w.Write(data)
		return
	}
	data := r.FormValue("username")
	expire := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Expires: expire,
		Value:   data,
		Name:    "session_id",
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)

}
func LogoutFunc(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	http.Redirect(w, r, "/", http.StatusFound)
}
func main() {
	http.HandleFunc("/", MainFunc)
	http.HandleFunc("/login", LoginFunc)
	http.HandleFunc("/logout", LogoutFunc)
	http.ListenAndServe(":8080", nil)
}
