package handler

import (
	"github.com/guilhermediasm/back-It/config"
)

var (
	logger *config.Logger
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	// db = config.GetSQLite()
}
