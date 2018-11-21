package main

import (
	"fmt"
	"os"
)

// PostgreSQL configs
var (
	host     = os.Getenv("DBHost")
	port     = "5432"
	user     = os.Getenv("DBUser")
	password = os.Getenv("DBPassword")
	database = os.Getenv("DBName")
	sslmode  = "disable"
)

// DbRoute contain full DataBase route
var DbRoute = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	host, port, user, database, password, sslmode)
