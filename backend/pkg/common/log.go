package common

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "[spa-ms] ", log.LstdFlags|log.Lshortfile)
