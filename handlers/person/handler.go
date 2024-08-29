package person

import (
	"fmt"
	"net/http"
	"time"
)

type PersonHandler struct {
	Name string
}

func (p *PersonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("This is %s home page", p.Name)))
}

func Hello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func TimeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

func TimeHandler2(format string) http.HandlerFunc {
	return TimeHandler(format).ServeHTTP
}

func TimeHandler3(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
}
