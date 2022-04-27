package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"

	// http-swagger middleware
	// httpSwagger ""
	_ "golang_pratice/docs"
)

var users = map[string]*User{}

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

// @title swaggo first Test
// @version 0.0.1?
// @description 처음써보는 스웨고~
// @host localhost:8000
// @BasePath /users

func main() {
	// mux := http.NewServeMux()

	mux := mux.NewRouter()
	userHandler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(rw).Encode(users)
		case http.MethodPost:
			var user User
			json.NewDecoder(r.Body).Decode(&user)

			users[user.Email] = &user

			json.NewEncoder(rw).Encode(user)
		}
		//rw.Write([]byte("hello"))
	})

	mux.Handle("/users", jsonContentTypeMiddleware(userHandler))

	// 스웨고
	// mux.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// http.Redirect(w, r, "/static/"+"swagger/", http.StatusSeeOther)
	// })

	opts := middleware.SwaggerUIOpts{SpecURL: "/docs/swagger.yaml"}

	sh := middleware.SwaggerUI(opts, nil)
	mux.Handle("/docs", sh)

	var dir string
	mux.PathPrefix("/docs/").Handler(http.FileServer(http.Dir(dir)))

	http.ListenAndServe(":8080", mux)

}
