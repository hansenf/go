package employees

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Repository struct{}

var db *sql.DB

func (r *Repository) InitDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)
	dbp, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	db = dbp

	return db, nil
}

func (r *Repository) connectDB(databse *sql.DB) {
	db = databse

	return
}

func (r *Repository) GetEmployees() (results Employees) {
	rows, err := db.Query("SELECT ID, Name, Age, Address FROM employees")
	checkErr(err)
	var row Employee

	for rows.Next() {
		err = rows.Scan(&row.ID, &row.Name, &row.Age, &row.Address)
		checkErr(err)
		results = append(results, row)
	}

	return results
}

func (r *Repository) GetEmployeeById(id int) Employee {
	var row Employee
	err := db.QueryRow("SELECT ID, Name, Age, Address FROM employees WHERE Id=$1;", id).Scan(&row.ID, &row.Name, &row.Age, &row.Address)
	checkErr(err)

	return row
}

func (r *Repository) GetEmployeesByString(query string) Employees {
	rows, err := db.Query("SELECT id, Name, Age, Address FROM employees WHERE LOWER(Name) like '%' || LOWER($1) || '%';", query)
	checkErr(err)
	results := Employees{}
	var row Employee

	for rows.Next() {
		err = rows.Scan(&row.ID, &row.Name, &row.Age, &row.Address)
		checkErr(err)
		results = append(results, row)
	}

	return results
}

func (r *Repository) InsertEmployee(u Employee) error {
	_, err := db.Exec("INSERT INTO employees(name, age, address) VALUES ($1 ,$2, $3);", u.Name, u.Age, u.Address)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateEmployee(u *Employee) error {
	stmt, err := db.Prepare("UPDATE employees SET Name=$1, Age=$2, Address=$3 WHERE id=$4;")
	checkErr(err)
	_, err = stmt.Exec(u.Name, u.Age, u.Address, u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteEmployee(id int) error {
	stmt, err := db.Prepare("DELETE FROM employees WHERE id=$4")
	checkErr(err)

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}