package main

import (
	"fmt"
	"net/http"
	"time"
)

var form_html = `<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Вход в систему</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }
        .login-container {
            background-color: white;
            padding: 20px;
            border-radius: 5px;
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
        }
        h2 {
            margin-bottom: 20px;
        }
        input[type="text"], input[type="password"] {
            width: 100%;
            padding: 10px;
            margin: 10px 0;
            border: 1px solid #ccc;
            border-radius: 5px;
        }
        input[type="submit"] {
            background-color: #5cb85c;
            color: white;
            border: none;
            padding: 10px;
            border-radius: 5px;
            cursor: pointer;
            width: 100%;
        }
        input[type="submit"]:hover {
            background-color: #4cae4c;
        }
    </style>
</head>
<body>
    <div class="login-container">
        <h2>Вход в систему</h2>
        <form action="/login" method="post">
            <label for="username">Логин:</label>
            <input type="text" id="username" name="username" required>

            <label for="password">Пароль:</label>
            <input type="password" id="password" name="password" required>

            <input type="submit" value="Войти">
        </form>
    </div>
</body>
</html>`

func RunServer(adr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("param")
		key := r.FormValue("key")
		fmt.Fprintf(w, "hi, url; %s, param: %v, key: %v", r.URL.String(), param, key)
	})
	server := http.Server{
		Handler:      mux,
		Addr:         adr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Printf("server starts in address : %v\n", adr)
	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("сервер не запустился")
	}
	return nil
}

func main() {
	go RunServer(":8080")
	RunServer(":8081")
}
