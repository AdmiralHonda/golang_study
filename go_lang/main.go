package main

import (
	"github.com/AdmiralHonda/iok/iok_packeage/handler_logic"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler_logic.SayhelloName) //アクセスのルーティングを設定します。
	err := http.ListenAndServe(":9090", nil)         //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
