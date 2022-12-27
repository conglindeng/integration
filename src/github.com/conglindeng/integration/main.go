// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	// "fmt" has methods for formatted I/O operations (like printing to the console)
	"fmt"
	"io"

	// The "net/http" library has methods to implement HTTP clients and servers
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// The "HandleFunc" method accepts a path and a function as arguments
	// (Yes, we can pass functions as arguments, and even trat them like variables in Go)
	// However, the handler function has to have the appropriate signature (as described by the "handler" function below)
	//http.HandleFunc("/", handler)

	// After defining our server, we finally "listen and serve" on port 8080
	// The second argument is the handler, which we will come to later on, but for now it is left as nil,
	// and the handler defined above (in "HandleFunc") is used
	http.ListenAndServe(":9090", newRouter())
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/hello", helloHandler).Methods("GET")

	dir := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(dir))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	//io.Copy(,r.Body)
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("body: %s \n", string(b))
	s := r.Header.Get("test")
	fmt.Println(s)
	fmt.Fprintf(w, "Hello World!")
}
