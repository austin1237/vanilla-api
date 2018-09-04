package handler

import (
	"net/http"
	"time"

	"github.com/user/api/hasher"
)

func GetHash() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// Validate form here
		userPassword := r.Form["password"][0]
		hashStr := hasher.HashString(userPassword)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(hashStr))
	})
}

func ArtificalWait(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		next.ServeHTTP(w, r)
	})
}

func PostOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// func ShutDown() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		test := r.Context().Value("test")
// 		fmt.Println(reflect.TypeOf(test))
// 		server, ok := r.Context().Value("test").(*http.Server)

// 		if !ok {
// 			fmt.Println("server not found in context")
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Shutting Down"))
// 		go func() {
// 			fmt.Println("shutting the server down")
// 			// server.SetKeepAlivesEnabled(false)
// 			if err := server.Shutdown(context.Background()); err != nil {
// 				log.Fatal(err)
// 			}
// 		}()
// 	})
// }
