package model

import (
	"os"
)

const (
	DBtype = "postgres"

	OK = 0
	IllegalName = -1
	IllegalPassword = -2
	NotOpening = -3
	NotExecution = -4
)

var DBUrl = os.Getenv("DB_URL")
