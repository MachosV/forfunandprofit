package models

import "time"

/*
IndexData struct holds the data
to be presented in the index view
*/
type IndexData struct {
	Visits uint64
	Uptime time.Duration
}
