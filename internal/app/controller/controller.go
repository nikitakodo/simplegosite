package controller

import (
	"net/http"
)

type Reader interface {
	Read(w http.ResponseWriter, r *http.Request)
}

type BulkReader interface {
	ReadAll(w http.ResponseWriter, r *http.Request)
}

type Creator interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type Updater interface {
	Update(w http.ResponseWriter, r *http.Request)
}

type Deleter interface {
	Delete(w http.ResponseWriter, r *http.Request)
}
