package myhandlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello struct
type Hello struct {
	l *log.Logger
}

// NewHello function
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("root")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}
