package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	client "traductor/data/dataclient"
	"traductor/data/model"
)

//Insert Función que inserta una palabra en la base de datos local
func Insert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathEnvioPalabras {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var palabra model.Palabra
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &palabra)

		if palabra.Texto != "" {
			palabra.Texto = strings.ToUpper(palabra.Texto)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La palabra está vacía")
			return
		}

		w.WriteHeader(http.StatusOK)

		w.Header().Add("Content-Type", "application/json")

		respuesta, _ := json.Marshal(palabra)
		fmt.Fprint(w, string(respuesta))

		go client.InsertarPalabra(&palabra)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//List Función que devuelve las palabras de la base de datos dado un filtro
func List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathListadoPalabras {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var filtro model.Filtro
		e = json.Unmarshal(bytes, &filtro)

		if e == nil {
			lista := client.ListarRegistros(&filtro)

			w.WriteHeader(http.StatusOK)

			w.Header().Add("Content-Type", "application/json")

			respuesta, _ := json.Marshal(&lista)
			fmt.Fprint(w, string(respuesta))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La palabra no pudo ser parseada")
			fmt.Fprintln(w, e.Error())
			return
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//Update Función que actualiza una palabra en la base de datos local
func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathEnvioPalabras {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var palabra model.Palabra
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &palabra)

		if palabra.Texto != "" {
			palabra.Texto = strings.ToUpper(palabra.Texto)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La palabra está vacía")
			return
		}

		w.WriteHeader(http.StatusOK)

		w.Header().Add("Content-Type", "application/json")

		respuesta, _ := json.Marshal(palabra)
		fmt.Fprint(w, string(respuesta))

		go client.InsertarPalabra(&palabra)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//Idioma Función que devuelve los Idiomas de la base de datos dado un filtro
func Idioma(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathIdioma {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var filtro model.Filtro
		e = json.Unmarshal(bytes, &filtro)

		if e == nil {
			lista := client.ListarRegistros(&filtro)

			w.WriteHeader(http.StatusOK)

			w.Header().Add("Content-Type", "application/json")

			respuesta, _ := json.Marshal(&lista)
			fmt.Fprint(w, string(respuesta))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La palabra no pudo ser parseada")
			fmt.Fprintln(w, e.Error())
			return
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//Traduccion Función que devuelve las traducciones de la base de datos dado un filtro
func Traduccion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathTraduccion {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var filtro model.Filtro
		e = json.Unmarshal(bytes, &filtro)

		if e == nil {
			lista := client.ListarRegistros(&filtro)

			w.WriteHeader(http.StatusOK)

			w.Header().Add("Content-Type", "application/json")

			respuesta, _ := json.Marshal(&lista)
			fmt.Fprint(w, string(respuesta))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La palabra no pudo ser parseada")
			fmt.Fprintln(w, e.Error())
			return
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}
