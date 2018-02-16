package main

import (
	"net/http"
	"fmt"
	"github.com/satori/go.uuid"
	"sync"
	"io"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func main() {
	http.Handle("/hello", helloWorldHandler{})
	http.Handle("/secureHello", authenticate(helloWorldHandler{}))
	http.HandleFunc("/login", handleLogin)

	http.ListenAndServe(":3000", nil)
}

const loginPage = "<html><head><title>Login</title></head><body><form action=\"login\" method=\"post\"> <input type=\"password\" name=\"password\" /> <input type=\"submit\" value=\"login\" /> </form> </body> </html>"


type helloWorldHandler struct {
}

func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!!!")
}

type authenticationMiddleware struct {
	wrappedHandler http.Handler
}

func (h authenticationMiddleware) ServeHTTP(w http.ResponseWriter,r *http.Request) {
}

func authenticate(h http.Handler) authenticationMiddleware {
	return authenticationMiddleware{h}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
}

