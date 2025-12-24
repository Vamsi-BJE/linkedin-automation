package logx

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "[linkedin] ", log.LstdFlags)
