// to see the working, move this file out of this folder, change the name to main.go, into the parent folder and change package dependencyinjection to main, then run go run main.go
package dependencyinjection

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(w io.Writer, name string) {
  fmt.Fprintf(w,"Hello, %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
  Greet(w, "world")
}

func main() {
  log.Fatal(http.ListenAndServe(":8080",http.HandlerFunc(GreetHandler)))
}
