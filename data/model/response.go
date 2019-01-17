package model

import "time"

//RPeticion struct
type RPeticion struct {
	ID               int
	Palabra          string
	Fecha            time.Time
	lenguaje         string
	palabratraducida string
	fechaconsulta    time.Time
}
