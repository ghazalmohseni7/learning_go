package Camera

import (
	"net/http"
)

func CameraRouters(mux *http.ServeMux) {
	//mux.HandleFunc("/camera/", ListCamera)
	//mux.HandleFunc("/camera/insert/", InsertCamera)
	mux.HandleFunc("/camera/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			ListCamera(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Handle POST requests for "/camera/insert/"
	mux.HandleFunc("/camera/insert/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			InsertCamera(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/camera/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {

			UpdateCamera(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/camera/delete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			DeleteCamera(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/camera/id/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			RetrieveCamera(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}
