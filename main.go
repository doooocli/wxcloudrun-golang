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

	mux := http.NewServeMux()

	mux.HandleFunc("/market_code/apply_code", apis.ApplyCodeHandler)
	mux.HandleFunc("/market_code/apply_code_query", apis.ApplyCodeQueryHandler)
	mux.HandleFunc("/market_code/apply_code_download", apis.GetApplyCodeDownloadHandler)
	mux.HandleFunc("/market_code/code_active", apis.CodeActiveHandler)
	mux.HandleFunc("/market_code/code_active_query", apis.CodeActiveQueryHandler)
	mux.HandleFunc("/market_code/ticket_to_code", apis.TicketToCodeHandler)

	//mux.HandleFunc("/api/count", service.CounterHandler)

	log.Fatal(http.ListenAndServe(":80", applyMiddleware(mux, errorHandlerMiddleware, authHandlerMiddleware)))
}

func errorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func authHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func applyMiddleware(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
