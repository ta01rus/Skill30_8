package httpserver

import (
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(`<h1>  тестирование </h1>`))
}
