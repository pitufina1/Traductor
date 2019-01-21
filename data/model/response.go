package model

import "time"

//RPeticion struct
type RPeticion struct {
	ID               int
	Palabra          string
	Nombre           string
	PalabraTraducida string
	FechaInicio      time.Time
	FechaConsulta    time.Time
}

type RIdioma struct {
	ID int
}

type RTraduccion struct {
	ID          int
	Peticion_ID int
	Idioma_ID   int
}
