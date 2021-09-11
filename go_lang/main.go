package main

import (
	"admiralhonda/trace"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/signature"
)

// templ represents a single template
type templateHandler struct {
	filename string
	templ    *template.Template
}

// Handle is a http.HandleFunc that renders this template.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if t.templ == nil {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	}
	t.templ.Execute(w, r)
}

var addr = flag.String("addr", ":8080", "アプリのポート番号")

func main() {

	flag.Parse()
	// Oauthのセットアップ
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		google.New("389525521286-h8m6i9pk7itnitbjti1cn53mr66ptscg.apps.googleusercontent.com",
			"KTXfWwpj906Lg84wNH8-k4M9", "http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/room", r)

	// 認証関連のPATH
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)

	// get the room going
	go r.run()
	log.Println("Webサーバーを開始します。:", *addr)

	// start the web server
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
