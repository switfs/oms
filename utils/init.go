package utils

import (
	logging "github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var (
	log = logging.Logger("utils")

	GVA_DB *gorm.DB
)
