package main

import (
	"admiralhonda/trace"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
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

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/", (&templateHandler{filename: "chat.html"}))
	http.Handle("/room", r)
	// get the room going
	go r.run()
	log.Println("Webサーバーを開始します。:", *addr)

	// start the web server
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
