package routes

import (
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header()
	id := r.URL.Path[len("/api/v1/employee/"):]
	w.Write([]byte("Update account for id: " + id))
	w.WriteHeader(http.StatusOK)
}
