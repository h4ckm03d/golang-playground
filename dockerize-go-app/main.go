package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	_ "time/tzdata"
)

func now(w http.ResponseWriter, req *http.Request) {
	local, _ := time.LoadLocation("Asia/Jakarta")
	fmt.Fprintf(w, "Now : %v\n", time.Now().In(local))
}

func main() {
	http.HandleFunc("/now", now)
	fmt.Println("server run on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
