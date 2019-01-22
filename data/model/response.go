package model

//RPalabra struct
type RPalabra struct {
	ID    int
	Texto string
}

type RIdioma struct {
	ID     int
	Nombre string
}

type RTraduccion struct {
	ID         int
	Texto      string
	Palabra_ID int
	Idioma_ID  int
}
