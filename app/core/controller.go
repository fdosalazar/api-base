package core

import (
	"encoding/json"
	"log"
	"net/http"
)

// Controller handle all base methods
type Controller struct {
}

// SendJSON marshals v to a json struct and sends appropriate headers to w
func (c *Controller) SendJSON(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")
	api := Api{Data: v, Status: code}
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(api); err != nil {
		log.Print(err)
	}
}

// GetContent of the request inside given struct
func (c *Controller) GetContent(v interface{}, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

// HandleError write error on response and return false if there is no error
func (c *Controller) HandleError(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}

	msg := map[string]string{
		"message": "An error occured",
	}

	c.SendJSON(w, &msg, http.StatusInternalServerError)
	return true
}
