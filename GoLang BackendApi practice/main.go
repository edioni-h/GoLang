package main

import (
	"net/http"
)

/*type server struct {
	addr string
}*/

/*func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello from the server!"))

	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Index Page"))
			return
		case "/users":
			w.Write([]byte("Users Page"))
			return
		}
	default:
		w.Write([]byte("404"))
		return
	}
}

func main() {
	s := &server{addr: ":8080"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}
}*/

func main() {
	api := &api{addr: ":8080"}

	//Initialize the ServeMux
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
