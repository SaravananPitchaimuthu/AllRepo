package main

import (
	"fmt"
	"net/http"
)

func middleware(OriginalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Excuting middleware before request phase")
		OriginalHandler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after request phase")
	})
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing main handler")
	w.Write([]byte("OK"))
}
func main() {
	originalHandler := http.HandlerFunc(handle)
	http.Handle("/", middleware(originalHandler))
	http.ListenAndServe(":8000", nil)

}
