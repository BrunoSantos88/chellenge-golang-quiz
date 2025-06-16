package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func main() {
    // Servir login.html em /login
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./login/login.html")
    })

    // Servir quiz.html em /quiz
    http.HandleFunc("/quiz", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./quiz/quiz.html")
    })

    // Servir arquivos estáticos (ex: CSS, JS, imagens)
    http.Handle("/login/", http.StripPrefix("/login/", http.FileServer(http.Dir("./login"))))
    http.Handle("/quiz/", http.StripPrefix("/quiz/", http.FileServer(http.Dir("./quiz"))))

    // Endpoint de login
    http.HandleFunc("/api/login", loginHandler)

    log.Println("Servidor rodando na porta 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    var creds LoginRequest
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

    if creds.Username == "admin" && creds.Password == "123456" {
        // Login válido, redirecionar para /quiz
        http.Redirect(w, r, "/quiz", http.StatusSeeOther)
    } else {
        // Login inválido
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]string{
            "status":  "error",
            "message": "Usuário ou senha incorretos",
        })
    }
}
