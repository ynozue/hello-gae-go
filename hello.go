package main

import (
	"net/http"
	"encoding/json"
)

type Result struct {
	Status int
	Description string
}

func init() {
	http.HandleFunc("/hello", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var result Result
	switch r.Method {
	case http.MethodGet:
		v := r.FormValue("get_value")
		result = Result{http.StatusOK, "http get result [" + v + "]"}
	case http.MethodPost:
		v := r.FormValue("post_value")
		result = Result{http.StatusOK, "http post result [" + v + "]"}
	default:
		result = Result{http.StatusNotImplemented, "not implemented http mehtod"}
	}
	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}
