package main

import (
	"log"
	"net/http"
	"wxcloudrun-golang/apis"
)

func main() {
	//if err := db.Init(); err != nil {
	//	panic(fmt.Sprintf("mysql init failed with %+v", err))
	//}

	http.HandleFunc("/market_code/apply_code", apis.ApplyCodeHandler)
	http.HandleFunc("/market_code/apply_code_query", apis.ApplyCodeQueryHandler)
	http.HandleFunc("/market_code/apply_code_download", apis.GetApplyCodeDownloadHandler)
	http.HandleFunc("/market_code/code_active", apis.CodeActiveHandler)
	http.HandleFunc("/market_code/code_active_query", apis.CodeActiveQueryHandler)

	//http.HandleFunc("/api/count", service.CounterHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
