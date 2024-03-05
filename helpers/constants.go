package helpers

import (
	"time"
)

var Today = time.Now().Truncate(24 * time.Hour)
var Yesterday = time.Now().Local().AddDate(0, 0, -1)
var ThisYr = time.Now().Local().Year()
