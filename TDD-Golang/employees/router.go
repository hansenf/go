package employees

import (
	"github.com/gorilla/mux"
	"net/http"
)

var controller = &Controller{Repository: Repository{}}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"Insert",
		"POST",
		"/Insert",
		controller.InsertEmployee,
	},
	Route{
		"New",
		"GEt",
		"/New",
		controller.New,
	},
	Route{
		"Edit",
		"GET",
		"/Edit/{id}",
		controller.GetEmployeeForEdit,
	},
	Route{
		"Update",
		"POST",
		"/Update",
		controller.UpdateEmployee,
	},
	Route{
		"Show",
		"GET",
		"/Show/{id}",
		controller.GetEmployee,
	},
	Route{
		"Delete",
		"GET",
		"/Delete/{id}",
		controller.DeleteEmployee,
	},
	Route{
		"Search",
		"POST",
		"/Search",
		controller.SearchEmployee,
	}}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
		http.Handle(route.Pattern, router)
	}

	return router
}