package routers

import (
	"fmt"
	"net/http"
)

func Healthy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy")
}
