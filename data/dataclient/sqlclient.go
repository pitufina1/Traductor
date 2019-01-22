package dataclient

import (
	"database/sql"
	"fmt"
	"traductor/data/model"

	_ "github.com/go-sql-driver/mysql" ///El driver se registra en database/sql en su función Init(). Es usado internamente por éste
)

//InsertarPalabra test
func InsertarPalabra(objeto *model.Palabra) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Traductor")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Query("INSERT INTO Traduccion(texto, palabra_id, idioma_id) VALUES (?, ?, ?)", objeto.Texto)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//ActualizarPalabra test
func ActualizarPalabra(objeto *model.Filtro) []model.RTraduccion {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Traductor")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Query("UPDATE Traduccion SET T.Texto = P.Texto FROM Traduccion T INNER JOIN Palabra P ON T.Palabra_ID = P.ID", objeto.Nombre)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//ListarRegistros test
func ListarRegistros(objeto *model.Filtro) []model.RTraduccion {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Traductor")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT * FROM Idioma WHERE (nombre <= '" + objeto.Nombre + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT * FROM Idioma WHERE (nombre >= ?) ", objeto.Nombre)

	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	resultado := make([]model.RTraduccion, 0)
	for query.Next() {
		var fila = model.RTraduccion{}

		err = query.Scan(&fila.ID, &fila.Texto, &fila.Palabra_ID, &fila.Idioma_ID)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, fila)
	}
	return resultado
}
