package main

import (
	"net/http"
)

func main() {
	// Cria o servidor HTTP responsável por receber as requisições e fazer o roteamento para os handlers específicos de cada rota (path) da aplicação.
	mux := http.NewServeMux()

	mux.Handle(ROOT_PATH, BuscaCepPage{Title: "Busca Cep - Home"})
	mux.HandleFunc(BUSCA_CEP_PATH, BuscaCepHandler)
	http.ListenAndServe(":8080", mux)
}

type BuscaCepPage struct {
	Title string
}

func (b BuscaCepPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.Title))
}
