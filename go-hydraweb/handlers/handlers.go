package handlers

import (
	"fmt"
	"net/http"
)

// Repository is the repository type
type Repository struct {
}

// creates new repository
func NewRepo() *Repository {
	r := Repository{}

	return &r
}

func (rp *Repository) HandleRoot(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to Hydra software system")

}
