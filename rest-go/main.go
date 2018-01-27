package main
	import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/", jsonHandler)
	http.ListenAndServe(":8080", nil)
}
func jsonHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello Docker World - GO REST example")
}