package employees

type Employee struct {
	ID       int    `bson:"_id"`
	Name 	 string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
}

type Employees []Employee

const (
	HOST     = "34.101.185.207"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "pgdev17!"
	DBNAME   = "tmi-1"
)