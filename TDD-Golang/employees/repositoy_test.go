package employees

import (
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestGetEmployees(t *testing.T) {
	r := Repository{}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	r.connectDB(db)

	rows := sqlmock.NewRows([]string{"Id", "Name", "Age", "Address"})
	rows.AddRow(1, "George", 23, "Tehran, Iran")
	rows.AddRow(3, "Harry", 12, "US, CA")
	var people = Employees{
		Employee{
			ID:      1,
			Name:    "George",
			Age:     23,
			Address: "Tehran, Iran",
		},
		Employee{
			ID:      3,
			Name:    "Harry",
			Age:     12,
			Address: "US, CA",
		},
	}

	mock.ExpectQuery("^SELECT (.+) FROM employees$").WillReturnRows(rows)
	employees := r.GetEmployees()
	assert.Equal(t, people, employees)

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "There were unfulfilled expectations.")
}

func TestGetEmployeeById(t *testing.T) {
	r := Repository{}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	r.connectDB(db)

	rows := sqlmock.NewRows([]string{"ID", "Name", "Age", "Address"})
	rows.AddRow(4, "David", 29, "indonesia")
	var employee = Employee{
		ID:      4,
		Name:    "David",
		Age:     29,
		Address: "indonesia",
	}
	mock.ExpectQuery("^SELECT (.+) FROM employee").WithArgs(4).WillReturnRows(rows)

	employeeById := r.GetEmployeeById(4)
	assert.Equal(t, employee, employeeById)

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "There were unfulfilled expectations.")
}

func TestGetEmployeesByString(t *testing.T) {
	r := Repository{}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	r.connectDB(db)
	rows := sqlmock.NewRows([]string{"ID", "Name", "Age", "Address"})
	rows.AddRow(4, "David", 29, "indonesia")

	var people = Employees{
		Employee{
			ID:      4,
			Name:    "David",
			Age:     29,
			Address: "indonesia",
		},
	}

	mock.ExpectQuery("^SELECT (.+) FROM").WithArgs("David").WillReturnRows(rows)
	employeeById := r.GetEmployeesByString("David")
	assert.Equal(t, people, employeeById)

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "There were unfulfilled expectations.")
}

func TestInsertEmployee(t *testing.T) {
	r := Repository{}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	r.connectDB(db)

	var insertedId = 4

	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(insertedId)

	mock.ExpectExec("^INSERT INTO").WithArgs("David", 29, "indonesia").WillReturnResult(sqlmock.NewResult(0, 1))

	var employee = Employee{
		ID:      4,
		Name:    "David",
		Age:     29,
		Address: "indonesia",
	}

	err = r.InsertEmployee(employee)
	assert.Nil(t, err)
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "There were unfulfilled expectations.")
}

func TestUpdateEmployee(t *testing.T) {
	r := Repository{}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	r.connectDB(db)

	var employee = &Employee{
		ID:      4,
		Name:    "David",
		Age:     29,
		Address: "indonesia",
	}
	mock.ExpectPrepare("UPDATE").ExpectExec().WithArgs("David", 29, "indonesia", 4).WillReturnResult(sqlmock.NewResult(0, 1))

	err = r.UpdateEmployee(employee)
	assert.Nil(t, err)
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "There were unfulfilled expectations.")
}

func TestDeleteEmployee(t *testing.T) {
	r := Repository{}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	r.connectDB(db)

	var employeeID = 4

	mock.ExpectPrepare("DELETE FROM employee").ExpectExec().WithArgs(employeeID).WillReturnResult(sqlmock.NewResult(0, 1))
	err = r.DeleteEmployee(employeeID)
	assert.Nil(t, err)

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err, "There were unfulfilled expectations.")
}
