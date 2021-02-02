package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = ":8081"

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	fmt.Println("Serve at ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session")
	loggedIn := (err != http.ErrNoCookie)
	if loggedIn {
		fmt.Fprintln(w, `<a href="/logout">Logout</a>`)
		fmt.Fprintln(w, `Wilkommen, `, session)
	} else {
		fmt.Fprintln(w, `<a href="/login">Login</a>`)
		fmt.Fprintln(w, `You are need to login.`)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "session",
		Value: "petrovic",
	}
	cookie.Expires = time.Now().Add(1 * time.Minute)
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	http.Redirect(w, r, "/", http.StatusFound)
}
