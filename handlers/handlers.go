package handlers

import "net/http"

//PathInicio Ruta raíz
const PathInicio string = "/"

//PathJSFiles Ruta a la carpeta de scripts de javascript
const PathJSFiles string = "/js/"

//PathEnvioPalabra Ruta de envío de palabras
const PathEnvioPalabras string = "/envio"

//PathListadoPalabras Ruta de obtención de las palabras de hoy
const PathListadoPalabras string = "/lista"

//PathActualizacionPalabras Ruta de actualización de las palabras de hoy
const PathActualizacionPalabras string = "/actualizacion"

//PathIdioma Ruta de obtención del listado de idiomas
const PathIdioma string = "/idioma"

//PathTraduccion Ruta de obtención del listado de traducciones
const PathTraduccion string = "/traduccion"

//ManejadorHTTP encapsula como tipo la función de manejo de palabras HTTP, para que sea posible almacenar sus referencias en un diccionario
type ManejadorHTTP = func(w http.ResponseWriter, r *http.Request)

//Lista es el diccionario general de las palabras que son manejadas por nuestro servidor
var Manejadores map[string]ManejadorHTTP

func init() {
	Manejadores = make(map[string]ManejadorHTTP)
	Manejadores[PathInicio] = IndexFile
	Manejadores[PathJSFiles] = JsFile
	Manejadores[PathEnvioPalabras] = Insert
	Manejadores[PathListadoPalabras] = List
	Manejadores[PathActualizacionPalabras] = Update
	Manejadores[PathIdioma] = Idioma
	Manejadores[PathTraduccion] = Traduccion
}
