package handlers

import (
	"fmt"
	"net/http"
)

//HomeHandler for moving on landing page
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Welcome to Home page")
}
