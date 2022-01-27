package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/justinas/alice"
)

type city struct {
	Name string
	Area uint64
}

func ContentTypeCheck(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Currenctly in the content type check middleware")
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("405 - Unsupported Media type. Please Send JSON"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func SetCookieTime(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		fmt.Println("Currently in the set cookie time middleware")
	})
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Got city Name:%s and Area %d", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201-Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405-Method Not allowed"))
	}
}

// func main() {
// 	//http.HandleFunc("/", PostHandler)
// 	originalHandler := http.HandlerFunc(PostHandler)
// 	http.Handle("/city", ContentTypeCheck(SetCookieTime(originalHandler)))
// 	http.ListenAndServe(":8000", nil)
// }

func main() {
	originalHandler := http.HandlerFunc(PostHandler)
	chain := alice.New(ContentTypeCheck, SetCookieTime).Then(originalHandler)
	http.Handle("/city", chain)
	http.ListenAndServe(":8000", nil)
}
