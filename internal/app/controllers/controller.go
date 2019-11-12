package controllers

import (
	"net/http"
)

type ControllerInterface interface {
	Action(w http.ResponseWriter, r *http.Request)
}
