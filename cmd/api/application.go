package api

import (
	"log"
)

type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
