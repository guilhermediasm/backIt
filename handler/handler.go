package handler

import (
	"github.com/guilhermediasm/backIt/config"
)

var (
	logger *config.Logger
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	// db = config.GetSQLite()
}
