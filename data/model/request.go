package model

import "time"

//Peticion struct
type Peticion struct {
	Palabra string
}

//Filtro struct
type Filtro struct {
	Fecha time.Time
}

//Idioma struct
type Idioma struct {
	lenguaje string
}

//Traduccion struct
type Traduccion struct {
	palabratraducida string
	fechaconsulta    time.Time
}
