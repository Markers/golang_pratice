package main
import (
        "net/http"
        "encoding/json"
)

var users = map[string]*User{}

type User struct {
        Nickname string `json:"nickname"`
        Email    string `json:"email"`
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
        return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
                rw.Header().Add("Content-Type", "application/json")
                next.ServeHTTP(rw, r)
        })
}

func main() {
        mux := http.NewServeMux()

        userHandler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
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
        http.ListenAndServe(":8080", mux)

}
