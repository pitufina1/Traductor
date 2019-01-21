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

//Insert Función que inserta una petición en la base de datos local
func Insert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathEnvioPeticion {
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
		var peticion model.Peticion
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &peticion)

		if peticion.Palabra != "" {
			peticion.Palabra = strings.ToUpper(peticion.Palabra)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La petición está vacía")
			return
		}

		w.WriteHeader(http.StatusOK)

		w.Header().Add("Content-Type", "application/json")

		respuesta, _ := json.Marshal(peticion)
		fmt.Fprint(w, string(respuesta))

		go client.InsertarPeticion(&peticion)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//List Función que devuelve las peticiones de la base de datos dado un filtro
func List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathListadoPeticiones {
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
			fmt.Fprintln(w, "La petición no pudo ser parseada")
			fmt.Fprintln(w, e.Error())
			return
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//Update Función que actualiza una petición en la base de datos local
func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathEnvioPeticion {
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
		var peticion model.Peticion
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &peticion)

		if peticion.Palabra != "" {
			peticion.Palabra = strings.ToUpper(peticion.Palabra)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La petición está vacía")
			return
		}

		w.WriteHeader(http.StatusOK)

		w.Header().Add("Content-Type", "application/json")

		respuesta, _ := json.Marshal(peticion)
		fmt.Fprint(w, string(respuesta))

		go client.InsertarPeticion(&peticion)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}
