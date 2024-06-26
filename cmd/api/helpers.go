package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Retrieve the "id" URL parameter from the current request context, then convert it to
// and integer and return it. If the operation isn't successfull, return 0 and an error.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id")
	}
	return id, nil
}

// Define a writeJSON() helper for sending responses.
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	// Encode data to JSON
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Append a new line to make it easier to view
	js = append(js, '\n')

	// At this point, we know that we won't encounter any more errors before writing
	// the response, so it's safe to add any headers that we want to include
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Add the Content-Type: application/json header, then write the status code and
	// JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
