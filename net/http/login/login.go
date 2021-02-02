package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

var loginForm = []byte(`
	<html>
	<body>
		<form action="/" method="post">
			Login: <input type="text" name="login"/>
			Password: <input type="password" name="password"/>
			<input type="submit" value="Login"/>
		</form>
	</body>
	</html>
`)

func main() {
	http.HandleFunc("/", handle)

	fmt.Println("Serve at ", port)
	http.ListenAndServe(port, nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write(loginForm)
		return
	}

	login := r.FormValue("login")
	fmt.Println("login", login)

	password := r.FormValue("password")
	fmt.Println("password", password)

	w.Write([]byte("Ok!\n"))
}
