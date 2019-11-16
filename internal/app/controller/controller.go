package controller

import (
	"net/http"
)

type Interface interface {
	Action(w http.ResponseWriter, r *http.Request)
}

type ResourceInterface interface {
	GetName() string
	Read(w http.ResponseWriter, r *http.Request)
	ReadAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
