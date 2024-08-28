package person

import (
	"fmt"
	"net/http"
)

type PersonHandler struct {
	Name string
}

func (p *PersonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("This is %s home page", p.Name)))
}
