package main

import (
	"log"
	"net/http"
	"wxcloudrun-golang/service"
)

func main() {
	//if err := db.Init(); err != nil {
	//	panic(fmt.Sprintf("mysql init failed with %+v", err))
	//}

	http.HandleFunc("/", service.ApplycodeHandler)
	//http.HandleFunc("/api/count", service.CounterHandler)

	log.Fatal(http.ListenAndServe(":8863", nil))
}
