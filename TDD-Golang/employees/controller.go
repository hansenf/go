package employees

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

type Controller struct {
	Repository Repository
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func (c Controller) Index(w http.ResponseWriter, r *http.Request) {
	employees := c.Repository.GetEmployees()
	tmpl.ExecuteTemplate(w, "Index", employees)

	return
}

func (c *Controller) New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)

	return
}

func (c *Controller) InsertEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	age, _ := strconv.Atoi(r.FormValue("age"))

	employee.Name = r.FormValue("name")
	employee.Age = age
	employee.Address = r.FormValue("address")

	err := c.Repository.InsertEmployee(employee)
	checkErr(err)

	employees := c.Repository.GetEmployees()
	tmpl.ExecuteTemplate(w, "Index", employees)

	return
}

func (c *Controller) SearchEmployee(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	employees := c.Repository.GetEmployeesByString(query)
	tmpl.ExecuteTemplate(w, "Index", employees)

	return
}

func (c Controller) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee *Employee
	id, _ := strconv.Atoi(r.FormValue("id"))
	age, _ := strconv.Atoi(r.FormValue("age"))

	employee.ID = id
	employee.Name = r.FormValue("name")
	employee.Age = age
	employee.Address = r.FormValue("address")

	err := c.Repository.UpdateEmployee(employee)
	checkErr(err)
	tmpl.ExecuteTemplate(w, "Show", employee)

	return
}

func (c *Controller) GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	employeeid, err := strconv.Atoi(id)
	checkErr(err)

	employee := c.Repository.GetEmployeeById(employeeid)
	tmpl.ExecuteTemplate(w, "Show", employee)

	return
}

func (c *Controller) GetEmployeeForEdit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	employeeid, err := strconv.Atoi(id)
	checkErr(err)

	employee := c.Repository.GetEmployeeById(employeeid)
	tmpl.ExecuteTemplate(w, "Edit", employee)

	return
}

func (c *Controller) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	employeeId, err := strconv.Atoi(id)
	checkErr(err)

	err = c.Repository.DeleteEmployee(employeeId)
	checkErr(err)

	employees := c.Repository.GetEmployees()
	tmpl.ExecuteTemplate(w, "Index", employees)

	return
}