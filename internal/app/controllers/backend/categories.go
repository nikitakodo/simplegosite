package backend

import (
	"github.com/gorilla/mux"
	"net/http"
	"simplesite/internal/app/store"
	"strconv"
)

type CategoriesResource struct {
	Name       string
	Repository store.RepositoryInterface
}

func NewCategoriesResource() *CategoriesResource {
	return &CategoriesResource{Name: "category"}
}

func (controller CategoriesResource) GetName() string {
	return controller.Name
}

func (controller CategoriesResource) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	item, err := controller.Repository.Find(id)
}

func (controller CategoriesResource) ReadAll(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (controller CategoriesResource) Create(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (controller CategoriesResource) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (controller CategoriesResource) Delete(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
