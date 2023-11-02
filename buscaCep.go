package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const ROOT_PATH string = "/"
const BUSCA_CEP_PATH string = "/buscaCep"
const BAD_PARAM_MESSAGE string = "O parâmetro cep é obrigatório"
const ERROR_FETCHING_CEP string = "Erro ao buscar o cep"
const VIA_CEP_URL string = "https://viacep.com.br/ws/%s/json/"
const QUERY_PARAM_CEP string = "cep"
const EMPTY_STRING string = ""

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func BuscaCepHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != BUSCA_CEP_PATH {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := req.URL.Query().Get(QUERY_PARAM_CEP)
	if cepParam == EMPTY_STRING {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BAD_PARAM_MESSAGE))
		return
	}

	data, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ERROR_FETCHING_CEP))
		return
	}

	returnJsonCepData(w, data)
}

func BuscaCep(cep string) (*ViaCEP, error) {
	preparedUrl := fmt.Sprintf(VIA_CEP_URL, cep)
	resp, error := http.Get(preparedUrl)
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var data ViaCEP
	error = json.Unmarshal(body, &data)
	if error != nil {
		return nil, error
	}

	return &data, nil
}

func returnJsonCepData(w http.ResponseWriter, data *ViaCEP) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
