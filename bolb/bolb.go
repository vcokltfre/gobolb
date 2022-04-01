package bolb

import (
	"fmt"
	"net/http"
)

func Bolb() {
	fmt.Println("bolb")
}

func Serve(port int, host string) {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bolb"))
	})

	http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}
