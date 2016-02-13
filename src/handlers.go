package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Usage shows usage
func Usage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Usage!\n")
}

// SubmissionShow return submission status
func SubmissionShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Show submission for %s!\n", vars["submissionID"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	submission := Submission{ID: 20, Name: "name", Completed: true}

	if err := json.NewEncoder(w).Encode(submission); err != nil {
		panic(err)
	}
	return
}

// SubmissionCreate accepts POST request
func SubmissionCreate(w http.ResponseWriter, r *http.Request) {
	var submission Submission

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &submission); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(submission); err != nil {
		panic(err)
	}
}
