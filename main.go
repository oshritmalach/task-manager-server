package main

import (
	"Checkmarx/router"
	"log"
	"net/http"
)

const serverPort = "8083"

func main() {
	r := router.NewRouter()
	log.Printf("Server is running on port %s\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, enableCors(r)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
