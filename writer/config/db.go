package config

import "fmt"

// PostgreSQL configs
const (
	host     = "127.0.0.1"
	port     = "5432"
	user     = "patrick"
	password = "erasmusmundus"
	database = "ciklum"
)

// DbRoute contain full DataBase route
var DbRoute = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
	host, port, user, database, password)
