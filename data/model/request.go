package model

import "time"

//Peticion struct
type Peticion struct {
	Palabra string
}

//Filtro struct
type Filtro struct {
	Palabra string
}

//Idioma struct
type Idioma struct {
	Nombre string
}

//Traduccion struct
type Traduccion struct {
	PalabraTraducida string
	FechaInicio      time.Time
	FechaConsulta    time.Time
}
